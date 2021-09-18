package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
	// 这里相当于是返回值类型对应的信息
	// 可以看看官网针对于类型返回的特定字符串信息：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.3.2
    u2 descriptor_index;
}
*/

/**
1. 描述符：
	①基本类型byte、short、char、int、long、float和double的描述符是单个字母，分别对应B、S、C、I、J、F和D。注意，long的描述符是J而不是L。
	②引用类型的描述符是L+类的完全限定名+分号。
	③数组类型的描述符是[+数组元素类型描述符。
2. 字段描述符就是字段类型的描述符；
3. 方法描述符是（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示；
比如 int add(int a, Long b) == (IJ)I  /  void sub() == ()V

正因为如此，所以支持方法的重载，我们根据 class_info 中的 name_index + NameAndType中的name_index + descriptor_index 对应的信息可以找到唯一的“方法”；
*/

// ConstantNameAndTypeInfo // 这里其实存的也是索引信息，可以通过对应的索引找到其 utf-8的字符串信息
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
