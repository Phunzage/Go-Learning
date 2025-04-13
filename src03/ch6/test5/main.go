// 接口入门
// 接口提供了一种方式来说明对象的行为：如果谁能搞定这件事，它就可以用在这儿。（关心“做什么”，而不是“怎样做”）

package main

import "fmt"

// interface类型可以定义一组方法，但是这些不需要实现，，并且interface不能包含任意变量。到某个自定义类型（比如结构体Phone）要使用的时候，再根据具体情况把这些方法写出来（实现）

// 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的 多态 和 高内聚低耦合 的思想

// Golang中的接口，不需要显式地实现。只要一个变量，含有接口类型的所有方法，那么这个变量就实现了这个接口。如下Usb和Usb2都可以被Phone和Camera实现
type Usber interface {
	Start()
	End()
}

type Usb2er interface {
	// 接口中不能创建变量
	Start()
	End()
}

// Usb3不可被Phone和Camera实现，因为Usb3有一个两者未实现的方法
type Usb3er interface {
	Start()
	End()
	Test()
}

type Phone struct {
}
type Camera struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) End() {
	fmt.Println("手机停止工作")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) End() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usber) { //这里Usb和Usb2接口都可以被实现
	// usb变量会根据传入的实参，来判断到底是Phone还是Camera，体现出多态的特点
	usb.Start()
	usb.End()
}

func main() {
	phone := Phone{}
	camera := Camera{}
	var computer Computer
	computer.Working(phone)
	computer.Working(camera)

}
