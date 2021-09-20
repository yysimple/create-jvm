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
}

// 初始化一个方法信息，也是从classFile转换过来
func newMethods(class *Class, classFileMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(classFileMethods))
	for i, classFileMethod := range classFileMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(classFileMethod)
		methods[i].copyAttributes(classFileMethod)
	}
	return methods
}

// 转换属性信息
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
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
