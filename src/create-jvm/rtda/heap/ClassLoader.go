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
	// 这里是把已经解析好的字节码文件 classfile文件信息传进来
	cp *classpath.Classpath
	// 用来标记是否需要打印类加载器加载类的过程
	verboseFlag bool
	// classMap字段记录已经加载的类数据，key是类的完全限定名
	// 可以把classMap字段当作方法区的具体实现
	classMap map[string]*Class // loaded classes
}

//NewClassLoader // NewClassLoader()函数创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
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
	// 验证阶段，这里比较麻烦，验证规则很多
	verification(class)
	// 准备阶段
	preparation(class)
	// 解析阶段
	resolution()
}

// verification // 验证阶段
func verification(class *Class) {
	// todo
	// 格式检查：魔数检查，版本检查等，这些在加载classFile里面已经进行了校验
	// 语义检查：是否继承了final类，是否由父类时抽象类，没去实现其方法
	// 字节码检查：跳转指令是否指向正确的位置，操作数类型是否合理，这里比如 goto 20 这个20位置是否存在，且是否合理？istore -21 这个 -21是否合法？
	// 符号引用验证：符号引用的直接引用是否存在？
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.2
// preparation // 准备阶段
/**
准备工作包括为类或接口创建静态字段，并将这些字段初始化为默认值(2.3,2.4)。这并不需要执行任何Java虚拟机代码;静态字段的显式初始化器是作为初始化的一部分执行的(5.5)，而不是准备。
	1. 准备阶段不包含 static final修饰的字段的内存分配和默认值初始化，因为final修饰的在编译的时候就分配了内存，准备阶段就会显示赋值了；
	2. 这里并不会为实例变量分配空间，其操作是跟对象一起分配到堆中的；而类变量这个操作是放在元空间的；
*/
func preparation(class *Class) {
	// 函数计算实例字段的个数，同时给它们编号
	calcInstanceFieldSlotIds(class)
	// 函数计算静态字段的个数，同时给它们编号
	calcStaticFieldSlotIds(class)
	// 给类变量分配空间，然后给它们赋予初始值
	allocAndInitStaticVars(class)
}

// 函数计算实例字段的个数，同时给它们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	// 这里是从父类开始，因为继承是可以使用父类的字段的
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	// 然后在拿自己的变量信息，遇到 double 则slot的索引+2
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

// 函数计算静态字段的个数，同时给它们编号
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

// 给类变量分配空间，然后给它们赋予初始值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	// 这里把所有静态变量全部取出来
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			// 初始化静态变量
			initStaticFinalVar(class, field)
		}
	}
}

// 为静态变量设置默认值，这里的初始化并不是显示初始化，真正进行显示初始化的时候是在 “初始化阶段”，而不是在准备阶段
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.rtConstantPool
	// 获取字段的定长属性
	cpIndex := field.ConstValueIndex()
	// 获取字段对应的插槽id
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

//
func resolution() {

}
