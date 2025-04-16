// select多路复用

package main

import "fmt"

func main() {
	// 创建一个1缓存通道
	ch := make(chan int, 1)
	for i := range 10 {
		select {
		// i = 0时， ch里没有值，此路不通
		// i = 1时，ch里的值为0，传给x并打印，break，ch空，进入i = 2，以此类推
		case x := <-ch:
			fmt.Println(x)
			// i = 0时，0存入ch
			// i = 2时，2存入ch
		case ch <- i:
		default:
			fmt.Println("无事发生")
		}
	}

	// 若ch := make(chan int, 10)
	// 输出结果每次都会随机
	// 因为select如果有多个 case 同时满足，select 会随机选择一个执行。

	// Select 语句具有以下特点。

	// 可处理一个或多个 channel 的发送/接收操作。
	// 如果多个 case 同时满足，select 会随机选择一个执行。
	// 对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

}
