// 接口与继承

// 当A结构体继承了B结构体，那么A结构体就自动继承了B结构体的字段和方法，并且可以直接使用
// 当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口，因此可以认为，实现接口是对继承机制的补充

package main

import "fmt"

type Monkey struct {
	Name string
}

// 小猴子继承大猴子的字段和方法
type LittleMonkey struct {
	Monkey
}

// 鸟类接口，可以实现飞翔
type Birder interface {
	Flying()
}

// 鱼类接口，可以实现游泳
type Fisher interface {
	Swimming()
}

func (m Monkey) Climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, "学会了飞翔")
}

func (lm LittleMonkey) Swimming() {
	fmt.Println(lm.Name, "学会了游泳")
}

func main() {
	// 创建小猴子实例
	m := LittleMonkey{
		Monkey{
			Name: "Wukong",
		},
	}

	m.Climbing()
	m.Flying()
	m.Swimming()

}
