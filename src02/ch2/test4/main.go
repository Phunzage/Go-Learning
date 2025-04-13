// defer用法01

package main

import "fmt"

func main() {
	fmt.Println("这个语句在上面")
	defer f1()
	fmt.Println("这个语句在下面")

}

func f1() {
	fmt.Println("延迟至调用函数结束")
}
