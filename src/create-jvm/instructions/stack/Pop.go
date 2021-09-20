package stack

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// POP // Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/**

POP：将栈顶元素出栈

| d | <-
| c |    	| c |
| b |   ->	| b |
| a |    	| a |

*/

// Execute // 大致就是上面的操作过程
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// POP2 // Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/**

POP2：将最上面两个元素出栈

| d | <-
| c | <-
| b |   ->	| b |
| a |    	| a |

*/

// Execute // 大致就是上面的操作过程
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
