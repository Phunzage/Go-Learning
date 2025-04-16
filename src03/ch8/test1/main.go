// 并发goroutine04

// 使用sync调度

package main

import (
	"fmt"
	"sync"
)

// 使用sync.WaitGroup
// 对于一个goroutine

var wait sync.WaitGroup

func Hello() {
	fmt.Println("Hello Goroutine")
	// 通知goroutine把计数器 - 1
	wait.Done()
}

func main() { // 开启一个主goroutine来执行main函数

	// 计数器 + 1
	wait.Add(1)
	// 开启一个goroutine来执行Hello函数
	go Hello()

	fmt.Println("Hello Main Goroutine")
	// 阻塞，等所有小弟都干完活之后才收兵
	wait.Wait()
}

// // 对于多个goroutine
// var wait sync.WaitGroup

// func Hello(i int) {
// 	fmt.Println("Hello Goroutine", i)
// 	// 通知goroutine把计数器 - 1
// 	wait.Done()
// }

// func main() { // 开启一个主goroutine来执行main函数

// 	// 计数器 + 1
// 	wait.Add(10000)
// 	// 开启一个goroutine来执行Hello函数
// 	for i := 0; i < 10000; i++ {
// 		go Hello(i)
// 	}

// 	// for循环可以写为：
// 	// for i := 0; i < 10000; i++ {
// 	// 	// fmt里的i用到外部的for循环，因此是一个闭包
// 	// 	go func() {
// 	// 		fmt.Println("hello", i)
// 	// 		wait.Done()
// 	// 	}()
// 	// }

// 	fmt.Println("Hello Main Goroutine")
// 	// 阻塞，等所有小弟都干完活之后才收兵
// 	wait.Wait()
// }
