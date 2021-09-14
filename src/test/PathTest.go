package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 这里的src可以理解为是等同java里面的，代码的工作目录
	absDir, _ := filepath.Abs("src/test/PathTest.go")
	fmt.Println("absDir =", absDir)
	// info := fs.FileInfo.IsDir("src/test/PathTest.go")
	// fmt.Println("info =", info)
	jh := os.Getenv("JAVA_HOME")
	fmt.Println("jh: ", jh)

	if _, err := os.Stat("D:\\java\\jdk1;D:\\java\\jdk;"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("dir not exist")
		}
	}

}
