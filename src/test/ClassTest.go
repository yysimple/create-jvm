package main

import "fmt"

type Person interface {
	talk()
}

type XiaoMing struct {
	Hobby string
}

func (self *XiaoMing) talk() {
	fmt.Println("我是小明")
}

type XiaoHong struct {
	Height float64
}

func (self *XiaoHong) talk() {
	fmt.Println("我是小红")
}

func main() {
	xm := &XiaoMing{"da lan qiu"}
	xh := &XiaoHong{160.0}
	testType(xm, xh)
}

func testType(persons ...Person) {
	for _, person := range persons {
		switch person.(type) {
		case *XiaoMing:
			person.talk()
			ming := person.(*XiaoMing)
			fmt.Println("我有兴趣爱好：", ming.Hobby)
		case *XiaoHong:
			person.talk()
			hong := person.(*XiaoHong)
			fmt.Println("我的身高是：", hong.Height)
		}
	}
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
