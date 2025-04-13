// 平方根工具

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	p := 0.0
	for i := 0; i < 10; i++ {
		p = z
		z -= (z*z - x) / (2 * z) //牛顿法
		fmt.Println(z)

		if math.Abs(z-p) < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
