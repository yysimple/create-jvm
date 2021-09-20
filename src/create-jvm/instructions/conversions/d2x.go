package conversions

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

/**
其实这里把D转换成其他的操作全部放在一起了，
还可以从 宽化 窄化 角度来分析，可以参考我的文章：https://www.wolai.com/ax9WjQNCUaKPqeNtLWf5fH
*/

// D2F Convert double to float
// 由于GO提供了很多类型转化的方法，所以实现起来很简单
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

//D2I Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// D2L Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
