package control

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// GOTO // Branch always
type GOTO struct{ base.BranchInstruction }

// Execute // 跳转到指定的offset对应的代码
func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
