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

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.3.1
// 这里解析过程就是按照官网来解析的，大致步骤可以看下官网
/**
要解析从D到用N表示的类或接口C的未解析符号引用，需要执行以下步骤：
	1. D的定义类加载器用于创建一个用n表示的类或接口。这个类或接口是c。因此，由于类或接口创建失败而引发的任何异常都可以由于类和接口解析失败而引发。
	2. 如果C是一个数组类，它的元素类型是引用类型，那么对表示元素类型的类或接口的符号引用将通过递归调用5.4.3.1中的算法来解析。
	3. 最后，检查C的访问权限。
		- 如果C对D不可访问(5.4.4)，类或接口解析将抛出IllegalAccessError
		- 例如，如果C是一个类，最初声明为public，但在D编译后被更改为非public，则会出现这种情况
如果第1步和第2步成功，但第3步失败，C仍然有效和可用。但是解析失败，D被禁止访问C。
*/
/**
通俗地讲，如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常。
Java虚拟机规范5.4.4节给出了类的访问控制规则，把这个规则翻译成Class结构体的isAccessibleTo()方法
*/
func (self *SymbolRef) resolveClassRef() {
	d := self.cp.class
	// 这里可以理解为，用 D 类（也就是SymbolRef里面存放的class）的类加载器去加载符号引用（className），得到一个类 C
	// 如果这里等得到则 1 成功了，数组先不考虑
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
