package classfile

import "fmt"

// ConstantPool // 常量池实际上是多个常量表组成的，ConstantInfo常量表其实里面又有tag，常量类型 info，具体可以参考官网，或者我的博客
// 我来博客地址：https://www.wolai.com/fkvaPasLsWQ1ghzv9h46ho
// 官网地址：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
type ConstantPool []ConstantInfo

/**
然后读取常量池的过程中，我们需要注意的是：
	1. 第一个字节码对应的是常量池的大小，但是这里有个问题，在jvm规范中，常量池是从索引位置1开始的
比如，常量池大小计数器：22，那么对应的常量表大小应该是 21项；0索引是无效索引，不会指向任何常量；
	2. 然后就是在常量池中有两种特殊的常量类型：CONSTANT_Long_info和CONSTANT_Double_info，这两者是会占据 两项 常量池大小，
所以最后如果存在这两种各一个，那么真正意思上的常量池项，只有 21 -1 -1 = 19项
官网：http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
*/
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	// 初始化大小为：常量池计数器数值 的数组，也就是常量表的数量
	cp := make([]ConstantInfo, cpCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	// 因为这里0号位索引是不需要读取的，也是不存在的，不会占据字节，所以从 1 号位开始读取
	for i := 1; i < cpCount; i++ {
		// 读取字节数据，到每个常量表中，这里是循环读，顺序也不能出错
		cp[i] = readConstantInfo(reader, cp)
		// 这里判断一下是否是上面说的那两种类型，如果是的话，那就占据常量池两项大小
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// getConstantInfo // 根据索引位置去拿到对应的常量值，可能是字符串之后
// 这里就是 符号引用  putfield #8 <com/atguigu/java/LoadAndStoreTest.aa> 这里的 #8 对应的就是常量池种的第八项
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))
}

// getNameAndType （参考 常量表 结构 - CONSTANT_NameAndType_info） // 从常量池查找字段或方法的名字和描述符
// 这里一般是通过上面的 #8 这种索引去找到对应的常量表位置。
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	// 转成字符串，这里一般就是名称信息比如方法
	// eg: 方法：private void add()
	// eg: 字段：private int num
	// 那么这个 name 就分别对应 add和num
	name := self.getUtf8(ntInfo.nameIndex)
	// 下面这个对应的就是返回值 ()V，I
	// 字段描述符表示类、实例或局部变量的类型
	// 可以看看官网针对于类型返回的特定字符串信息：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.3.2
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// getClassName （参考 常量表 结构 - CONSTANT_Class_info）// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 （参考 常量表 结构 - CONSTANT_uft8_info）// 从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	// 最后这个就是解析出来的字符串，比如方法名 add(), 那么这个 str就是 s t r 三个字节转成的字符串
	return utf8Info.str
}
