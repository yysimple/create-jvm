package heap

import "create-jvm/classfile"

// FieldRef  字段引用
type FieldRef struct {
	MemberRef
	// field字段缓存解析后的字段指针
	field *Field
}

/**
在Java中，我们并不能在同一个类中定义名字相同，但类型不同的两个字段，那么字段符号引用为什么还要存放字段描述符呢？
答案是，这只是Java语言的限制，而不是Java虚拟机规范的限制。
也就是说，站在Java虚拟机的角度，一个类是完全可以有多个同名字段的，只要它们的类型互不相同就可以
*/

// 实例化字段引用
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

// ResolvedField //解析字段引用
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// 字段的符号引用解析成直接引用
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	// 这里依旧先拿到 类 c，然后在进行下一步操作
	c := self.ResolvedClass()
	// 解析字段引用
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

/**
如果类D想通过字段符号引用访问类C的某个字段，首先要解析符号引用得到类C，然后根据字段名和描述符查找字段。
如果字段查找失败，则虚拟机抛出NoSuchFieldError异常。如果查找成功，但D没有足够的权限访问该字段，则虚拟机抛出IllegalAccessError异常。
*/
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 从接口中找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 从父类中找
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
