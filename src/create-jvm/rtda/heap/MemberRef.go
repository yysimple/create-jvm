package heap

import "create-jvm/classfile"

type MemberRef struct {
	SymbolRef
	name       string
	descriptor string
}

// 这里就是类的字段信息和方法信息的公共类，转换相对应的类名和 方法/字段名 及其 描述（返回值类型+方法签名）
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
