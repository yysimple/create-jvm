package stack

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

// SWAP //  Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/

/**
SWAP: 交换最上面两个元素的位置

| d |		| c | <-
| c |		| d | <-
| b |	->	| b |
| a |		| a |

*/

//Execute // 大致流程
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
