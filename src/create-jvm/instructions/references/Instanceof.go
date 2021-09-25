package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// INSTANCE_OF // Determine if object is of given type
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.instanceof
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	// 先弹出对象引用，如果是null，则把0推入操作数栈。
	// 用Java代码解释就是，如果引用obj是null的话，不管ClassYYY是哪种类型，都返回false：
	// If objectref is null, the instanceof instruction pushes an int result of 0 as an int on the operand stack.
	if ref == nil {
		stack.PushInt(0)
		return
	}

	// 如果对象引用不是null，则解析类符号引用，判断对象是否是类的实例，然后把判断结果推入操作数栈。
	cp := frame.Method().Class().RtConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
