package stack

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
)

/**
指令的详细用法参考我的文章：https://www.wolai.com/ax9WjQNCUaKPqeNtLWf5fH
*/

//DUP // Duplicate the top operand stack value
// 将栈顶元素复制一遍
type DUP struct{ base.NoOperandsInstruction }

/**

dup：下面的过程：其实就是将c复制一遍，这里说他再栈顶还是栈顶下一位都无所谓，应该是栈顶下一位

		     | c | <-
| c |        | c | <-
| b |   ->   | b |
| a |        | a |

*/

// Execute // 大致是上面的这个操作，将栈顶元素复制一遍的到两个 c
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// DUP_X1 // Duplicate the top operand stack value and insert two values down
// 带_x的指令是复制栈顶数据并插入栈顶以下的某个位置 _x1的算法是 1 + 1
type DUP_X1 struct{ base.NoOperandsInstruction }

/**

DUP_X1: 就是将栈顶元素复制一位，并放在其下面 1+1位置处

			| d | <-
| d |		| c |
| c |	->	| d | <-
| b |		| b |
| a |		| a |

*/

// Execute // 大致是上面的这个操作
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// DUP_X2 // Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }

/**

DUP_X2: 就是将栈顶元素复制一位，并放在其下面 1+2 位置处

			| d | <-
| d |		| c |
| c |	->	| b |
| b |		| d | <-
| a |		| a |

*/

// Execute // 大致是上面的这个操作
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// DUP2 // Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperandsInstruction }

/**

DUP2: 复制最上面的两个元素
			| d | <-
			| c | <-
| d |		| d | <-
| c |	->	| c | <-
| b |		| b |
| a |		| a |

*/

// Execute // 大致是上面的这个操作
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// DUP2_X1 // Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperandsInstruction }

/**

DUP2_X1: 复制最上面的两个元素,并往下移动一位
			| d | <-
			| c | <-
| d |		| b |
| c |	->	| d | <-
| b |		| c | <-
| a |		| a |

*/

// Execute // 大致是上面的这个操作
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// DUP2_X2 // Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperandsInstruction }

/**

DUP2_X2: 复制最上面的两个元素,并往下移动两位
			| d | <-
			| c | <-
| d |		| b |
| c |	->	| a |
| b |		| d | <-
| a |		| c | <-

*/

// Execute // 大致是上面的这个操作
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
