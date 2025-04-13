// 切片

package main

import (
	"fmt"
)

func main() {
	ageList := make([]int, 5)

	for i := range ageList {
		ageList[i] = i
	}

	// 将元素从数组切出来
	ageSlice := ageList[0:3]

	fmt.Printf("%v\n", ageList)
	fmt.Println(ageSlice)
}
