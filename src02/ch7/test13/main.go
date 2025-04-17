// sync.Map

// 并发读写sync.Map

// sync.Map方法：
/*
// 根据一个key读取值，返回值会返回对应的值和该值是否存在
func (m *Map) Load(key any) (value any, ok bool)

// 存储一个键值对
func (m *Map) Store(key, value any)

// 删除一个键值对
func (m *Map) Delete(key any)

// 如果该key已存在，就返回原有的值，否则将新的值存入并返回，当成功读取到值时，loaded为true，否则为false
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)

// 删除一个键值对，并返回其原有的值，loaded的值取决于key是否存在
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)

// 遍历Map，当f()返回false时，就会停止遍历
func (m *Map) Range(f func(key, value any) bool)
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var syncMap sync.Map

func main() {
	wait := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wait.Add(1)
		go func(n int) {
			// 将int类型转换成字符串
			key := strconv.Itoa(n)
			// 存储key和value数据
			syncMap.Store(key, n)
			// 查询key对应的value
			value, _ := syncMap.Load(key)
			fmt.Printf("key = %v, value = %v\n", key, value)
			wait.Done()
		}(i)
	}
	wait.Wait()
}

/*
输出（乱序）：
key = 0, value = 0
key = 1, value = 1
key = 19, value = 19
key = 4, value = 4
key = 5, value = 5
key = 2, value = 2
key = 6, value = 6
key = 18, value = 18
key = 11, value = 11
key = 14, value = 14
key = 12, value = 12
key = 15, value = 15
key = 3, value = 3
key = 16, value = 16
key = 7, value = 7
key = 17, value = 17
key = 8, value = 8
key = 9, value = 9
key = 10, value = 10
key = 13, value = 13
*/
