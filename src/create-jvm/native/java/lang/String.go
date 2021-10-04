package lang

import (
	"create-jvm/native"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
// 如果字符串还没有入池，把它放入并返回该字符串，否则找到已入池字符串并返回
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
