// 文件操作2

// 文件读取
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// 方式一，os 包下的 ReadFile 函数
// 优点，自动打开和关闭文件，读取全部内容，适合小文件
// 缺点，大文件会一次性加载到内存，可能导致内存溢出

func OsReadFile() {
	bytes, err := os.ReadFile("test3.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}

// 方式二，io 包下的 ReadAll 函数
// 优点，一行代码即可读取所有内容
// 适用于任何 io.Reader 类型（文件、网络流、内存等）
// 自动处理读取结束：一直读取到 io.EOF，无需手动循环
// 缺点，内存占用高：将全部内容加载到内存，大文件可能导致内存溢出
// 无法控制读取过程：无法逐行或分块处理数据
// 不适合流式处理：必须等待所有数据读取完毕才能处理

func IoReadAll() {
	file, err := os.OpenFile("test3.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件访问异常")
	} else {
		bytes, err := io.ReadAll(file)
		fmt.Println("文件打开成功", file.Name())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes))
		}
		file.Close()
	}
}

// 方式三，使用 bufio.scanner 以固定缓冲读取文件
// 优点，内存高效：默认按行读取（缓冲大小 4KB），适合大文件
// 支持自定义分割函数（如按单词、JSON 分隔符）
// 自动处理 UTF-8 编码
// 缺点，需要手动打开和关闭文件
// 单行长度超过缓冲区时会报错（可通过 Buffer() 调整）

func Scanner() {
	file, err := os.Open("test3.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// 循环的读取文件的内容，逐行处理
	for scanner.Scan() {
		line := scanner.Text() // 逐行处理
		fmt.Println(line)

	}
	// 检查读取过程中是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取错误", err)
	}
	fmt.Println("文件读取结束")
}

func main() {

	OsReadFile()
	time.Sleep(time.Second)
	fmt.Println()
	IoReadAll()
	time.Sleep(time.Second)
	fmt.Println()
	Scanner()
}
