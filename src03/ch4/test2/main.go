//结构体初始化

package main

import "fmt"

type boy struct {
	name string
	age  int
}

//当结构体数据量较少，可以使用值传递，而当数据量很大时，使用指针传递已提高性能
func main() {
	b1 := boy{
		name: "zhangsan",
		age:  18,
	}
	//可以按顺序直接赋值
	b2 := boy{
		"lisi",
		20,
	}
	fmt.Printf("%#v,\n%#v", b1, b2)
}
