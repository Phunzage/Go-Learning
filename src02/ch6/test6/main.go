// 类型断言

package main

import (
	"fmt"
)

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

	Justify(m)
}

// 对传入的空接口类型变量x进行类型断言，以便进行后续操作
func Justify(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
