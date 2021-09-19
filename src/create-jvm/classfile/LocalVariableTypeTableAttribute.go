package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/

// LocalVariableTypeTableAttribute
// localvariableypetable属性与LocalVariableTable属性(4.7.13)的区别在于，它提供签名信息而不是描述符信息。
// 这种差异仅对其类型使用类型变量或参数化类型的变量才有意义。这样的变量将出现在两个表中，而其他类型的变量将只出现在LocalVariableTable中。
type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (self *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	self.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range self.localVariableTypeTable {
		self.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}
}
