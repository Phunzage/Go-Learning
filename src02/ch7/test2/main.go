// 并发goroutine02

// 针对goroutine01存在的问题，可以加入time.Sleep()来缓解
package main

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello Goroutine")
}
func main() {
	go Hello()
	fmt.Println("Hello Main Goruntine")
	// 等待一秒后再结束main.goroutine
	time.Sleep(time.Second)
}

/*
输出
Hello Main Goruntine
Hello Goroutine
*/

// 为什么总是先输出Hello Main Goruntine
// 因为在程序中创建 goroutine 执行函数需要一定的开销，而与此同时 main 函数所在的 goroutine 是继续执行的。
