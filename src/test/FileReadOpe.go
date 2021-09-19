package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{202, 254, 186, 190, 0, 0, 0, 52}
	bytes := readUint32(data)
	fmt.Println("bytes: ", bytes)
	readInt8()
}

func readUint32(data []byte) []byte {
	val := binary.BigEndian.Uint32(data)
	// true
	fmt.Println("is true? ", val == 0xCAFEBABE)
	data = data[4:]
	fmt.Println("val: ", val)
	return data
}

func readInt8() {
	// 00001010  00001111
	code := []byte{uint8(10), uint8(15)}
	ub := code[0]
	uc := code[1]
	fmt.Println("ub: ", ub)
	b := int8(ub)
	c := int8(uc)
	fmt.Println("test b | c: ", b<<8-c)
	b16 := uint16(ub)
	c16 := uint16(uc)
	fmt.Println("b16: ", b16, ", c16: ", c16)
	fmt.Println("test: ", b16<<8|c16)
	fmt.Println("int test: ", int16(2575))
}
