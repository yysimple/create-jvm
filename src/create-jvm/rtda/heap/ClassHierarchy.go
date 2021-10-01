package heap

// IsAssignableFrom  一些继承关系的判断
func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self

	// 也就是说，在三种情况下，S类型的引用值可以赋值给T类型：S和T是同一类型；T是类且S是T的子类；或者T是接口且S实现了T接口
	// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.instanceof
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// IsSubClassOf self extends c
// 这里是去判断 该类 是否是 other 的子类
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// IsImplements self implements iface
// 判断S是否是T的子类，实际上也就是判断T是否是S的（直接或间接）超类
func (self *Class) IsImplements(iface *Class) bool {
	// 从父类接口开始一直遍历下来
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// IsSubInterfaceOf self extends iface
// 判断S是否实现了T接口，就看S或S的（直接或间接）超类是否实现了某个接口T' ,T’要么是T，要么是T的子接口
func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// IsSuperClassOf // c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
