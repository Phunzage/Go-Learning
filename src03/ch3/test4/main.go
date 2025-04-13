// 根据用户输入的不同，返回不同的函数

package main

import (
	"fmt"
)

func main() {

	fmt.Println(`
		1. 登录
		2. 注册
		3. 用户操作
	`)
	var num int
	fmt.Print("请输入操作代码：")
	fmt.Scan(&num)

	m := map[int]func(){
		1: signIn,
		2: signUp,
		3: user,
	}

	// 直接调用函数：signIn()，需要 () 触发执行。
	// 通过映射调用函数：m[num]()，同样需要 () 触发执行。

	// () 是函数调用的语法标记。无论函数是通过变量、映射还是其他方式引用，必须通过 () 显式调用才能执行其逻辑。
	m[num]()

}

func signIn() {
	fmt.Println("登录")
}

func signUp() {
	fmt.Println("注册")
}

func user() {
	fmt.Println("用户操作")
}
