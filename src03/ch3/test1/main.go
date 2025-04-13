//函数
// 打印九九乘法表

package main

import (
	"fmt"
)

func main() {
	nine()
}

func nine() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, j*i)
		}
		fmt.Println()
	}
}
