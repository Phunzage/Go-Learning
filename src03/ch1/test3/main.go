// 输入输出与err

package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入你的名字：")
	fmt.Scan(&name)
	fmt.Println("你的名字是：", name)

	// 加入err
	var age int
	_, err := fmt.Scan(&age)
	fmt.Println(age, err)
}
