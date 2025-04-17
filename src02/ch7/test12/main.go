// sync.Map方法

// Go 语言中内置的 map 不是并发安全的
// 而sync.Map 是官方提供的一种并发安全 Map 的实现

package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发不安全
// 初始化一个map，键是string， 值是int
var m = make(map[string]int)

// 获取对应键的值
func get(key string) int {
	return m[key]
}

// 更改键和值
func set(key string, value int) {
	m[key] = value
}

func main() {
	wait := sync.WaitGroup{}
	// 开启10个goroutine
	for i := 0; i < 10; i++ {
		// 计数器 + 1
		wait.Add(1)
		// 数据竞态
		go func(n int) {
			// 把整型转换成字符串
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("key = %v, value = %v\n", key, get(key))
			// 计数器 - 1
			wait.Done()
		}(i)
	}
	wait.Wait()
}

/*
输出：
很大概率报错
*/

// 解决方法：
// 互斥锁（高频写操作）
// sync.Map（读多写少，且键长期稳定）
