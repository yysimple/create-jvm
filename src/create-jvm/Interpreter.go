package main

import (
	"create-jvm/instructions"
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"create-jvm/rtda/heap"
	"fmt"
)

//func interpret(methodInfo *classfile.MemberInfo) {
//	codeAttr := methodInfo.CodeAttribute()
//	maxLocals := codeAttr.MaxLocals()
//	maxStack := codeAttr.MaxStack()
//	bytecode := codeAttr.Code()
//
//	thread := rtda.NewThread()
//	frame := thread.NewFrame(maxLocals, maxStack)
//	thread.PushFrame(frame)
//
//	defer catchErr(frame)
//	loop(thread, bytecode)
//}

/*func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}*/
// logInst参数控制是否把指令执行信息打印到控制台。
func interpret(method *heap.Method, logInst bool) {
	// 涉及到方法的调用，所以先初始化一个固定大小的栈
	thread := rtda.NewThread()
	// 在基于栈开始创建栈帧
	frame := thread.NewFrame(method)
	// 入栈
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

// catchErr()函数会打印出错信息
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

/**
在每次循环开始，先拿到当前帧，然后根据pc从当前方法中解码出一条指令。指令执行完毕之后，判断Java虚拟机栈中是否还有帧。如果没有则退出循环；否则继续
*/
func loop(thread *rtda.Thread, logInst bool) {
	// 读取字节码指令
	reader := &base.BytecodeReader{}
	for {
		// 获取到当前栈帧，然后一直往下读取
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		// 获取字节码
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

// logInstruction()函数在方法执行过程中打印指令信息
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

// logFrames()函数打印Java虚拟机栈信息
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
