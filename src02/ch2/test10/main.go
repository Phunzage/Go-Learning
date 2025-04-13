// 闭包
// 闭包就是一个函数和与其相关的引用环境组合的一个整体
package main

import "fmt"

//累加器案例

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}
func main() {

	//当反复调用f函数时, 因为n是内部函数引用的(传的是地址), 因此每一次调用n都会进行更改

	f := AddUpper()   //因为AddUpper()是一个返回函数的函数, 所以f相当于执行了内部的匿名函数
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}

/*这一部分就是闭包
函数是一个匿名函数, 这个匿名函数引用了函数外的n, 因此这个匿名函数就和n形成一个整体, 构成闭包

var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}

	函数在自身函数体内找变量n, 未找到就返回上一层(此时为外层函数)寻找, 这样就满足函数 + 外部变量
*/
