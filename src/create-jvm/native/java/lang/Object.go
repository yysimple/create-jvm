package lang

import (
	"create-jvm/native"
	"create-jvm/rtda"
	"unsafe"
)

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
	native.Register(jlObject, "notifyAll", "()V", notifyAll)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
	// 首先，从局部变量表中拿到this引用
	this := frame.LocalVars().GetThis()
	// 有了this引用后，通过Class()方法拿到它的Class结构体指针，进而又通过JClass()方法拿到它的类对象
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
// 如果类没有实现Cloneable接口，则抛出CloneNotSupportedException异常
// 否则调用Object结构体的Clone()方法克隆对象，然后把对象副本引用推入操作数栈顶
func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()

	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}

// public final native void notifyAll();
// ()V
func notifyAll(frame *rtda.Frame) {
	// todo
}
