// channel练习

package main

import "fmt"

// 向管道发送数据
func Producer() chan int {
	// 定义一个具有2缓冲的管道
	// 为什么有5个奇数，两个缓冲完全够用，因为消费函数能够从管道及时拿取数据，甚至会发生刚向管道存入一个3消费者函数就拿走的情况
	ch := make(chan int, 2)
	go func() {
		// 取出奇数
		for i := range 10 {
			if i%2 == 1 {
				// 存入管道
				ch <- i
			}
		}
		// 任务完成后关闭管道
		close(ch)
	}()
	return ch
}

// 从管道接收数据
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
}
