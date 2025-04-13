// 函数定义

package main

import (
	"fmt"
)

func add(n1, n2 int) int {
	return n1 + n2
}

// 多参数传入
func moreNum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(add(123, 127))

	fmt.Println("----------多参数传入----------")
	fmt.Println(moreNum(1, 2, 3))

	// 给函数传入切片
	aggSlice := []int{20, 21, 22, 23}
	// 在后面加...
	fmt.Println("----------传入切片----------")
	fmt.Println(moreNum(aggSlice...))

}
