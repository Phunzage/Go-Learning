package main

import "fmt"

type MyInt []int

func (m MyInt) Get(index int) int {
	return m[index]
}
func main() {
	myint := MyInt{2, 4, 6, 8, 10}
	fmt.Printf("%#v", myint.Get(3))
}
