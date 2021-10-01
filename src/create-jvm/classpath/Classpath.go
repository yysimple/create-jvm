package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	// 如果是用户自己去指定加载的路径
	// 需要这样传：create-jvm.exe -Xjre D:\java\jdk\jre  -classpath D:\golang\source\bin\create-jvm\bin\java\ Demo
	// 或者不需要自己指定 jre 环境，会自动去系统变量中获取 ：create-jvm.exe -classpath D:\golang\source\bin\create-jvm\bin\java\ Demo
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析启动类路径和扩展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 这里是获取jdk环境，其实就是jre环境，去加载并执行class文件的时候，会用到jre中的lib的类库
	jreDir := getJreDir(jreOption)

	// jre/lib/* 这步操作就是去获取 jre下面的的 lib文件夹下面的所有 文件
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/* 这步操作就是去获取 jre下面的 lib 下面的 ext 文件夹下面的所有 文件
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该选项，则在当前目录下寻找jre目录。如果找不到，尝试使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	// 指令上面指定的路径去寻找
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 当前目录寻找
	if exists("./jre") {
		return "./jre"
	}
	// 这个是获取电脑的环境变量，也是 path 的值 D:\java\jdk
	if jdkAbsPath := os.Getenv("JAVA_HOME"); jdkAbsPath != "" {
		return filepath.Join(jdkAbsPath, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	// 校验该路径在系统上是否存在
	// D:\java\jdk1;D:\java\jdk; 这里的就是证明存在，D:\java\jdk 因为这个在path中是存在的，满足即可
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径。
ReadClass()方法依次从启动类路径、扩展类路径和用户类路径中搜索class文件
*/
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// ReadClass className: fully/qualified/ClassName：这里可以理解为 java中的全限定名
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		// fmt.Println("==== 从BootStrap加载器路径下读取 ====")
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		// fmt.Println("==== 从ext加载器路径下读取 ====")
		return data, entry, err
	}
	// 上面都没有加载到，则使用用户类路径去加载
	// fmt.Println("==== 从用户自定义加载器路径下读取 ====")
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
