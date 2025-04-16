// sync.Once方法

// 当在使用一些数据结构时，如果这些数据结构太过庞大，可以考虑采用懒加载的方式，即真正要用到它的时候才会初始化该数据结构。
package main

import (
	"fmt"
	"sync"
)

/*

type MySlice []int

func (m *MySlice) Get(i int) (int, bool) {
	if *m == nil {
		return 0, false
	} else {
		return (*m)[i], true
	}
}

func (m *MySlice) Add(i int) {
	// 当真正用到切片的时候，才会考虑去初始化
	// 问题：当并发时，初始化操作可能会执行多次
	if *m == nil {
		*m = make([]int, 0, 10)
	}
	*m = append(*m, i)
}

*/

// 如果只有一个协程使用肯定是没有任何问题的，但是如果有多个协程访问的话就可能会出现问题了。比如协程 A 和 B 同时调用了 Add 方法，A 执行的稍微快一些，已经初始化完毕了，并且将数据成功添加，随后协程 B 又初始化了一遍，这样一来将协程 A 添加的数据直接覆盖掉了，这就是问题所在。

// 解决方法：
// 使用sync.Once

var (
	wait sync.WaitGroup
	// lock sync.Mutex
)

type MySlice struct {
	s    []int
	once sync.Once
}

func main() {
	var slice MySlice
	wait.Add(100)
	// 开启4个并发
	for i := 0; i < 100; i++ {
		go func() {
			slice.Add(1)
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(slice.Len())
}

func (m *MySlice) Add(i int) {
	// 当真正用到切片的时候，才会考虑去初始化
	// 此操作只被执行一次，之后的并发不会触发
	m.once.Do(func() {
		fmt.Println("初始化")
		if m.s == nil {
			m.s = make([]int, 0, 10)
		}
	})

	// 初始化之后，可能存在竞态问题，可以加锁解决
	// lock.Lock()
	m.s = append(m.s, i)
	// lock.Unlock()
}

func (m *MySlice) Len() int {
	return len(m.s)
}
