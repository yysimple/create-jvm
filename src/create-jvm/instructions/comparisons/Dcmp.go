package comparisons

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

//DCMPG Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

/**
这里跟float是一样的到里，需要考虑特殊情况
*/
func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
