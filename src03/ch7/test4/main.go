// 文件操作3

// 写文件操作

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 方式一，os.WriteFile
func OsWriteFile() {
	err := os.WriteFile("test4.txt", []byte("hello world\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// 方式二，io.WriteString

func IoWriteString() {

	file, err := os.OpenFile("test4.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	// os.O_APPEND模式打开的文件，所以在写入文件时会将数据添加到文件尾部
	if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file.Name())
		for i := 0; i < 5; i++ {
			_, err := io.WriteString(file, "hello world!\n")
			if err != nil {
				fmt.Println(err)
			}
		}
		defer file.Close()
	}
}

func main() {
	OsWriteFile()
	time.Sleep(time.Second)
	IoWriteString()
}
