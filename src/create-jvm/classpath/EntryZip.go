package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry /** absPath字段存放ZIP或JAR文件的绝对路径
type ZipEntry struct {
	absPath string
}

// 这里就直接传文件所在的路径就行了
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 读取文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 从zip里面读取，相当于开启一个流
	r, err := zip.OpenReader(self.absPath)
	// 报错则返回报错信息，比如文件不存在
	if err != nil {
		return nil, nil, err
	}
	// 这里相当于 finally 里面去关闭流
	defer r.Close()
	// 遍历 zip 中可能存在的文件
	for _, f := range r.File {
		// 找到对应的字节码文件，则读出
		if f.Name == className {
			// 获取对应的文件流
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			// 导出字节数据
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

/**
返回路径字符串
*/
func (self *ZipEntry) String() string {
	return self.absPath
}
