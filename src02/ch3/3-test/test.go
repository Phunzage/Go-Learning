package test3

import "fmt"

//执行第一步, A = 3
var A = test()

func test() int {
	fmt.Println("执行第一步, test()")
	return 3
}

func init() {
	A = 5
	fmt.Println("执行第二步, A =", A)
}
