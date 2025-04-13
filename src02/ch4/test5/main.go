//结构体转换

package main

import "fmt"

type number struct {
	f float64
}
type nr number

func main() {
	//正确的命名/转换方式
	a := number{5.0}
	b := nr{5.0}
	c := number(b)

	//错误的命名/转换方式
	// var i float64 = b
	// var i = float64(b)
	// var i number = b
	fmt.Println(a, b, c)

}
