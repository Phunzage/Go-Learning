// 匿名函数

package main

import (
	"fmt"
)

func main() {
	fmt.Println("---匿名函数---")

	var age = func(a, b int) int {
		return a * b
	}

	fmt.Println(age(12, 5))
}
