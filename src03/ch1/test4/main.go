// 变量类型最大值

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.0f\n", math.Pow(2, 63))
	// int64最大值：2^63 - 1
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	// var n2 int = 9223372036854775808	报错
	// fmt.Println(n2)
}
