// 又一个错误示例

package main

import (
	"fmt"
	"time"
)

func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}

// 问题：21行的case <-time.After(time.Second)，在等待1s后退出，即select部分return，demo2()也会结束，result := <-ch也就没有了，但是同时，原设定3s后执行的无缓冲通道ch，在select退出的2s后接收到了数据，但是此时已经没有可以接收它数据的result了，无缓冲通道ch就会阻塞，程序什么都不会输出

func main() {
	demo2()
}

// 解决方法：
// 1.可以给select的超时时间设置大于3s，等待数据立马接收后再退出
// 2.可以设置一个有缓冲通道，即使程序无输出，但是不会发生ch阻塞
