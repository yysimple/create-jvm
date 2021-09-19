package main

import (
	"create-jvm/classfile"
	"create-jvm/classpath"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

// 启动vm去加载类
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

/**
加载对应的类，然后将对应的二进制文件解析成对应的字节码文件格式
*/
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	// 这里是将class文件装成字节类型
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	// 这里就是解析过程
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

/**
打印一些信息：类似于idea中 使用 javap 指令去解析 class文件，然后输出再控制台的那种方式
*/
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
