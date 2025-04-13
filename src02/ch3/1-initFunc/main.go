// 流程执行顺序：
//1.全局变量定义，2.init函数，3.main函数

package main

import (
	"fmt"
)

var a int

func init() {
	a = 5
	fmt.Println("在init中a的值为", a)
}

func main() {

	a = 10
	fmt.Println("在main中a的值为", a)
}
