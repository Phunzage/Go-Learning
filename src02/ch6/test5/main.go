// 接口值的底层实现
package main

import "fmt"

type Any interface{}

type Dog struct {
}

// 接口值除了需要记录具体值（动态值）之外，还需要记录这个值属于的类型（动态类型）

func main() {
	var x Any

	// 此处接口存储的动态值为 “你好” ，动态值类型为字符串型
	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	// 此处接口存储的动态值为 100 ，动态值类型为int类型
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	// 此处接口存储的动态值为 true ，动态值类型为bool类型
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)

	// 此处接口存储的动态值为 nil ，动态值类型为结构体类型
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)
	// 注意此时接口变量与nil并不相等，因为它只是动态值的部分为nil，而动态类型部分保存着对应值的类型。
}
