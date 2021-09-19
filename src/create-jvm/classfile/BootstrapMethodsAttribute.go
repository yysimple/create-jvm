package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/

// BootstrapMethodsAttribute // 这个BootstrapMethods属性，用于支持新增的invokedynamic指令 是java7才有的
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (self *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	self.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range self.bootstrapMethods {
		self.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
