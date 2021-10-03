package base

import (
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// InvokeMethod 函数的前三行代码创建新的帧并推入Java虚拟机栈，剩下的代码传递参数
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	// 这里是去获取参数，这里是要去获取参数的数量，其实很好判断，在方法/字段的描述符里，这样的：(IL)J，所以IL就是参数
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			// == 这里的位置对应的跟方法是 类方法 还是 实例方法有关 ==
			/**
			这个数量并不一定等于从Java代码中看到的参数个数，原因有两个：
				第一：long和double类型的参数要占用两个位置。
				第二，对于实例方法，Java编译器会在参数列表的前面添加一个参数，这个隐藏的参数就是this引用。
			假设实际的参数占据n个位置，依次把这n个变量从调用者的操作数栈中弹出，放进被调用方法的局部变量表中，参数传递就完成了
			*/
			slot := invokerFrame.OperandStack().PopSlot()
			// 在代码中，并没有对long和double类型做特别处理。因为操作的是Slot结构体，所以这是没问题的，因为之前解析的时候也是解析成两个slot
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
