package main

import "fmt"

func main() {
	fmt.Println(`hello \n world`)
	//使用``打印多行文本
	fmt.Println(`
	makabaka
	yigubigu
	`)

	//给类型起别名
	type hello = int
	var a hello = 8
	fmt.Printf("a是%T类型，a的值 = %v\n", a, a)

	//用%c打印所对应的字符
	//任何整数类型都可以使用%c来打印，但是rune意味着该数值表示了一个字符
	var (
		pi    rune = 960
		alpha rune = 940
		omega rune = 969
		bang  byte = 33
	)
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c", pi, alpha, omega, bang)

}
