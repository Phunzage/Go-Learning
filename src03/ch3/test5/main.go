//值传递与引用传递

package main

import "fmt"

func main() {
	name := "xiaoming"

	// 值传递演示：传递 name 的副本
	value := valueFunc(name)
	fmt.Printf("[值传递] 原 name 的值：%s\n", name)      // 原值不变
	fmt.Printf("[值传递] 函数返回的 value：%s\n\n", value) // 返回新值

	// "指针传递"演示：传递 name 的指针（本质是传递指针的副本值）
	newName := pointerFunc(&name)
	fmt.Printf("[指针传递] 原 name 被修改为：%s\n", name)      // 原值被修改
	fmt.Printf("[指针传递] 函数返回的 newName：%s\n", newName) // 返回相同的新值
}

func valueFunc(string) string {
	name := "zhangsan" // 修改的是副本
	return name
}

func pointerFunc(name *string) string {
	*name = "lisi" // 通过指针修改原值
	return *name
}
