package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// CodeAttribute // 这个属性只会存在方法属性中 method_info 中
// 里面包含了，最大栈深，变量个数等等，其解析过程还是挺多的
// 有兴趣的可以看看我写的文章：https://www.wolai.com/wwehi8W6Pckpy98cnhvPFt
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

// 这里就是读取过程
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

// MaxStack 下面这些就是获取其对应的值
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

// MaxLocals // 获取变量的数量
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

// Code 获取对应字节码指令
func (self *CodeAttribute) Code() []byte {
	return self.code
}

// ExceptionTable // 获取异常表
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

// ExceptionTableEntry // 异常表信息
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// 读取异常信息，这里的异常一般指的是 try-catch-finally 这种模式，会再 Code中维护一个异常表信息
// 比如： try{ xxx } catch (FileNotFoundException e) { e.printStackTrace();}
// 这种情况下就会再code中维护一个异常表信息 ： 类似于：
// Nr.	startPC	 endPC	handlerPC	catchType
//  0	   0	   22	   25	 	cp_info #14(java/io/FileNotFoundException)
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
