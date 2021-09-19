package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/

/**
Signature属性是ClassFile、字段信息或方法信息结构(4.1、4.5、4.6)的属性表中的固定长度属性。
Signature属性记录类、接口、构造函数、方法或字段的签名(4.7.9.1)，这些类、接口、构造函数、方法或字段的声明在Java编程语言中使用类型变量或参数化类型。
*/

// SignatureAttribute // 用于处理泛型，可以看个例子
// eg： public <T> T getT(T s) { s.toString(); return s; }
// 然后其对应的是存储了 signature_index，指向字符串表信息对应的值是：<<T:Ljava/lang/Object;>(TT;)TT;>
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (self *SignatureAttribute) readInfo(reader *ClassReader) {
	self.signatureIndex = reader.readUint16()
}

func (self *SignatureAttribute) Signature() string {
	return self.cp.getUtf8(self.signatureIndex)
}
