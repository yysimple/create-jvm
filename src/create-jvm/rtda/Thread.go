package rtda

/*
JVM
// 虚拟机栈是线程私有的，所以每个线程都会为自己创建一个私有的栈
  - Thread
  - pc
  - Stack
	  // 栈帧，可以理解为是每个方法调用形成的
      Frame001
        LocalVars
		// 操作数栈，用来处理中方法中间的一些调用计算操作等
        OperandStack
	  Frame002
        LocalVars
        OperandStack
*/

// Thread //
type Thread struct {
	// cpu线程上下文切换的时候，这里就会记录当前线程执行到了哪个位置
	pc int // the address of the instruction currently being executed
	// 这里传指针进来就是表示线程私有
	stack *Stack
	// todo
}

// NewThread // 初始化
func NewThread() *Thread {
	// 这里初始化操作，先默认初始化大小为1024
	return &Thread{
		stack: newStack(1024),
	}
}

// PC // get set方法
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

// PushFrame // 入栈操作
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// PopFrame // 出栈操作
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// CurrentFrame
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
