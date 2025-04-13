// 斐波拉契函数
package main

import (
	"fmt"
)

func main() {
	x, y := 5, 7
	x, y = y, x
	fmt.Println(x, y)

	fmt.Println(fib(7))
}

// 斐波拉契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}

	return x
}
