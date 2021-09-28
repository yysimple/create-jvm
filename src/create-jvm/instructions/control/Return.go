package control

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

//RETURN  Return void from method
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	// 栈帧出栈的操作，这里会去更新栈顶
	frame.Thread().PopFrame()
}

// ARETURN Return reference from method
type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	// 拿到最新的栈顶元素
	invokerFrame := thread.TopFrame()
	// 移除当前栈帧所有的引用，然后出栈，这个引用会传递给调用的那个方法
	ref := currentFrame.OperandStack().PopRef()
	// 这里就拿到上个栈帧最后返回的结果，然后推入栈顶
	invokerFrame.OperandStack().PushRef(ref)
}

// Return double from method
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// Return float from method
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// Return int from method
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// Return double from method
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
