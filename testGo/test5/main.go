// 通过调用方法打印矩形及矩形面积
// 结构体可以没有字段

package main

import "fmt"

type methodUtils struct{}

//打印特定矩形
func (mu methodUtils) print(m, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//打印矩形面积
func (mu methodUtils) printArea(len, width int) int {
	return len * width
}
func main() {
	var me methodUtils

	//打印特性矩形方法调用
	me.print(10, 8)

	//打印矩形面积方法调用
	res := me.printArea(10, 8)
	fmt.Println(res)
}
