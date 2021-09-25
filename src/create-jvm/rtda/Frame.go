package rtda

import "create-jvm/rtda/heap"

// Frame // 栈帧
type Frame struct {
	// 这里其实就是维护一个链表，指向下一个节点的指针
	lower *Frame // stack is implemented as linked list
	// 局部变量表
	localVars LocalVars
	// 操作数栈
	operandStack *OperandStack
	// 当前线程
	thread *Thread
	// 放入方法信息
	method *heap.Method
	// the next instruction after the call
	nextPC int
}

// NewFrame // 新建栈帧操作，初始化
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// newFrame // 这里的话是再初始化的时候，指定是哪个线程
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
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

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
