// 继承与匿名结构体
package main

import "fmt"

type A struct {
	Name string
	age  int
}

// 结构体可以使用嵌套匿名结构体的所有字段和方法,不管首字母是大写还是小写
type B struct {
	Name string
	A
}

// 传地址的匿名结构体嵌套
type C struct {
	*A
}

func (a *A) SayHello() {
	fmt.Printf("A.name %s SaySello\n", a.Name)
}

func (a *A) SayOK() {
	fmt.Printf("A.name %s, age = %d SayOK\n", a.Name, a.age)
}

func (b *B) SayGood() {
	fmt.Printf("B.name %s SayGood\n", b.Name)
}

func main() {
	var b B
	// 当B与A有相同字段时,遵循就近原则,即先在B中找,没有向A中寻找
	// 如果B中没有Name,则B.Name会给A.Name赋值
	b.Name = "Tom"
	b.SayGood()
	// 精确赋值
	b.A.Name = "Eli"
	b.age = 20
	// 因为b.A.Name = "Eli",所以使用Eli调用SayHello和SayOK方法
	b.SayHello()
	b.SayOK()

	c := C{
		&A{
			Name: "David",
			age:  30,
		},
	}

	fmt.Println(c)      // 此时输出一个地址
	fmt.Println(*(c.A)) // 取地址输出, 可写为 *c.A
}
