// 文件操作4

// 复制文件
package main

import (
	"fmt"
	"os"
)

// 方式一，将原文件中的数据读取出来，然后写入目标文件中

func RWFiles() {
	// 从源文件中读取数据
	data, err := os.ReadFile("from.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入目标文件
	err = os.WriteFile("to.txt", data, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("复制成功")
	}
}

func main() {
	RWFiles()
}
