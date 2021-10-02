package references

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
)

// MULTI_ANEW_ARRAY // Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	// multianewarray指令的第一个操作数是个uint16索引，通过这个索引可以从运行时常量池中找到一个类符号引用，解析这个引用就可以得到多维数组类
	index uint16
	// 第二个操作数是个uint8整数，表示数组维度
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().RtConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	// popAndCheckCounts()函数从操作数栈中弹出n个int值，并且确保它们都大于等于0。
	// 如果其中任何一个小于0，则抛出NegativeArraySizeException异常
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	// 这里就是初始化每维数组
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			// 递归初始化，这里其实是处理每维数组里面的对象引用
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
