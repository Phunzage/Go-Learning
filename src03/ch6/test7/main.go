// 接口注意事项和细节(下)
package main

import "fmt"

// ----6.一个自定义类型可以实现多个接口----

// 两个接口
type BInterfacer interface {
	Say()
}
type CInterfacer interface {
	Hello()
}

type EInterfacer interface {
	BInterfacer
	CInterfacer
	Test()
}

// 一个自定义类型
type Teacher struct{}

func (t Teacher) Say() {
	fmt.Println("Teacher Say()")
}

func (t Teacher) Hello() {
	fmt.Println("Teacher Hello()")
}

func (t Teacher) Test() {
	fmt.Println("Teacher Test()")
}

// ----6.End----

// 7.go接口中不能有任何变量

// ----8.一个接口（比如E接口）可以继承多个别的接口（比如B、C接口），这时如果要实现E接口，也必须将B、C接口的方法也全部实现----

// 9.接口interface类型默认是一个指针（引用类型），接收指针，如果没有对interface初始化就使用，那么会输出nil

// 10.空接口interface{} 没有任何方法，所以所有类型都实现了空接口，即我们可以把任何一个变量赋给空接口
// 定义一个空接口
// 方法一
type T1 interface{}

func main() {
	// ----6.main.Start----
	var ter Teacher
	var b BInterfacer = ter
	var c CInterfacer = ter
	b.Say()
	c.Hello()
	// ----6.main.End----
	// 输出
	// Teacher Say()
	// Teacher Hello()

	// ----8.main.Start----
	var ter2 Teacher
	var e EInterfacer = ter2
	e.Test()
	// 输出
	// Teacher Test()
	// 此时删除B或C任意一个接口，都会报错，因为实现E接口，也必须将B、C接口的方法也全部实现
	// ----8.main.End----

	// 10.main.Start
	// 定义空接口方法二
	var t2 T1 = ter

	// 也可以直接赋值
	var t3 interface{} = ter
	fmt.Println(t2, t3)
	// 输出{} {}
}
