// 接口注意事项和细节(上)

package main

import "fmt"

// ----1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量----
type AInterfacer interface {
	Say()
}

// 自定义（结构体）类型
type Stu struct {
	Name string
}

// 自定义类型实现了接口
func (stu Stu) Say() {
	fmt.Println("stu Say()~")
}

// ----1.End----
func main() {
	// ----1.main.Start----
	// 此处的stu就是一个实现了接口的自定义类型变量
	var stu Stu
	var a AInterfacer = stu
	a.Say()
	// ----1.End----

	// 4.main.Start----
	var i integer
	var b AInterfacer = i
	b.Say()
	// ----4.main.End----
}

// 2.接口中的所有方法都没有方法体,即都是没有实现的方法----
// 3.在go中,一个自定义类型需要将某个接口的所有方法都实现,我们说这个自定义类型实现了该接口

// ----4.只要是自定义数据类型,都可以实现接口,不仅仅是结构体类型----
type integer int

func (i integer) Say() {
	fmt.Println("integer say i = ", i)
}

// ----4.End----
