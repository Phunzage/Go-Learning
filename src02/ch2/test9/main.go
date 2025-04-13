// 匿名函数2

package main

import (
	"fmt"
)

func main() {

	// 匿名函数
	// 声明后立即执行
	func(a, b int) {
		fmt.Println("hello我是匿名函数, 两值相加为:", a+b)
	}(3, 5)
	fmt.Println("--------")
	// 为匿名函数赋变量
	app := func(a, b int) {
		fmt.Println("有变量赋值的匿名函数, 两值相加为:", a+b)
	}
	// 可以任意时刻调用
	app(5, 10)

}
