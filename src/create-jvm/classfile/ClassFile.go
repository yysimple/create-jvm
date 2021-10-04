package classfile

import "fmt"

// ClassFile /** @link  https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html /
// jdk 官网
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// ClassFile /** 定义字节码文件格式 /
type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse /* 解析字节码 字节形式，转成 ClassFile 格式 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 这里大概得意思是防止宕机处理
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	// 这里把数据交给 ClassReader
	cr := &ClassReader{classData}
	// 初始化一个 ClassFile 类
	cf = &ClassFile{}
	// 放在 read方法里面处理
	cf.read(cr)
	return
}

/**
读取字节数据到ClassFile类中
这里有几点需要注意：
	1. 初始化的值都放在了 reader 中
	2. 因为 reader 中所有读取操作都是将前面对应的字节跳过的，所以
*/
func (self *ClassFile) read(reader *ClassReader) {
	// 读取魔数
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 1. 读取魔数  u4 magic;
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	// 这个 magic 读取出来对应的二进制是：3405691582（如果在GoLand编辑器的控制台就是这个值）
	// 可以使用在线进制转换或者计算器，计算其对应的16进制 cafebabe
	magic := reader.readUint32()
	// 每个字节码文件必须是以这4个字节的魔数来开头对应的，就拿之前的Object来看看
	// 前四位是：202 254 186 190，这里需要注意的是，前面通过 ioutil.ReadFile读出来的数据，其实是二进制的，但是在idea的控制台中会变成 10 进制的值
	// 然后转换一下，每个数对应 8位1字节：11001010 11111110 10111010 10111110 转成对应的16进制就是：CA FE BA BE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 2. 读取版本号
// u2 minor_version; 次版本号，这个在java 1.2之后就没怎么使用了
// u2 major_version; 主版本，每次有大版本都会+1；
// java1.2 ~ java1.8 对应的版本号是：46 ~ 52
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 次版本，赋值
	self.minorVersion = reader.readUint16()
	// 主版本，赋值
	self.majorVersion = reader.readUint16()

	// 校验版本号是否正确
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

// MinorVersion /** ---- 这里往下，相当于是6个Get方法，对应java中的，且全部是共有的方法 ---- */
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// SourceFileAttribute 注意，并不是每个class文件中都有源文件信息，这个因编译时的编译器选项而异
func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
