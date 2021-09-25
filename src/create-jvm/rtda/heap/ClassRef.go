package heap

import "create-jvm/classfile"

// ClassRef 类的引用直接继承即可,一般情况下只需要知道类型就可以找到指向了
type ClassRef struct {
	SymbolRef
}

// newClassRef 从常量池的类信息表里拿到类名
func newClassRef(cp *RtConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	// 拿到类信息的名称
	ref.className = classInfo.Name()
	return ref
}
