// 接口的值接收者和指针接收者

package main

import "fmt"

type Payer interface {
	Pay()
}

type AliPay struct {
}

// 下面两种方法关掉其一演示

// 使用值接收者定义方法
func (a AliPay) Pay() {
	fmt.Println("值类型，客户使用支付宝付款")
}

// 使用指针接收者定义方法
// func (a *AliPay) Pay() {
// 	fmt.Println("指针类型，客户使用支付宝付款")
// }

func UsePay(p Payer) {
	p.Pay()
}

func main() {
	// 使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。
	var a1 = AliPay{}  //传值类型，无论调用值接收者方法还是指针接收者方法都可以
	var a2 = &AliPay{} // 此处传指针可以

	// 在指针类型中报错，可改为&a1
	UsePay(a1) // 值类型中，输出: 值类型，客户使用支付宝付款

	// 两种类型都正确
	UsePay(a2) // 值类型，客户使用支付宝付款

	// 另一种赋值方法
	var p Payer
	var a6 = AliPay{}
	p = a6
	UsePay(p) //值类型，客户使用支付宝付款
}

// 综上
// 使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。
// 使用指针接收者实现接口后，只有结构体指针类型的变量可以赋值给该接口变量。
