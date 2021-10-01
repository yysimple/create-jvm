package rtda

// Stack // jvm stack
type Stack struct {
	// 这是设置最大的大小，之后可以用来抛异常
	maxSize uint
	size    uint
	// 栈帧，这里是用链表实现的
	_top *Frame // stack is implemented as linked list
}

// newStack // 初始化栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 入栈操作
func (self *Stack) push(frame *Frame) {
	// 大于最大栈大小，这里先报错
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	// 栈顶存在元素，则将栈顶元素设置成当前栈帧的下一个节点
	if self._top != nil {
		frame.lower = self._top
	}

	// 当前栈帧设置成栈顶
	self._top = frame
	self.size++
}

// pop // 出栈操作
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	// 先取出栈顶
	top := self._top
	// 将栈顶元素 指向 其下一个节点
	self._top = top.lower
	// 将栈顶的下一个元素置空，也就是移除top节点
	top.lower = nil
	self.size--

	return top
}

// 这里是获取栈顶的值，但是不会出栈
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	return self._top
}

// 判断栈是否为空
func (self *Stack) isEmpty() bool {
	return self._top == nil
}
