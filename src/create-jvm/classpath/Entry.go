package classpath

import (
	"os"
	"strings"
)

// :(linux/unix) or ;(windows)这里是用来读取系统的分割符的
const pathListSeparator = string(os.PathListSeparator)

// Entry /*  实现类路径：可以把类路径想象成一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径构成。三个小路径又分别由更小的路径构成*/
type Entry interface {
	// className: fully/qualified/ClassName.class
	// 比如要读取java.lang.Object类，传入的参数应该是java/lang/Object.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	// 多个路径的组合 D:\golang\source\bin\create-jvm\test\PathTest.go --（;）以这个分割符做分割--E:\golang\source\bin\create-jvm\test\PathTest.go
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// 这里是判断如果是jar包或者zip包，都以这种读取方式去加载文件
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
