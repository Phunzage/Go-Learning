// map映射
// map(映射、字典)是一种内置的数据结构，它是一个无序的key-value对的集合
package main

import (
	"fmt"
)

func main() {
	// map使用之前一定要初始化
	// 1.短声明变量 + 字面量初始化
	m1 := map[string]int{
		"apple":  5,
		"banana": 4,
	}

	// 2.空map初始化
	m2 := map[int]string{}

	fmt.Println(m1)

	// 添加
	m2[1] = "zhangsan"
	m2[3] = "lisi"
	m2[5] = "del"

	// 删除
	delete(m2, 3)
	fmt.Println(m2)

	fmt.Println("----------复杂类型map----------")
	// slice类型
	scores := map[string][]int{
		"小明": {80, 90, 87},
		"李华": {90, 78., 99},
	}

	fmt.Println("slice类型\n", scores)

	fmt.Println()

	//结构体类型
	type User struct {
		Name  string
		Score float32
	}

	users := map[int]User{
		1: {"Tom", 78.6},
		2: {"John", 88.9},
	}

	fmt.Println("结构体类型\n", users)

}
