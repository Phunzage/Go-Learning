// defer与闭包 测试题

package main

import "fmt"

// 1.defer

func main() {
	defer fmt.Println(Fn1())
	fmt.Println("3")
}

func Fn1() int {
	fmt.Println("2")
	return 1
}

/*
输出：
2
3
1
*/

// 2.defer
/*
func main() {
	var a, b int
	a = 1
	b = 2
	defer fmt.Println(sum(a, b))
	a = 3
	b = 4
  }

  func sum(a, b int) int {
	return a + b
  }

-------

输出：
3
*/

// 3.闭包
/*
func main() {
	var a, b int
	a = 1
	b = 2
	f := func() {
	  fmt.Println(sum(a, b))
	}
	a = 3
	b = 4
	f()
  }

--------

  输出：
  7
*/

// 4.defer + 闭包
/*
func main() {
  var a, b int
  a = 1
  b = 2
  defer func() {
    fmt.Println(sum(a, b))
  }()
  a = 3
  b = 4
}

--------

输出：
7
*/

//5.在4的基础上把闭包去除
/*
func main() {
  var a, b int
  a = 1
  b = 2
  defer func(num int) {
    fmt.Println(num)
  }(sum(a, b))
  a = 3
  b = 4
}

--------
输出：
3
*/

// 综上：
// defer fmt.Println(sum(a,b))
// 等价于
// defer fmt.Println(3)
