package base

import "create-jvm/rtda"

// Branch // 用于条件判断跳转
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
