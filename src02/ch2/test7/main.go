//将函数作为参数传递

package main

import (
	"fmt"
)

func main() {
	F(5, add)

}

func add(a, b int) {
	fmt.Println(a, b, a+b)
}

func F(i int, f func(int, int)) {
	f(i, i+1)
}
