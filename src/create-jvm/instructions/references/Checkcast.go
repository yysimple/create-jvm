package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// CHECK_CAST // Check whether object is of given type
// 这里是类型转换操作
type CHECK_CAST struct{ base.Index16Instruction }

// Execute //
// checkcast指令和instanceof指令很像，区别在于：instanceof指令会改变操作数栈（弹出对象引用，推入判断结果）;
// checkcast则不改变操作数栈（如果判断失败，直接抛出ClassCastException异常）
func (self *CHECK_CAST) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()

	/**
	先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态。
	32 aload_3
	33 checkcast #15 <com/simple/java/SubOrder>
	36 astore 4  -- 这里的4 是因为生成了一个新的局部变量，所以放在新的位置上，但是中间发生的 转换 操作是不会影响操作数栈的状态的
	如果引用是null，则指令执行结束。也就是说，null引用可以转换成任何类型，否则解析类符号引用，判断对象是否是类的实例。
	如果是的话，指令执行结束，否则抛出ClassCastException。instanceof和checkcast指令一般都是配合使用的
	*/
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().RtConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
