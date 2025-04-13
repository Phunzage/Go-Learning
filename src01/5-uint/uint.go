package main

import (
	"fmt"
	"time"
)

func main() {

	var red, green, blue uint8 = 0, 141, 213
	//用十六进制表示数值
	// var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", red, green, blue) //用%x格式化输出

	//指定最小宽度和填充,0是指用0填充，2是占两位
	fmt.Printf("color: #%02x%02x%02x\n", red, green, blue)

	//整数环绕
	var num1 uint8 = 255 //uint类型0~255，再加会回到0
	num1++
	fmt.Println(num1)
	var num2 int8 = 127 //int类型是-128~127，再加会回到-128
	num2++
	fmt.Println(num2)

	//返回当前时间(%T返回数据类型)
	nowtime := time.Now()
	fmt.Printf("%T\n现在时间是：%v", nowtime, nowtime)
}
