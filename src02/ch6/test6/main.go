// 类型断言

package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d Dog) Move() {
	fmt.Println("run~~")
}

func main() {

	var m Mover = &Dog{Name: "Lucky"}
	v, ok := m.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "David"
	} else {
		fmt.Println("类型断言失败")
	}

	fmt.Println(v.Name)
}
