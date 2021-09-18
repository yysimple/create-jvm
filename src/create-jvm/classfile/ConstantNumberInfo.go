package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes; // int能存储的最大数值
}
*/

// ConstantIntegerInfo // 这里是存的有符号的数值，常量值，是按高位在前存储的int值
type ConstantIntegerInfo struct {
	val int32
}

// readInfo // 读取出对应的字节，然后转成 10 进制
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/

// ConstantFloatInfo // 读取 float的值
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

// ConstantLongInfo // 读取Long的值
// 这里的Long类型存储方式是：((long) high_bytes << 32) + low_bytes 分为高字节位个低字节位，先优先存高字节位，再存低字节位
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

// ConstantDoubleInfo // 读取double的值，这里跟Long类型是一样的 ((long) high_bytes << 32) + low_bytes 分为高字节位个低字节位，先优先存高字节位，再存低字节位
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
