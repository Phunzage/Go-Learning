package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var command = "hello world"

	//strings.Contains(), 查看字符串中是否包含特定的词，bool类型
	var exit = strings.Contains(command, "hello")
	fmt.Println("the result is", exit) //此处exit自动为bool类型

	var a = 7
	for i := 5; i < a; i-- {
		if i < 0 {
			break
		}
		fmt.Println(i)
		//隔2秒输出
		time.Sleep(time.Second)
	}
	fmt.Print("finish")

	//一个猜数游戏
	//首先定义一个1到100的整数。然后让计算机生成一个1到100的随机数，并显示
	// 计算机猜测的结果是太大了还是太小了，没猜对的话就继续猜，直至猜对。

	var (
		numin  = 20
		numout int
	)
	for {
		numout = rand.Intn(100) + 1
		if numout < numin {
			fmt.Printf("你猜的%v,太小了\n", numout)
		} else if numout > numin {
			fmt.Printf("你猜的%v,太大了\n", numout)
		} else {
			fmt.Println("猜对了")
			break
		}
		time.Sleep(time.Second)
	}

}
