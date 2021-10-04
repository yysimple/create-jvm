package heap

import "create-jvm/classfile"

// ClassMember //用来存储类的字段信息个方法信息，对应转换classFile中的字段表和方法表
type ClassMember struct {
	accessFlags    uint16
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
	class          *Class
}

// copyMemberInfo // 进行转换操作
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Signature() string {
	return self.signature
}
func (self *ClassMember) AnnotationData() []byte {
	return self.annotationData
}

func (self *ClassMember) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// 用通俗的语言描述字段访问规则。如果字段是public，则任何类都可以访问。如果字段是protected，则只有子类和同一个包下的类可以访问。
// 如果字段有默认访问权限（非public，非protected，也非privated），则只有同一个包下的类可以访问。否则，字段是private，只有声明这个字段的类才能访问。
// 放在这里时因为方法引用到时候也差不多
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
