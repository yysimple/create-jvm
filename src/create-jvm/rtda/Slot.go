package rtda

// Slot // 实际上再jvm虚拟机自己的实现里，也是使用slot的概念，double和long都是占两个slot
// 其中每个slot占四个字节，也就是32位，然后争对于对象类型，也即是引用类型，则存放的是引用，只占一个slot
type Slot struct {
	num int32
	ref *Object
}
