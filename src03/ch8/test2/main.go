// channel 通道

// 命名格式
// var 变量名 chan 存储的变量类型

// slice，map，channel都是引用类型，要使用make初始化
package main

import "fmt"

func main() {

	// 无缓冲区的通道，称为同步通道
	// var ch1 chan int = make(chan int)	// 尾部不添加缓冲值

	// 有缓冲区的通道，异步通道
	var ch1 chan int = make(chan int, 1) // 可选缓冲区大小，即通道里最多存放几个值
	// 等价于
	// ch1 := make(chan int, 1)

	// 给通道变量赋值
	ch1 <- 10

	// 从通道变量取值
	x := <-ch1

	// len(ch1)	// 取通道中元素的数量
	// cap(ch1)	// 拿到通道的数量

	fmt.Println(x)

	// 关闭通道
	close(ch1)

}
