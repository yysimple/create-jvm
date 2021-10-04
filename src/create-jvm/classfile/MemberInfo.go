package classfile

/*
u2 fields_count;
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}

u2 methods_count;
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// MemberInfo /** 用来记录字段表信息和方法表信息；因为这两个表信息都是通用的字段 */
// “表”，“表”，“表”，“表”，“表” /// 强调 -- 这里就是表信息，除了cp之外，其他的都是 jvm 对字节码的规范约束
type MemberInfo struct {
	// cp字段保存常量池指针，后面会用到它
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// readMembers()读取字段表或方法表, 这里之后就可以
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 这里是对应的属性表和方法表的计数器；u2 fields_count; 两个字节
	memberCount := reader.readUint16()
	// 这里初始化“表信息”，长度就是前面的计数器的大小
	members := make([]*MemberInfo, memberCount)
	// 遍历对应索引的member；也就是“表”的基本信息，并将对应的字节转成其属性
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

/**
读取字节操作，转成对应的属性；
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	// 最终返回 通用 “表信息”对象
	return &MemberInfo{
		cp: cp,
		// 访问标识，两个字节
		accessFlags: reader.readUint16(),
		// 名称索引，两个字节
		nameIndex: reader.readUint16(),
		// 字段/方法描述，对应的是返回值类型，两个字节
		descriptorIndex: reader.readUint16(),
		// 字段表/方法表里面的属性
		attributes: readAttributes(reader, cp),
	}
}

//AccessFlags // 类似get方法
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// Name //  cp （符号引用），指向常量池表中的位置，可以得到对应的name（比如一个方法，add(),然后这个指向的就是 add字符串 ）
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Descriptor // 这个就是对应的常量池里面返回值的描述，比如 void add1()；那么这个描述在常量池中对应位置处，是 ()V 这写字符串
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

// CodeAttribute 获取方法里的code属性
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

// ConstantValueAttribute // 这里是获取常量表达式的值
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

// ExceptionsAttribute 解析异常表信息
func (self *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (self *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (self *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return self.getUnparsedAttributeData("AnnotationDefault")
}

func (self *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
