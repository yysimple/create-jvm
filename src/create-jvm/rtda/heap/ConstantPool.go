package heap

import (
	"create-jvm/classfile"
	"fmt"
)

/**
运行时常量池主要存放两类信息：字面量（literal）和符号引用（symbolicreference）。
字面量包括整数、浮点数和字符串字面量；
符号引用包括类符号引用、字段符号引用、方法符号引用和接口方法符号引用
*/

type Constant interface{}

// RtConstantPool  这里是运行时常量池
type RtConstantPool struct {
	class  *Class
	consts []Constant
}

// 将classFile中的常量转换成运行时常量池中的常量
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *RtConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &RtConstantPool{class, consts}

	// 争对不同常量进行转换
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}
	}

	return rtCp
}

//GetConstant 根据索引返回常量
func (self *RtConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
