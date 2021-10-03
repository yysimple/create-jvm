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
	// extra字段用来记录Object结构体实例的额外信息
	extra interface{}
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

// GetRefVar // 对应的字段信息
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

// SetRefVar // 设置引用值，并放入到slot中
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
