// 文件操作2

// 文件读取
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("test2.md")
	if err != nil {
		fmt.Println("open file err: ", err)
	}

	defer file.Close()

	// 创建一个Reader，是带缓冲的，默认的缓冲区为4096（缓冲的好处：不是一次性读完文件，而是读一部分，处理一部分，方便处理一些比较大的文件）
	reader := bufio.NewReader(file)
	// 循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行符结束
		// 将内容放到str中，将错误放到err中
		if err == io.EOF {
			break // 读到文件的末尾
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}
