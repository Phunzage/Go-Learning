// 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。

package main

import (
	"fmt"
	"sync"
)

// 案例

var (
	x    int64
	wait sync.WaitGroup
	m    sync.Mutex
)

// 对全局变量x执行100000次 +1 操作

func add() {
	for i := 0; i < 100000; i++ {

		// 操作数据前加锁
		m.Lock()
		x += 1
		// 操作完成后解锁
		m.Unlock()
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	// 开启两个goroutine同时操作x
	go add()
	go add()
	wait.Wait()

	fmt.Println(x)

}

// 输出
// 200000

// 同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。
