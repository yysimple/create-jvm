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
	code := []byte{0x0A, 0x0F}
	b := code[0]
	fmt.Println("b: ", b)
	i := int8(b)
	fmt.Println("i: ", i)
}
