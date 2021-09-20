package loads

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// ILOAD // 从常量池指定索引位位置获取对应的值，结合操作数一起使用
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

// 从局部变量表中指定索引位值获取值，然后推入到栈顶
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
