package heap

/**
因为4种类型的符号引用有一些共性，所以仍然使用继承来减少重复代码
所以定义一个父类：常量池符号引用
*/

//SymbolRef symbolic reference
// 对于类符号引用，只要有类名，就可以解析符号引用。
// 对于字段，首先要解析类符号引用得到类数据，然后用字段名和描述符查找字段数据。
// 方法符号引用的解析过程和字段符号引用类似。
type SymbolRef struct {
	// cp字段存放符号引用所在的运行时常量池指针，这样就可以通过符号引用访问到运行时常量池，进一步又可以访问到类数据
	cp *ConstantPool
	// className字段存放类的完全限定名。
	className string
	// class字段缓存解析后的类结构体指针，这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值
	class *Class
}

func (self *SymbolRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
func (self *SymbolRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
