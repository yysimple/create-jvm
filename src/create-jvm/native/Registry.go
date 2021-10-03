package native

import "create-jvm/rtda"

// NativeMethod // 本地方法
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

// Register // 注册到registry变量中，其是个哈希表，值是具体的本地方法实现
// 类名、方法名和方法描述符加在一起才能唯一确定一个方法，所以把它们的组合作为本地方法注册表的键，Register()函数把前述三种信息和本地方法实现关联起来
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// FindNativeMethod // 这个是查找本地方法
// FindNativeMethod()方法根据类名、方法名和方法描述符查找本地方法实现，如果找不到，则返回nil
// java.lang.Object等类是通过一个叫作registerNatives()的本地方法来注册其他本地方法的
// 像registerNatives()这样的方法就没有太大的用处。为了避免重复代码，这里统一处理，如果遇到这样的本地方法，就返回一个空的实现
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}
