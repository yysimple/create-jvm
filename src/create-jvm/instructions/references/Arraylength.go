package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// ARRAY_LENGTH // Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

// Execute arraylength指令只需要一个操作数，即从操作数栈顶弹出的数组引用。Execute()方法把数组长度推入操作数栈顶
func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	// 如果数组引用是null，则需要抛出NullPointerException异常，否则取数组长度，推入操作数栈顶即可
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
