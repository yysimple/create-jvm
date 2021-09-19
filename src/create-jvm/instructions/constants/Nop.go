package constants

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// NOP // 不做任何操作
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
