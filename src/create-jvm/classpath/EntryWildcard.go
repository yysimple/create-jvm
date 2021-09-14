package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
这里是使用通配符 “*” 来和表示多个文件，所以其原理也是和 CompositeEntry 类似的
*/
func newWildcardEntry(path string) CompositeEntry {
	// remove *，这里是去掉 D:\golang\* 后面的 *
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	// 这里是创建一个函数，这里类似函数式编程，最后的实现都在这个函数里面
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果这个文件是目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			// 这里如果是jar包，则使用newZipEntry这个实现类去获取 entry
			jarEntry := newZipEntry(path)
			// 然后放到 组合Entry、中
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 这个其实就是会去循环找到目录下的文件，这里可以理解为是递归
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}

	return compositeEntry
}
