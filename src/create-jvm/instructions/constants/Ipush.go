package constants

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

/**
这里为什么不直接用 iconst 来表示 short的最大接受值呢? 可以看看我的文章：
https://www.wolai.com/ax9WjQNCUaKPqeNtLWf5fH
*/

// BIPUSH // Push byte
type BIPUSH struct {
	val int8
}

// FetchOperands // 这里会去读取其后面对应的操作数；if(字节码存在操作数)从字节码流中取出操作数；会走到这个流程
// 而这个操作数其实就是对应的常量值
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

// Execute // 将其获取到的值推入到操作数栈中
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH 这个原理其实跟上面是一样的
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
