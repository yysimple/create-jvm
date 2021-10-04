package heap

import "create-jvm/classfile"

// Method // 这里就是类中方法的信息
type Method struct {
	ClassMember
	// 这个编译器就能确定
	maxStack  uint
	maxLocals uint
	// 字节码指令操作
	code []byte
	// 异常表信息
	exceptionTable ExceptionTable
	// 异常属性
	exceptions *classfile.ExceptionsAttribute
	// 记录该方法有多少个参数
	argSlotCount uint
	// 方法的描述信息
	parsedDescriptor *MethodDescriptor
	// 方法执行的行号
	lineNumberTable *classfile.LineNumberTableAttribute
	// 参数信息
	parameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	// 注解信息
	annotationDefaultData []byte // AnnotationDefault_attribute
}

// 初始化一个方法信息，也是从classFile转换过来
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

// 详细初始化每个方法，包括native方法
func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	// 先计算argSlotCount字段，如果是本地方法，则注入字节码和其他信息
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

// 本地方法在class文件中没有Code属性，所以需要给maxStack和maxLocals字段赋值
func (self *Method) injectCodeAttribute(returnType string) {
	// 本地方法帧的操作数栈至少要能容纳返回值，为了简化代码，暂时给maxStack字段赋值为4
	self.maxStack = 4 // todo
	// 因为本地方法帧的局部变量表只用来存放参数值，所以把argSlotCount赋给maxLocals字段刚好
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}

// 这里是去解析方法中的参数占用插槽的个数
func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++ // `this` reference
	}
}

// 转换属性信息
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			self.class.rtConstantPool)
	}
	self.exceptions = cfMethod.ExceptionsAttribute()
	self.annotationData = cfMethod.RuntimeVisibleAnnotationsAttributeData()
	self.parameterAnnotationData = cfMethod.RuntimeVisibleParameterAnnotationsAttributeData()
	self.annotationDefaultData = cfMethod.AnnotationDefaultAttributeData()
}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

// getters
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) AnnotationDefaultData() []byte {
	return self.annotationDefaultData
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

// FindExceptionHandler
// FindExceptionHandler()方法调用ExceptionTable.findExceptionHandler()方法搜索异常处理表，如果能够找到对应的异常处理项，则返回它的handlerPc字段，否则返回-1
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}

// 判断是否是构造器
func (self *Method) isConstructor() bool {
	return !self.IsStatic() && self.name == "<init>"
}

//ParameterTypes  获取所有的参数信息 reflection
func (self *Method) ParameterTypes() []*Class {
	if self.argSlotCount == 0 {
		return nil
	}

	paramTypes := self.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = self.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}

// ExceptionTypes 异常类型
func (self *Method) ExceptionTypes() []*Class {
	if self.exceptions == nil {
		return nil
	}

	exIndexTable := self.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := self.class.rtConstantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}

// ReturnType 获取方法的返回值信息
func (self *Method) ReturnType() *Class {
	returnType := self.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return self.class.loader.LoadClass(returnClassName)
}

//ParameterAnnotationData 获取参数信息
func (self *Method) ParameterAnnotationData() []byte {
	return self.parameterAnnotationData
}

// 是否是静态变量的初始化操作
func (self *Method) isClinit() bool {
	return self.IsStatic() && self.name == "<clinit>"
}
