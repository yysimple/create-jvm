package heap

import (
	"create-jvm/classfile"
	"strings"
)

/**
name、superClassName和interfaceNames字段分别存放类名、超类名和接口名。
注意这些类名都是完全限定名，具有java/lang/Object的形式。
constantPool字段存放运行时常量池指针，fields和methods字段分别存放字段表和方法表
*/

//Class // name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.2.1
// 我们也可以叫做他为类元信息，就是一个类的描述信息
type Class struct {
	// 访问标识
	accessFlags uint16
	// 当前类名称
	name string // thisClassName
	// 父类名称
	superClassName string
	// 接口名称，多实现，所以可能存在多个
	interfaceNames []string
	// 将常量池的内容引入
	rtConstantPool *RtConstantPool
	// 字段信息,会存在多个字段
	fields []*Field
	// 方法信息，也同样会存在多个
	methods []*Method
	// 加载该类的类加载器
	loader *ClassLoader
	// 父类指针
	superClass *Class
	// 对应的接口
	interfaces []*Class
	// 实例变量占用插槽数量
	instanceSlotCount uint
	// 类变量占用的插槽数量
	staticSlotCount uint
	// 插槽数组，之后用于判断变量的位置信息
	staticVars Slots
	// 用来判断是否开始初始化的
	initStarted bool
	// 类类型-通过jClass字段，每个Class结构体实例都与一个类对象关联
	jClass *Object
	// 属性信息
	sourceFile string
}

// newClass // 把ClassFile格式的数据转换成 class结构
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.rtConstantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

/**
下面是判断其访问标识
*/

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// Loader 返回类加载器
func (self *Class) Loader() *ClassLoader {
	return self.loader
}

// getters
func (self *Class) RtConstantPool() *RtConstantPool {
	return self.rtConstantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// jvms 5.4.4 https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.GetPackageName() == other.GetPackageName()
}

// GetPackageName 获取包名
func (self *Class) GetPackageName() string {
	// 这里拿到最后一个名字
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

//GetMainMethod //  获取到方法对象
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown" // todo
}

// GetClinitMethod 这是获取到静态方法初始化信息
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

// getters
func (self *Class) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *Class) Interfaces() []*Class {
	return self.interfaces
}

func (self *Class) Name() string {
	return self.name
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

// 下面都是判断是否是指定的方法
func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

// ArrayClass // 获取到数组的信息，并初始化
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

// getField // 获取对应的字段信息
func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}

//GetConstructors 获取所有的构造器
func (self *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}

// GetConstructor 返回构造器
func (self *Class) GetConstructor(descriptor string) *Method {
	return self.GetInstanceMethod("<init>", descriptor)
}

// GetFields // 反射获取所有的public字段信息
func (self *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(self.fields))
		for _, field := range self.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return self.fields
	}
}

// GetStaticMethod 获取当前类的静态方法
func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}

// GetMethods 获取方法信息
func (self *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}
