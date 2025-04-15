// channel 通道

// Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过“通信共享内存”而不是通过“共享内存而实现通信”。

// 如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

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
	// 关闭后的通道有以下特点：

	// 对一个关闭的通道再发送值就会导致 panic。
	// 对一个关闭的通道进行接收会一直获取值直到通道为空。
	// 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	// 关闭一个已经关闭的通道会导致 panic。

}
