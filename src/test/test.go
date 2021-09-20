package main

import (
	"flag"
	"fmt"
	"os"
)

var help bool
var getIp bool

func main() {
	flag.Usage = printUsage
	flag.BoolVar(&help, "help", false, "help message for")
	flag.BoolVar(&getIp, "getIp", false, "help message for")
	flag.Parse()
	if getIp {
		fmt.Println("i'm getIp", getIp)
		printUsage()
	} else if help {
		fmt.Println("i'm help", help)
		printUsage()
	}

	bitMove()
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}

func bitMove() {
	s := uint32(2) & 0x1f
	fmt.Println("16: ", 0x1f) // 00011111
	fmt.Println("16: ", 0x3f) // 00111111
	fmt.Println("s: ", s)
	result := 1 << s
	fmt.Println("result: ", result)
}
