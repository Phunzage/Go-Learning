// 切片排序

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 定义数组
	ages := []int{3, 5, 12, 1, 45, 55}

	// 升序
	sort.Ints(ages)
	fmt.Println(ages)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	fmt.Println(ages)
}
