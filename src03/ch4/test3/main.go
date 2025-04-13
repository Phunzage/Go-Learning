//结构体嵌套

package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name    string
	gender  string
	age     int8
	address address
}

func main() {
	p1 := person{
		name:   "张三",
		gender: "男",
		age:    20,
		address: address{
			province: "山东",
			city:     "德州",
		},
	}

	fmt.Println(p1)
}
