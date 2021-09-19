package classfile

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/

// LocalVariableTableAttribute // 其实这个类型跟LineNumberTable是类似的
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc: reader.readUint16(),
			// 这个长度应该是其作用域范围
			length: reader.readUint16(),
			// 变量名
			nameIndex: reader.readUint16(),
			// 变量的描述，跟之前解析是一样的，NameAndType那里
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
