// fmt包通过反射检查传入的值是否实现了 fmt.Stringer 接口。

package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	stu := student{
		name: "tom",
		age:  20,
	}
	fmt.Println(&stu)
}

func (stu *student) String() string {
	str := fmt.Sprintf("name = [%v], age = [%v]\n", stu.name, stu.age)
	return str
}
