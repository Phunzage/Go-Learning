// 常量const与变量var相关

package main

import "fmt"

//ear的作用域是package级别，且此处不能用短声明
var ear = 90

func main() {
	const pi = 3.14159265

	var weight = 23.4
	// 自加,go只有后加
	weight++
	fmt.Printf("this weight %v %v\n", "is", weight)
	// 正数向左填充空格,负号向右
	fmt.Printf("%-4vhello%4v", 12, pi)

	// 同时声明多个变量
	// 第一种形式

	// var (
	// 	one = 1
	// 	two = 2
	// )

	//第二种形式
	// var one, two = 1, 2

	//第三种形式
	one, two := 1, 2
	fmt.Printf("one = %v, two = %v", one, two)

	//for, if和switch里不能用var声明
	for i := 0; i < 5; i++ {
		fmt.Println("oioi", i)
	}

}
