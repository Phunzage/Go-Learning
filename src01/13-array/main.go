// 数组与切片

package main

import "fmt"

func main() {
	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = i * 2
	}

	b := a[2:5]
	fmt.Println(a)

	fmt.Println(b)
}
