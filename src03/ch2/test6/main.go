// map查找

package main

import (
	"fmt"
)

func main() {
	age := map[string]int{
		"浩然": 20,
		"梅梅": 23,
	}

	// 查找是否有"浩然"
	value, ok := age["浩然"]

	fmt.Println(value, ok) //20 true
}
