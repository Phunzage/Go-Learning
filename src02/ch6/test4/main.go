// 空接口

// 空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
package main

import "fmt"

// Any 不包含任何方法的空接口类型
type Any interface{}

// Dog 狗结构体
type Dog struct{}

// 空接口的应用
// 1.空接口类型可以作为函数的参数（比如Println可以接收任意类型的变量输出）（常用）
// 2.空接口类型可以作为map的value

func main() {
	var x Any

	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)

	// 此处可以传任意类型的变量
	var m = make(map[string]interface{})
	m["name"] = "修章鱼"
	m["age"] = 20
	m["hobby"] = []string{"rua宝宝龙", "嘶溜"}
	fmt.Println(m)
}
