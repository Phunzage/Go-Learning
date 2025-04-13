// 工厂模式
// 导入外部包时,结构体首字母要大写,但是当结构体首字母是小写仍然想要导出时,就要使用工厂模式

// 核心思路:
// 结构体导出不能使用,可以创建一个可以导出的方法,此方法顺带将一些任务(结构体初始化)完成
package main

import (
	"fmt"
	test "test/src03/ch5/test4/factory"
)

func main() {
	stu := test.NewStudent("tom", "3班", 20)
	fmt.Println(stu) //因为返回的指针,所以前有&

	stu.Age = 30
	fmt.Println(stu.Age) //可访问
	// fmt.Println(stu.class) 报错,不可访问
	fmt.Println(stu.GetClass()) //输出 3班

	//对于结构体首字母大写而里面字段首字母小写的字段,字段同样不可访问
	cat1 := test.Cat{
		// age: 3,	不可访问
		Color: "black",
	}
	fmt.Println(cat1.Color)
	// fmt.Println(cat1.age)
}
