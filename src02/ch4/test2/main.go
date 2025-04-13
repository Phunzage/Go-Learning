//使用选项模式实例化结构体

package main

import (
	"fmt"
)

// 结构体声明
type Person struct {
	Name    string
	Age     int
	Address string
	Salary  float64
}

type PersonOptions func(p *Person)

//声明一个PersonOptions类型，它接受一个 *Person 类型的参数， 它必须是指针，因为我们要在闭包中对Person赋值

//创建选项函数，一般以With开头，它们的返回值就是一个闭包函数

func WithName(name string) PersonOptions {
	return func(p *Person) {
		p.Name = name
	}
}
func WithAge(age int) PersonOptions {
	return func(p *Person) {
		p.Age = age
	}
}
func WithAddress(address string) PersonOptions {
	return func(p *Person) {
		p.Address = address
	}
}
func WithSalary(salary float64) PersonOptions {
	return func(p *Person) {
		p.Salary = salary
	}
}

// 实际声明的构造函数签名如下，它接受一个可变长PersonOptions类型的参数。
func NewPerson(options ...PersonOptions) *Person {
	//优先应用options
	p := &Person{}
	for _, option := range options {
		option(p)
	}

	//默认值处理
	if p.Age < 0 {
		p.Age = 0
	}
	if p.Salary < 0 {
		p.Salary = 0
	}

	return p
}

func main() {
	// 对于不同实例化的需求只需要一个构造函数即可完成，只需要传入不同的 Options 函数即可
	p1 := NewPerson(
		WithName("Tom"),
		WithAge(23),
		WithAddress("66 London"),
		WithSalary(10000.0),
	)
	fmt.Println(p1)
}
