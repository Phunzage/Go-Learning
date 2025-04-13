package main

import "fmt"

type Vertex struct {
	a int
	b int
}

func ver() {

	v := Vertex{}

	var (
		a = Vertex{1, 2}
		b = Vertex{a: 4}
		c = Vertex{5, 6}
	)

	fmt.Println(a, b, c, v.b)
}
