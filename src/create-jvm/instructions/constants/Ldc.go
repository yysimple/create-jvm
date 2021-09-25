package constants

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// LDC // Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

// LDC_W // Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().RtConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// LDC2_W // Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().RtConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
