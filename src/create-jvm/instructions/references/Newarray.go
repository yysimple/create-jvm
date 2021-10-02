package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

/**
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.newarray
这里数组操作指令是有固定格式的；
The count must be of type int. It is popped off the operand stack. The count represents the number of elements in the array to be created.
*/
const (
	//Array Type  atype
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// NEW_ARRAY Create new array
type NEW_ARRAY struct {
	atype uint8
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

// Execute newarray指令的第二个操作数是count，从操作数栈中弹出，表示数组长度
func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	// 如果count小于0，则抛出NegativeArraySizeException异常，否则根据atype值使用当前类的类加载器加载数组类，然后创建数组对象并推入操作数栈
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

// 判断数组类型
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}
