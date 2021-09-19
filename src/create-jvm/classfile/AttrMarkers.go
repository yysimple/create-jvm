package classfile

/**
由于不包含任何数据，所以attribute_length的值必须是0。
Deprecated属性用于指出类、接口、字段或方法已经不建议使用，编译器等工具可以根据Deprecated属性输出警告信息。
J2SE 5.0之前可以使用Javadoc提供的deprecated标签指示编译器给类、接口、字段或方法添加Deprecated属性
eg：
@Deprecated
public void add(){}
*/

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

// DeprecatedAttribute // 这个属性 其实就是用来让jvm知道该方法已经过时了
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

// SyntheticAttribute // Synthetic属性用来标记源文件中不存在、由编译器生成的类成员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

// 这里不会去读取，只是用来标记的
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
