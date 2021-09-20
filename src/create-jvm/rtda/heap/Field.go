package heap

import "create-jvm/classfile"

// Field // 字段信息，存放字段相关的信息，继承于 ClassMember
type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

// newFields 新建一个字段信息
func newFields(class *Class, classFileFields []*classfile.MemberInfo) []*Field {
	// 初始化大小
	fields := make([]*Field, len(classFileFields))
	for i, classFileField := range classFileFields {
		fields[i] = &Field{}
		// 指定该字段属于哪个类的
		fields[i].class = class
		// 转换ClassFile中的基本信息
		fields[i].copyMemberInfo(classFileField)
		// 转换ClassFile中的属性信息
		fields[i].copyAttributes(classFileField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}
