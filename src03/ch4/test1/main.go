// 匿名结构体
// 没有显式类型名称的结构体，通常用于临时或一次性的数据组合场景。
package main

import "fmt"

//方式一
var user struct {
	name string
	age  int8
}

//方式二
var dog = struct {
	name  string
	age   int8
	color string
}{
	name:  "lucky",
	age:   3,
	color: "black",
}

func main() {
	user.name = "lisi"
	user.age = 18
	fmt.Printf("%#v\n", user)

	fmt.Println(dog)
}
