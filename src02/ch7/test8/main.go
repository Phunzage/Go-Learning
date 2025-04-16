// 竞态问题：多个 goroutine 同时操作一个资源（临界区）的情况，这种情况下就会发生竞态问题（数据竞态）。

package main

import (
	"fmt"
	"sync"
)

// 案例

var (
	x  int64
	wg sync.WaitGroup
)

// 对全局变量x执行100000次 +1 操作

func add() {
	for i := 0; i < 100000; i++ {
		x += 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	// 开启两个goroutine同时操作x，发生数据竞争
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)

}

/* 输出：
181505
111946
125764
173487
...
开启了两个 goroutine 分别执行 add 函数，这两个 goroutine 在访问和修改全局的x变量时就会存在数据竞争，某个 goroutine 中对全局变量x的修改可能会覆盖掉另一个 goroutine 中的操作，所以导致最后的结果与预期不符。
*/

// 解决方法：
// 互斥锁
