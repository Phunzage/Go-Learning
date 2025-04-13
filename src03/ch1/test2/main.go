package main

import (
	"fmt"
)

const (
	num1 = 13
	num2 = 15
)

func main() {
	// 查看是什么类型
	fmt.Printf("%T\t", "hello")
	// 打印空字符串
	fmt.Printf("%#v", "")

	// 将格式化之后的内容赋给一个变量
	var op = fmt.Sprintf("%d", num1)

	fmt.Println(op)
}
