// 一个错误演示

/*

错误示例

func demo1() {
	wg := sync.WaitGroup{}

	// 向管道发送数据
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 结束后关闭管道
	close(ch)

	wg.Add(3)
	// 启动3个goroutine
	for j := 0; j < 3; j++ {
		go func() {
			for {
				// 从管道取值并且赋给task
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				// time.Sleep(1 * time.Second)	// 慢速输出观察错误

				// 打印task	因为是3个goroutine并发，所以0-9随机打印
				fmt.Println(task)

				// 问题：取完0-9后管道关闭，但goroutine还在取值，一直输出0且进行不到下一步wg.Done()，无限输出0

				// 解决办法
				// 使用for range接收值：for task := range ch
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

*/

// 正确示例

package main

import (
	"fmt"
	"sync"
)

func demo1() {
	wg := sync.WaitGroup{}

	// 向管道发送数据
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 结束后关闭管道
	close(ch)

	wg.Add(3)
	// 启动3个goroutine
	for j := 0; j < 3; j++ {
		go func() {
			// 从管道取值并且赋给task
			for task := range ch {
				// for task := range ch是一个语法糖，等价于
				// task, ok := <-ch
				// if !ok {
				//	break
				// }
				// 当通道被关闭后，for range循环会自动退出

				// 这里假设对接收的数据执行某些操作

				// 打印task	因为是3个goroutine并发，所以0-9随机打印
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	demo1()
}
