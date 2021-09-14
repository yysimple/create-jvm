package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry /** 定义一个类，里面只有 绝对路径
type DirEntry struct {
	absDir string
}

/**
这里相当于是自己手动创建了一个构造器，传入参数，然后返回初始化好的对象
*/
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// 这里通过绝对路径 + 文件所在的相对路径 + 文件名称
// 最终的目的是返回字节
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 反对对象里面的字段，以字符串的形式
func (self *DirEntry) String() string {
	return self.absDir
}
