// 常量

package main

import (
	"fmt"
)

// 常量定义，定义时就要赋值
// 赋值之后不能再修改
const pi = 3.14

func main() {
	// pi = 4.7 错误
	fmt.Print(pi)
}
