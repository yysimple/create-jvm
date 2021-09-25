package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// NEW // Create new object
type NEW struct{ base.Index16Instruction }

/**
new指令的操作数是一个uint16索引，来自字节码。通过这个索引，可以从当前类的运行时常量池中找到一个类符号引用。
解析这个类符号引用，拿到类数据，然后创建对象，并把对象引用推入栈顶，new指令的工作就完成了
*/

// Execute // 执行具体的操作，这个命令执行的时候，一般会跟上一个 dup 指令
func (self *NEW) Execute(frame *rtda.Frame) {
	// 获取运行时常量池的信息
	cp := frame.Method().Class().ConstantPool()
	// 获取类引用信息
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析拿到类
	class := classRef.ResolvedClass()
	// todo: init class

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	// 将新建的对象推送到栈顶
	frame.OperandStack().PushRef(ref)
}
