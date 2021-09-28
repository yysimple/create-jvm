package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// INVOKE_STATIC Invoke a class (static) method
// 用来静态方法，后面是获取符号引用的指向 #12
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	// 获取到运行时常量池
	cp := frame.Method().Class().RtConstantPool()
	// 获取到方法引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 获取到 method 对象
	resolvedMethod := methodRef.ResolvedMethod()
	// 如果不是静态方法，则报错
	/**
	M必须是静态方法，否则抛出Incompatible-ClassChangeError异常。
	M不能是类初始化方法。类初始化方法只能由Java虚拟机调用，不能使用invokestatic指令调用。
	这一规则由class文件验证器保证，这里不做检查。如果声明M的类还没有被初始化，则要先初始化该类
	*/
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取到类信息
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
