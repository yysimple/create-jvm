package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// PUT_FIELD // Set field in object,其字指令本身只有 一个指向常量池的索引
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.putfield
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	// 获取当前栈帧对应的方法
	currentMethod := frame.Method()
	// 获取对应的类信息
	currentClass := currentMethod.Class()
	// 获取运行时常量池信息
	cp := currentClass.RtConstantPool()
	// 获取对应的索引信息上的字段引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段
	field := fieldRef.ResolvedField()

	// 如果是静态的则抛出异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是final修饰的，看是不是在构造器也即初始化方法中，否则抛出异常
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		// 这里先拿到栈顶的引用
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		// 为对象里的实例变量赋值
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}
