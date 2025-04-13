package main

import (
	"fmt"
)

func main() {
	type Element struct {
		a int
		b int
	}

	fmt.Println(Element{3, 2})

	e := Element{1, 2}
	fmt.Printf("a = %v, b = %v\n", e.a, e.b)
	fmt.Println(e)
}
