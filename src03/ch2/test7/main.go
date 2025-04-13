// map遍历

package main

import (
	"fmt"
)

func main() {
	m := map[int]string{}

	s := []string{"zhangsan", "lisi", "wangwu"}

	for key, value := range s {
		m[key] = value
	}

	fmt.Println(m)
}
