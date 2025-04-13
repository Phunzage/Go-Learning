// String() 方法特殊机制

// fmt包中的Stringer接口定义为:
// type Stringer interface {
//     String() string
// }

package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

// 任何实现了String()方法的类型，都会隐式实现该接口。当使用fmt包的打印函数（如Printf、Println）时，若值的类型实现了Stringer接口，会自动调用其String()方法生成字符串表示。

// String()方法定义在指针接收者*TwoInts上，因此只有指针类型的TwoInts实例（如two1）会触发该方法。值类型的实例（如TwoInts{}）不会触发。
func main() {
	two1 := new(TwoInts) // 此处若使用TwoInts{}则不会调用
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is %v\n", two1)  // two1 is (12/10)
	fmt.Println("two1 is ", two1)     // two1 is  (12/10)
	fmt.Printf("two1 is %T\n", two1)  // two1 is *main.TwoInts
	fmt.Printf("two1 is %#v\n", two1) // two1 is &main.TwoInts{a:12, b:10}

	// %v：默认格式，会调用String()方法（若存在）。
	// %#v：输出Go语法表示，直接显示结构体字段，不调用String()方法。
	// %T：输出值的类型。
	// Println：自动选择格式，对实现了Stringer接口的值调用String()。
	// 总结: %v和Println会触发，%#v和%T不会。

}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
