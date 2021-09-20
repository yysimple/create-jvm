package comparisons

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// LCMP // Compare long
// 这里就是比较两个数值的大小，转成指定的 1 0 -1 然后入栈
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
