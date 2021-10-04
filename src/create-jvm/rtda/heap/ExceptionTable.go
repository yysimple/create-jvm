package heap

import "create-jvm/classfile"

type ExceptionTable []*ExceptionHandler

// ExceptionHandler 这个也是属性表里的一种类型：异常表
type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

// newExceptionTable()函数把class文件中的异常处理表转换成ExceptionTable类型。
// 有一点需要特别说明：异常处理项的catchType有可能是0。我们知道0是无效的常量池索引，但是在这里0并非表示catch-none，而是表示catch-all
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *RtConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

func getCatchType(index uint, cp *RtConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 这里注意两点：
// 第一，startPc给出的是try{}语句块的第一条指令，endPc给出的则是try{}语句块的下一条指令。
// 第二，如果catchType是nil（在class文件中是0），表示可以处理所有异常，这是用来实现finally子句的
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
