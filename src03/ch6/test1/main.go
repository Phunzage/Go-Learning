//封装
//封装就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作（方法），才能对字段进行操作
//区别于传统面向对象，Golang更强调简洁性

// 主要表现在结构体首字母的大小写
package main

import (
	"fmt"
	test1 "test/src03/ch6/test1/pack"
)

func main() {

	p1 := test1.NewPerson("Eli")
	fmt.Println(p1) //此时输出&{Eli 0 0}

	// 使用Set方法获取字段的值
	p1.SetPersonAge(20)
	p1.SetPersonSalary(5000)
	//使用Get方法输出字段的值
	fmt.Println(p1.GetPersonAge()) //20
	fmt.Println(p1.Name, p1.GetPersonAge(), p1.GetPersonSalary())
}
