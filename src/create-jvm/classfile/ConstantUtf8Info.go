package classfile

import "fmt"
import "unicode/utf16"

/*
官网地址：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.7
CONSTANT_Utf8_info {
    u1 tag;
    // 这个表示的是，接下来会有几个字符串常量
    u2 length;
	// 这里是每个字符串所占的字节数，但是这里的字符串的编码并不是标准的 UTF-8，
	而是 String content is encoded in modified UTF-8 简称 MUTF-8，可以参阅上面的 官网
    u1 bytes[length];
}
*/

/**
MUTF-8编码方式和UTF-8大致相同，但并不兼容。差别有两点：
	1. null字符（代码点U+0000）会被编码成2字节：0xC0、0x80；
	2. 补充字符（SupplementaryCharacters，代码点大于U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair）分别编码的
*/

// ConstantUtf8Info // 字符串 常量表 信息
type ConstantUtf8Info struct {
	str string
}

// readInfo // 读取信息
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	// 字符串的长度
	length := uint32(reader.readUint16())
	// 字符串对应的字节
	bytes := reader.readBytes(length)
	// 转成字符串
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) Str() string {
	return self.str
}

/*
func decodeMUTF8(bytes []byte) string {
	return string(bytes) // not correct!
}
*/

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
