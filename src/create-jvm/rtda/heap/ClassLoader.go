package heap

import (
	"create-jvm/classfile"
	"create-jvm/classpath"
	"fmt"
)

/**
加载过程：
	1. 首先找到class文件并把数据读取到内存；
	2. 然后解析class文件，生成虚拟机可以使用的类数据，并放入方法区；
	3. 最后进行链接
这里前两步其实对应java中的概念是：加载过程；后面是链接；最后还有一个初始化；
*/

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/

// ClassLoader 类加载器
type ClassLoader struct {
	cp *classpath.Classpath
	// classMap字段记录已经加载的类数据，key是类的完全限定名
	// 可以把classMap字段当作方法区的具体实现
	classMap map[string]*Class // loaded classes
}

//NewClassLoader // NewClassLoader()函数创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

/**
先查找classMap，看类是否已经被加载。如果是，直接返回类数据，否则调用loadNonArrayClass()方法加载类。
数组类和普通类有很大的不同，它的数据并不是来自class文件，而是由Java虚拟机在运行期间生成
*/

// LoadClass // 加载类信息
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}

	return self.loadNonArrayClass(name)
}

// loadNonArrayClass // 加载非数组的类，也是缓存中没有的
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// readClass 读取类信息
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	// 这里是吧 classFile传进来了
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5 类的定义
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 解析类
	class := parseClass(data)
	class.loader = self
	// 获取父类
	resolveSuperClass(class)
	// 获取接口
	resolveInterfaces(class)
	// 然后放在缓存中
	self.classMap[class.name] = class
	return class
}

// parseClass // 相当于解析之前的classFile成Class结构
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1 这里是获取父类信息，除java.lang.Object以外，所有的类都有且仅有一个超类。因此，除非是Object类，否则需要调用LoadClass()方法加载它的超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 获取接口，因为接口可能存在多实现，所以放在接口数组里面
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
