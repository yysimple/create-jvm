package classfile

/*
官网：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

// AttributeInfo // 属性表
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes // 这里就是单纯的读取，跟之前读取常量池表操作差不多，但是这里是从索引位置 0 开始的
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	// 这里是根据 属性表计数器值 对应的长度去读取
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute 这里是读取单个属性表信息的
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 先读取出对应的属性名，然后再去官网找其对应的具体描述
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	// 长度字段也是通用的，表示属性的长度，这里可能会包含 字节码指令等信息
	attrLen := reader.readUint32()
	// 根据不同的属性名去读取不同的信息
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

//newAttributeInfo 对应不同类型的属性表信息
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		// 暂未解析的字段，使用通用格式表示一下
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
