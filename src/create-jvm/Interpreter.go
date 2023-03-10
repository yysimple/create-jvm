package main

import (
	"create-jvm/instructions"
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"fmt"
)

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
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

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

// ---
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
// logInst????????????????????????????????????????????????????????????
//func interpret(method *heap.Method, logInst bool, args []string) {
//	// ?????????????????????????????????????????????????????????????????????
//	thread := rtda.NewThread()
//	// ??????????????????????????????
//	frame := thread.NewFrame(method)
//	// ??????
//	thread.PushFrame(frame)
//	jArgs := createArgsArray(method.Class().Loader(), args)
//	frame.LocalVars().SetRef(0, jArgs)
//
//	defer catchErr(thread)
//	loop(thread, logInst)
//}
//
//// interpret()???????????????startJVM()????????????????????????args?????????????????????createArgs-Array()?????????????????????Java????????????????????????????????????????????????????????????
//func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
//	stringClass := loader.LoadClass("java/lang/String")
//	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
//	jArgs := argsArr.Refs()
//	for i, arg := range args {
//		jArgs[i] = heap.JString(loader, arg)
//	}
//	return argsArr
//}
//
//// catchErr()???????????????????????????
//func catchErr(thread *rtda.Thread) {
//	if r := recover(); r != nil {
//		logFrames(thread)
//		panic(r)
//	}
//}
//
///**
//?????????????????????????????????????????????????????????pc???????????????????????????????????????????????????????????????????????????Java???????????????????????????????????????????????????????????????????????????
//*/
//func loop(thread *rtda.Thread, logInst bool) {
//	// ?????????????????????
//	reader := &base.BytecodeReader{}
//	for {
//		// ????????????????????????????????????????????????
//		frame := thread.CurrentFrame()
//		pc := frame.NextPC()
//		thread.SetPC(pc)
//
//		// decode
//		reader.Reset(frame.Method().Code(), pc)
//		// ???????????????
//		opcode := reader.ReadUint8()
//		inst := instructions.NewInstruction(opcode)
//		inst.FetchOperands(reader)
//		frame.SetNextPC(reader.PC())
//
//		if logInst {
//			logInstruction(frame, inst)
//		}
//
//		// execute
//		inst.Execute(frame)
//		if thread.IsStackEmpty() {
//			break
//		}
//	}
//}
//
//// logInstruction()????????????????????????????????????????????????
//func logInstruction(frame *rtda.Frame, inst base.Instruction) {
//	method := frame.Method()
//	className := method.Class().Name()
//	methodName := method.Name()
//	pc := frame.Thread().PC()
//	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
//}
//
//// logFrames()????????????Java??????????????????
//func logFrames(thread *rtda.Thread) {
//	for !thread.IsStackEmpty() {
//		frame := thread.PopFrame()
//		method := frame.Method()
//		className := method.Class().Name()
//		fmt.Printf(">> pc:%4d %v.%v%v \n",
//			frame.NextPC(), className, method.Name(), method.Descriptor())
//	}
//}
