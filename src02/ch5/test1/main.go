//方法
//方法与函数的区别在于，方法拥有接收者，而函数没有，且只有自定义类型能够拥有方法。

// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
// 函数体
// }

package main

import "fmt"

// MyInt是[]int类型的别名
type MyIntArray []int

// 声明一个得到int数组某个元素的方法
//go推荐使用接收者类型名称首字母的小写，如MyInt m
func (m MyIntArray) Get(index int) int {
	return m[index]
}
func main() {
	myint := MyIntArray{2, 4, 6, 8, 10}
	fmt.Printf("%#v", myint.Get(3))
}

// 方法的使用就类似于调用一个类的成员方法，先声明，再初始化，再调用。
