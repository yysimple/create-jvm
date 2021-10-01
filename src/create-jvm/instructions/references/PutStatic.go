package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

//PUT_STATIC // Set static field in class
type PUT_STATIC struct{ base.Index16Instruction }

/**
putstatic指令给类的某个静态变量赋值，它需要两个操作数。
第一个操作数是uint16索引，来自字节码。通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用，解析这个符号引用就可以知道要给类的哪个静态变量赋值。
第二个操作数是要赋给静态变量的值，从操作数栈中弹出
*/

// Execute //
func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	// 获取方法信息
	currentMethod := frame.Method()
	// 获取方法对应的类信息
	currentClass := currentMethod.Class()
	// 获取运行时常量池
	cp := currentClass.RtConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段并拿到对应的类
	field := fieldRef.ResolvedField()
	class := field.Class()
	// todo: init class
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	// 如果不是静态变量，抛出异常，这里是jvm规定的异常信息
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 这里如果是final修饰的字段，且不是在 <clinit> 中进行初始化的
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 下面是获取字段的一系列信息：描述 / 对应的插槽id / 对象的插槽数组，最后用来确定位置 / 栈帧
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	// 根据描述符来判断对应的解析操作,解析和初始化操作其实是相关联的
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
