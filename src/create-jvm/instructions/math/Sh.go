package math

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// ISHL // Shift left int
type ISHL struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 这两部操作是基于前一步 的入栈操作，所以栈顶元素其实是 a << 2 这里面的 2
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 这里 0x1f 是 31 int变量只有32位，所以只取v2的前5个比特就足够表示位移位数了
	// Go语言位移操作符右侧必须是无符号整数，所以需要对v2进行类型转换
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// ISHR // Arithmetic shift right int
// 这里是有符号的位运算
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// IUSHR // Logical shift right int
// 无符号运算
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

//LSHL // Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// long变量有64位，所以取v2的前6个比特
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// LSHR // Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//LUSHR // Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
