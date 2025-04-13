//方法的值接收与指针接收，方法定义注意要点

package main

import "fmt"

// 为int的别名定义方法，不能为int定义
type MyInt int

/* 值接收组合，不会修改值，传递的是副本
// func (m MyInt) Set(num int) {
// 	m = MyInt(num)
}

输出：1
*/

//指针接收组合
func (m *MyInt) Set(num int) {
	*m = MyInt(num)
}

func main() {
	myInt := MyInt(1)
	myInt.Set(2)
	fmt.Println(myInt) //输出2
}

//注意，不能为int类型定义方法（即不能在非本地的包里定义方法）
//如，import "...\list"，不能为list.List(别的包的元素)定义方法，但是可以为其起别名定义方法
