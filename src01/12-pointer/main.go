// 指针与结构体

package main

import (
	"fmt"
)

func main() {
	a := 13
	type Vertex struct {
		b int
		c int
	}

	v := Vertex{34, 56}
	p := &v
	p.b = 25
	p.c = 78

	fmt.Println(a, p.b, p.c)

	ver()
}
