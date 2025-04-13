// map 和 struct vs new() 和 make()

// 通过 make 或字面量初始化底层数据：
// map / slices / channels

// 为什么这三者不使用new():
// 对于值类型new返回的是零值的指针，即有地址（分配了内存），地址里的元素是0（假如是int）。而对于引用类型，返回的零值是nil，这个时候指针指向nil，即没有对应的地址，这时候赋值就会有问题，故map等不用new()
package main

import "fmt"

type dog struct {
	name  string
	age   int8
	color string
}

func main() {
	//使用make声明map
	m := make(map[int]string)
	m[1] = "hello"
	m[2] = "map"
	fmt.Printf("%#v", m)

	fmt.Println()

	//使用new声明结构体
	a := new(dog)
	a.name = "lucky"
	a.age = 3
	a.color = "black"
	fmt.Printf("%#v", a)
}
