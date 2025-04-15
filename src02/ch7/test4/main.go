// channel无缓冲通道

package main

import "fmt"

func main() {

	// 无缓冲管道

	ch := make(chan int)
	// ch := make(chan int)是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
	// 可以想象一个无缓冲通道是一个货拉拉司机（没有车），他手上有你的一份货物，他必须交到你（接收方）的手上，程序才可以进行下一步，否则他将一直等待，程序也将发出错误

	go recv(ch) // 解决办法1，放在这里可以，因为接收方已经运行并准备好接收
	ch <- 10
	// go recv(ch)	注意放在这里不可以，因为ch <- 10已经发生阻塞，永远无法到达这一步

	fmt.Println("发送成功")
	//输出
	// fatal error: all goroutines are asleep - deadlock!
	// 没有接收方，所以发生错误

}

// 解决办法
// 1.创建一个goroutine去接收值

func recv(ch chan int) {
	res := <-ch
	fmt.Println("接收成功", res)
}

// 2.创建一个有缓冲通道，见test5
