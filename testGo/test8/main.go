// 结构体练习2
// 定义一个box

package main

import "fmt"

type box struct {
	long   float64
	width  float64
	height float64
}

func (b *box) getInput() float64 {
	return b.long * b.width * b.height
}

func main() {
	b := new(box)
	fmt.Print("long: ")
	fmt.Scan(&b.long)
	fmt.Print("width: ")
	fmt.Scan(&b.width)
	fmt.Print("height: ")
	fmt.Scan(&b.height)
	res := b.getInput()
	fmt.Println(res)
}
