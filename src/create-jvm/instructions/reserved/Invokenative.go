package reserved

import (
	"create-jvm/instructions/base"
	"create-jvm/native"
	"create-jvm/rtda"
)

import _ "create-jvm/native/java/lang"
import _ "create-jvm/native/sun/misc"

// INVOKE_NATIVE // Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	// 根据类名、方法名和方法描述符从本地方法注册表中查找本地方法实现，如果找不到，则抛出UnsatisfiedLinkError异常，否则直接调用本地方法
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
