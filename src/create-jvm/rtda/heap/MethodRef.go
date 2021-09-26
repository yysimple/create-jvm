package heap

import "create-jvm/classfile"

// MethodRef // 跟字段引用是差不多的
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *RtConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

// ResolvedMethod // 这里传入方法引用，如果这里存在缓存，直接取method
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.3.3
// 这里其实就是按照官网意思一步一步解析
func (self *MethodRef) resolveMethodRef() {
	// 如果类D想通过方法符号引用访问类C的某个方法，先要解析符号引用得到类C
	d := self.cp.class
	c := self.ResolvedClass()
	// 如果C是接口，则抛出IncompatibleClassChangeError异常，否则根据方法名和描述符查找方法
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 如果找不到对应的方法，则抛出NoSuchMethodError异常
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 检查类D是否有权限访问该方法，如果没有，则抛出IllegalAccessError异常
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

// lookupMethod 先从C的继承层次中找，如果找不到，就去C的接口中找
func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先从该类中去找
	method := LookupMethodInClass(class, name, descriptor)
	// 没找到，就往其接口中去找
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
