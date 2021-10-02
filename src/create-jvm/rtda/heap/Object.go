package heap

/**
Java虚拟机可以操作两类数据：基本类型（primitive type）和引用类型（reference type）。
*/

// Object // 模拟引用类型，先临时表示对象
type Object struct {
	// 存的当前类的指针信息
	class *Class
	// 存放实例变量
	// fields Slots
	data interface{}
}

/**
新创建一个对象
*/
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

// IsInstanceOf // 用来判断 class 类型是否是 self
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}
