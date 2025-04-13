//指针副本与值传递的思考

//go的传递本质上都是值传递，即使是指针传递，也是先创建一个指针的副本，副本指针与原指针指向同一个地址，通过副本指针修改目标地址的值之后，副本指针生命周期结束后被销毁，对应的值也得到了改变

package main

import (
	"fmt"
)

func checkPointer(p *string) {
	fmt.Printf("指针地址（副本）: %p\n", &p) // 与原指针地址不同
}

func main() {
	name := "xiaoming"
	p := &name
	fmt.Printf("原指针地址: %p\n", &p)
	checkPointer(p) // 传递指针的副本
}
