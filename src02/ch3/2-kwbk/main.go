//init函数在跨包调用中, 先执行
// 1.被调用包的全局变量, 2.被调用包的init函数, 3.主包的全局变量定义, 4.主包的init函数, 5.主包的main函数

package main

import (
	"fmt"
	test3 "test/src02/ch3/3-test"
)

func init() {
	test3.A = 15
	fmt.Println("执行第三步, A =", test3.A)
}

func main() {
	test3.A = 20
	fmt.Println("执行第四步, A =", test3.A)
}
