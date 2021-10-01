package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	// 帮助标签
	helpFlag bool
	// 版本信息
	versionFlag bool
	/**
	java命令提供了-verbose:class（简写为-verbose）选项，可以控制是否把类加载信息输出到控制台。
	也增加这样一个选项，另外参照这个选项增加一个-verbose:inst选项，用来控制是否把指令执行信息输出到控制台。
	*/
	verboseClassFlag bool
	verboseInstFlag  bool
	// 指令选项
	cpOption string
	// Java虚拟机将使用JDK的启动类路径来寻找和加载Java标准库中的类，所以这里是指定jre目录的位置
	XjreOption string
	// 需要处理的类
	class string
	// 参数
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "path to classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "path to classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

/**
打印使用的命令
*/
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
