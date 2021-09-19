package base

type BytecodeReader struct {
	// 这里是读取到的字节码指令
	code []byte // bytecodes
	// 1 istore_1：这个前面的1 就是对应字节码的 pc
	pc int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// ReadUint8 读取对应的字节码指令，只读一位
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

// ReadUint16 读取连续两位自己的指令
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	// 这里很好理解，就是组成这样的数据：byte1 = 00010101 << 8  |  byte2 = 10101001 = 0001010100000000 | 0000000010101001 = 0001010110101001
	// 这样这里就是表示连续的两个字节的数据
	return (byte1 << 8) | byte2
}

// ReadInt32 连续读取四个字节的数据
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
