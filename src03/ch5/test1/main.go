//方法定义：golang中的方法是作用在指定数据类型上的（即和指定数据类型绑定），因此自定义类型都可以有方法，而不仅仅是struct

//方法的调用与传参机制：
// 1.通过一个变量去调用方法时，其调用机制和函数一样
// 2.不一样的地方，变量调用方法时，该变量本身也会作为一个参数传递到方法（如果变量是值类型，则进行值拷贝，如果变量是引用类型，则进行地址拷贝）

//实现一个方法，输入name和age，将两者组合后输出

package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {

	p1 := person{
		"小明",
		18,
	}
	p2 := person{age: 20, name: "Tom"}
	p1.com()
	p2.com()
}

//方法
func (p person) com() {
	fmt.Printf("%s的年龄是%d\n", p.name, p.age)
}
