package lang

import (
	"create-jvm/native"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	//native.Register(jlClass, "isInterface", "()Z", isInterface)
	//native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	// getPrimitiveClass()是静态方法。先从局部变量表中拿到类名，这是个Java字符串，需要把它转成Go字符串
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	// 基本类型的类已经加载到了方法区中，直接调用类加载器的LoadClass()方法获取即可。最后，把类对象引用推入操作数栈顶
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
	// 首先从局部变量表中拿到this引用，这是一个类对象引用，通过Extra()方法可以获得与之对应的Class结构体指针
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	// 然后拿到类名，转成Java字符串并推入操作数栈顶。注意这里需要的是java.lang.Object这样的类名，而非java/lang/Object
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}
