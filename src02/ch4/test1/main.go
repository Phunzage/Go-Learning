//结构体

package main

import "fmt"

//结构体的声明
type Person struct {
	name string
	age  int
}

//可以连续声明
type CSS struct {
	class, color   string
	height, weight float64
}

func main() {

	//初始化字段值的几种方法
	//方法一
	//声明并初始化，返回值类型(完全等价于方法二)
	cat := CSS{ //注意此处&
		class:  "box",
		color:  "red",
		height: 23.4,
		weight: 13.5,
	}
	//声明并初始化，返回指针类型
	// &CSS{a, b, c} 是一种简写，底层仍然会调用 new()
	dog := &CSS{
		class:  "box",
		color:  "black",
		height: 23.4,
		weight: 13.5,
	}

	//方法二
	//强调声明一个变量，适合无需初始化的场景
	var p Person
	p.name = "zhangsan"
	p.age = 18

	//方法三(不推荐，因为不够灵活，无法即时赋值)，返回指针
	tom := new(Person)
	tom.name = "tom"
	tom.age = 18

	fmt.Println(cat)
	fmt.Println(dog)
	fmt.Printf("%#v\n", p)
	fmt.Println(tom)
}
