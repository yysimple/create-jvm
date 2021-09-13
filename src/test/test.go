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
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
