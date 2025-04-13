// 编写两种函数计算一个float64类型的平方根

package main

import (
	"fmt"
	"math"
)

func main() {

	var num float64 = 46.5

	fmt.Printf("%.2f的平方根是：%.2f\n", num, MySqrt1(num))

	fmt.Printf("%.2f的平方根是：%.2f", num, MySqrt2(num))
}

// 非命名返回值
func MySqrt1(x float64) float64 {
	return math.Sqrt(x)
}

// 命名返回值
func MySqrt2(x float64) (result float64) {
	result = math.Sqrt(x)
	return
}
