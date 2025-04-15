// 并发gotoutine01
// Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。

package main

import "fmt"

// 这种情况的缺点，不稳定
// main函数本身是一个main.goroutine，两个goroutine并发执行，若main.goroutine执行速度快一些，hello函数的goroutine未执行完，就不会被打印
func Hello() {
	fmt.Println("Hello Goroutine")
}
func main() {
	go Hello()
	fmt.Println("Hello Main Goruntine")
}

/*
输出
Hello Goroutine
Hello Main Goruntine
或
Hello Main Goruntine
或
Hello Main Goruntine
Hello Goroutine
*/
