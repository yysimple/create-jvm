package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// INVOKE_INTERFACE // Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

// FetchOperands // 注意，和其他三条方法调用指令略有不同，在字节码中，invokeinterface指令的操作码后面跟着4字节而非2字节
// 前两字节的含义和其他指令相同，是个uint16运行时常量池索引。
// 第3字节的值是给方法传递参数需要的slot数，其含义和给Method结构体定义的argSlotCount字段相同。
// 正如我们所知，这个数是可以根据方法描述符计算出来的，它的存在仅仅是因为历史原因。
// 第4字节是留给Oracle的某些Java虚拟机实现用的，它的值必须是0。该字节的存在是为了保证Java虚拟机可以向后兼容。
func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().RtConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

/**
invokestatic指令调用静态方法，很好理解。
invokespecial指令也比较好理解。
首先，因为私有方法和构造函数不需要动态绑定，所以invokespecial指令可以加快方法调用速度。
其次，使用super关键字调用超类中的方法不能使用invokevirtual指令，否则会陷入无限循环。
那么为什么要单独定义invokeinterface指令呢？统一使用invokevirtual指令不行吗？答案是，可以，但是可能会影响效率。
这两条指令的区别在于：当Java虚拟机通过invokevirtual调用方法时，this引用指向某个类（或其子类）的实例。
因为类的继承层次是固定的，所以虚拟机可以使用一种叫作vtable（Virtual MethodTable）的技术加速方法查找。
但是当通过invokeinterface指令调用接口方法时，因为this引用可以指向任何实现了该接口的类的实例，所以无法使用vtable技术。
*/
