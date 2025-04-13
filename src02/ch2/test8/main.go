// 匿名函数1

package main

import (
	"fmt"
)

func main() {
	fv := func() {
		fmt.Println("Hello World")
	}
	fv()

	// 这个地方，打印fv返回的是函数代码的入口地址，打印&fv是变量fv的地址，变量fv内存储的值是函数代码的入口地址
	fmt.Printf("fv的类型是%T, 打印输出fv的地址：%v\n", fv, &fv)
}
