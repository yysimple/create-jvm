package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

// ConstantClassInfo // 这个就是跟之前 String_info 表类似，同样是通过索引位置，找到对应的 Utf-8_info 的字符串信息
// 类和超类索引，以及接口表中的接口索引指向的都是CONSTANT_Class_info常量
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
