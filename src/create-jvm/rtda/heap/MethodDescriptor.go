package heap

// MethodDescriptor 方法的描述信息，这里是将 (IL)J 这种描述信息 分成 参数 + 返回值
type MethodDescriptor struct {
	// 这里是参数列表
	parameterTypes []string
	// 返回值信息，java 只支持单返回
	returnType string
}

// addParameterType 解析参数
func (self *MethodDescriptor) addParameterType(t string) {
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}

	self.parameterTypes = append(self.parameterTypes, t)
}
