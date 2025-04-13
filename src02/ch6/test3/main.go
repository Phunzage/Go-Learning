// 类型与接口的关系

// 一个类型可以同时实现多个接口
// 不同的类型还可以实现同一接口

package main

import "fmt"

type Mover interface {
	Move()
}

type Sayer interface {
	Say()
}

type Dog struct {
	Name string
}

type Car struct {
	Brand string
}

func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

func (d Dog) Say() {
	fmt.Printf("%s会叫\n", d.Name)
}

func (c Car) Move() {
	fmt.Printf("%s飞驰\n", c.Brand)
}
func main() {
	var d = Dog{
		Name: "Lucky",
	}
	// 一个类型可以同时实现多个接口
	var m Mover = d
	var s Sayer = d
	m.Move()
	s.Say()

	// 不同的类型还可以实现同一接口
	m = Car{Brand: "宝马"}
	m.Move()
}
