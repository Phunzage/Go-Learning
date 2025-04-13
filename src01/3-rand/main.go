package main

import (
	"fmt"
	//所在位置
	"math/rand"
)

func main() {
	randnum()
	// 随机返回0~9, 加1返回1到10
	for i := 0; i < 3; i++ {
		var num1 = rand.Intn(7) + 1
		var num2 = rand.Intn(13) + 1
		fmt.Printf("你选择第%v行，第%v列的英雄\n", num1, num2)
	}
}
