// 接口
// 接口是一种类型，是一种抽象的类型
// 接口是一种类型，是一种抽象的类型
// 接口是一种类型，是一种抽象的类型

package main

import "fmt"

// 定义一个场景，有许多小动物，比如猫和狗，它们饿了会叫

type Cat struct {
}

type Dog struct {
}

// 猫会叫
func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

// 狗会叫
func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

// 定义一个饿肚子就会叫的场景
// func CatHungry(c Cat) {
// 	c.Say()
// }

// func DogHungry(d Dog) {
// 	d.Say()
// }

// 现在来了一只小羊，小羊饿了会叫，后面又来了一只……
type Sheep struct {
}

func (s Sheep) Say() {
	fmt.Println("咩咩咩")
}

// 定义一个饿肚子就会叫的场景
// func SheepHungry(s Sheep) {
// 	s.Say()
// }

// 当有越来越多的小动物来，我们的代码该如何扩展??
// 这时就用到了接口

// 我们可以约定一个Sayer类型，它必须实现一个Say()方法，只要饿肚子了，我们就调用Say()方法。

type Sayer interface {
	Say()
}

// 通用的饿了会叫的方法
func MakeHungry(s Sayer) {
	s.Say()
}

// 此时上面的 “定义一个饿肚子就会叫的场景” 都可以删掉

func main() {
	// 无接口情况下
	// 猫和狗饿了在叫
	// var c1 Cat
	// CatHungry(c1)	// 喵喵喵
	// var d1 Dog
	// DogHungry(d1)	// 汪汪汪

	// 有接口情况下
	var c2 Cat
	var d2 Dog
	var s2 Sheep
	// 我们通过使用接口类型，把所有会叫的动物当成Sayer类型来处理，只要实现了Say()方法都能当成Sayer类型的变量来处理。
	MakeHungry(c2) //喵喵喵	此时c2就成了Sayer
	MakeHungry(d2) //汪汪汪
	MakeHungry(s2) //咩咩咩

	// 一个接口类型的变量能够存储所有实现了该接口的类型变量。
	var s Sayer
	var c3 Cat
	s = c3
	s.Say() // 喵喵喵
}
