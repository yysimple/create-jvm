package loads

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// AALOAD // Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

// Execute arrayref必须是类型引用的，并且必须引用其组件是类型引用的数组。索引必须为int类型。
// arrayref和index都是从操作数堆栈中弹出的。索引处数组组件中的引用值被检索并压入操作数堆栈。
func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 这里是数组索引
	index := stack.PopInt()
	// 拿到数组的引用信息
	arrRef := stack.PopRef()
	// 判断是否为空指针异常
	checkNotNil(arrRef)
	// 这里是拿到Object的数组
	refs := arrRef.Refs()
	// 校验索引是否越界
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
