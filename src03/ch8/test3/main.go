// channel和goroutine协同工作

package main

import (
	"fmt"
)

// 单向通道：只能完成发送或接收其中一个操作
// 在函数形参中加入

func f1(ch chan<- int) { // 此函数的ch只能接收
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// 从ch1中取出数据计算平方，把结果发送到ch2
func f2(ch1 <-chan int, ch2 chan<- int) { // 此函数的ch1只能发送，ch2只能接收
	for { // 从通道中取值的方式1
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
