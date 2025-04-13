package main

import (
	"fmt"
	"math/rand"
)

func randnum() {
	num := rand.Intn(10)
	fmt.Println("我最喜欢的数字是：", num)
}
