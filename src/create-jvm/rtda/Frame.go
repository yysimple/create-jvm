package rtda

// Frame // 栈帧
type Frame struct {
	// 这里其实就是维护一个链表，指向下一个节点的指针
	lower *Frame // stack is implemented as linked list
	// 局部变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	// todo
}

// NewFrame // 新建栈帧操作，初始化
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// LocalVars // get set
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
