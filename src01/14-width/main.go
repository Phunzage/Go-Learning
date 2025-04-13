package main

import (
	"fmt"
)

func main() {
	var i, j float32
	for i = 16; i < 1000; i++ {
		for j = 9; j < i; j++ {
			if i/j == 1.7777777777777 {
				fmt.Println(i, j)
			} else {
				continue
			}
		}
	}
}
