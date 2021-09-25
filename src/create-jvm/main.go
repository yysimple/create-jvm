package main

import (
	"create-jvm/classfile"
	"create-jvm/classpath"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 初始化一个类加载器
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

// 启动jvm去加载类
//func startJVM(cmd *Cmd) {
//	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	classFile := loadClass(className, cp)
//	// 从classFile中找到对应的方法表里面的属性表code里面的 main方法
//	mainMethod := getMainMethod(classFile)
//	if mainMethod != nil {
//		interpret(mainMethod)
//	} else {
//		fmt.Printf("Main method not found in class %s\n", cmd.class)
//	}
//}

// startJVM // 测试指令
/*func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}*/

/**
获取main方法
*/
func getMainMethod(classFile *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range classFile.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
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
	classFile, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return classFile
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

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
