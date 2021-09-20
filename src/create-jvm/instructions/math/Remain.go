package math

import (
	"create-jvm/instructions/base"
	"create-jvm/rtda"
	"math"
)

// DREM Remainder double
type DREM struct{ base.NoOperandsInstruction }

/**
Go语言没有给浮点数类型定义求余操作符，所以需要使用math包的Mod()函数。
另外，浮点数类型因为有Infinity（无穷大）值，所以即使是除零，也不会导致ArithmeticException异常抛出
*/

// Execute //
func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	// GO语言里面
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// FREM Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// IREM Remainder int
type IREM struct{ base.NoOperandsInstruction }

/**
先从操作数栈中弹出两个int变量，求余，然后把结果推入操作数栈。
这里注意一点，对int或long变量做除法和求余运算时，是有可能抛出ArithmeticException异常的
*/

// Execute //
func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// LREM // Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
