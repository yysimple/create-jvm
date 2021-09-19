package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

// SourceFileAttribute //  SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件名
// 这里有点需要注意的是：attribute_length的值必须是2，虽然后面他只有一个 sourceFileIndex 属性，但是官方要求是 2
// The value of the attribute_length item of a SourceFile_attribute structure must be two.
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
