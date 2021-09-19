package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/

// ExceptionsAttribute // 这里只是针对于 throws/throw 两种抛异常的方式，会生成异常信息
// 可以阅读我的文章：https://www.wolai.com/ax9WjQNCUaKPqeNtLWf5fH
type ExceptionsAttribute struct {
	// 这里是以表格形式进行维护的，可能会抛出多个异常
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
