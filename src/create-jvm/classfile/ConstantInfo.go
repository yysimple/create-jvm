package classfile

// Constant pool tags，常量池中信息对应的 tag，这里都是10进制数据
// 官网地址：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
// 具体每项数据对应的含义，可参阅我的文章：https://www.wolai.com/fkvaPasLsWQ1ghzv9h46ho
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag; // 这个就是对应上面的数值，但是这里分析工具一般都是以 16 进制表示，需要转成 10 进制去找对应项
    u1 info[];
}
*/

// ConstantInfo // 常量表 信息，这里是个接口，具体读取过程交由各个实现类去完成，也就是 各个 实际常量表 去读
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// readConstantInfo // 读取对应 常量表 的信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// 读取对应的 tag 值
	tag := reader.readUint8()

	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// 这里就是
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
