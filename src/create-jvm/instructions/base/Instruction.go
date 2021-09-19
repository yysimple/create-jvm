package base

import "create-jvm/rtda"

/**
这里看个逻辑：
	do{
	  自动计算PC寄存器的值加1；
	  根据PC寄存器的指示位置，从字节码流中取出操作码；
	  if(字节码存在操作数)从字节码流中取出操作数；
	  执行操作码所定义的操作；
    } while(字节码长度>0) ；
我们可以把真个字节码的指令操作从逻辑上分成上面几个步骤，其实就两个：1. 读取字节码指令  2. 执行字节码指令
*/

// Instruction // 根据上面的逻辑，我们其实也不难将其抽象出来，做两件事情，也就是对应两个方法；
// 但是这里不管具体的实现；
type Instruction interface {
	// FetchOperands 这里就是去读取字节码指令
	FetchOperands(reader *BytecodeReader)
	// Execute 这就是去执行
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction // 这里就是不做任何操作，nop
type NoOperandsInstruction struct {
	// empty
}

//FetchOperands 这里就是就是去获取对应的指令信息
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction // 这个是分支跳转指令，Offset字段存放跳转偏移量  goto 15 (+6)，类似于这样的
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction // 用于读取一个字节的 操作数
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction // 用于读取两个字节的索引指向 比如指向常量池的索引，如 #12
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
