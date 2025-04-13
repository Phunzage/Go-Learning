// 和函数不同，方法的值接收和指针接收可以相互转换（底层编译器实现）
package main

import "fmt"

type twoInts struct {
	a int8
	b int8
}

func main() {
	//值类型调用指针接收
	t1 := twoInts{3, 4}
	fmt.Println(t1.AddThem()) //3 + 4 = 7

	//指针类型调用指针接收
	t2 := new(twoInts)
	t2.a = 10
	t2.b = 20
	fmt.Println(t2.AddToNum(8)) //10 + 20 + 8 = 38

	//值类型调用值接收
	t3 := twoInts{5, 6}
	fmt.Println(t3.TestAdd()) //5 + 6 = 11

	//指针类型调用值接收
	t4 := &twoInts{7, 8}
	fmt.Println(t4.TestAdd()) //7 + 8 = 15
}

//无论外面传值还是指针，都是指针拷贝
func (t *twoInts) AddThem() int8 {
	return t.a + t.b //等同于(&t).a + (&t).b, go的语法糖
}

//无论外面传值还是指针，都是指针拷贝
func (t *twoInts) AddToNum(num int8) int8 {
	return t.a + t.b + num
}

//无论外面传值还是指针，都是值拷贝
func (t twoInts) TestAdd() int8 {
	return t.a + t.b
}
