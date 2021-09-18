package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}*/

// ConstantFieldRefInfo // 这个是字段的符号引用信息，其实就是 Class_info信息 + NameAndType信息
type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}*/

// ConstantMethodRefInfo // 这个是方法的符号引用信息，其实就是 Class_info信息 + NameAndType信息
type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

// ConstantInterfaceMethodRefInfo // 这个是接口的符号引用信息，其实就是 Class_info信息 + NameAndType信息
type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberRefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
