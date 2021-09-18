package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/

// ConstantStringInfo // 这里虽然是 String 常量表，但是其并不会存实际的 字符串
// 什么意思呢？ 比如再代码里面定义了这样的常量：String name = "张三";
// 这个“张三”是会存到 utf-8 字符串常量表中的 ， String_info中只会存一个 指向对应字符串的索引
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// 读取索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

//
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
