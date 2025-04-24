// 文件操作4

// 复制文件
package main

import (
	"fmt"
	"io"
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

// 方法二，使用io.Copy，一边读一边写，先将内容读到缓冲区中，再写入到目标文件中，缓冲区默认大小为 32KB。

func IoCopy() {
	// 以只读的方式打开源文件
	from, err := os.OpenFile("from.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer from.Close()

	// 以只写的方式打开目标文件
	to, err := os.OpenFile("to.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer to.Close()

	// 复制
	written, err := io.Copy(to, from)
	if err != nil {
		fmt.Println(err)
	} else {
		// 输出22，在Windows环境下，若文件以文本模式打开，\r\n会被转换为\n读取，再转换回\r\n写入。Go的os.OpenFile默认使用二进制模式，不进行换行符转换，直接按原样复制字节
		// 所以：前两行每行占6（数字）+2（换行符）=8字节，共16字节。第三行仅6字节。总计：16 + 6 = 22字节
		fmt.Println(written)
	}
}

func main() {
	// RWFiles()
	IoCopy()
}
