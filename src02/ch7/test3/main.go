// 并发goroutine03

// 一个测试题
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

}

/*
1.没有输出
因为main.goroutine一直在运行，只要创建完五个goroutine就结束，五个goroutine未运行到fmt.Println，main.goroutine就结束了

2.输出几个5
因为这是一个闭包，func包含外层的i，是引用的i的地址，i的地址最终值是5，一个goroutine运行完后，输出时i的地址里的值已变成5
*/
