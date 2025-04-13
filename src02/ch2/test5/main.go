// defer用法02

package main

import "fmt"

func main() {
	f()

	//输出0
}

func f() (i int) {
	i = 0
	// 此时fmt.Println(i)的值已经确定为0（此时已压入栈），只是在return返回后输出，后续的i++不会影响
	defer fmt.Println(i)
	i++
	return
}
