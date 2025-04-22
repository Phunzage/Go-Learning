// 文件操作1

// 文件打开有两种方式

package main

import (
	"fmt"
	"os"
)

func FileOpen1() {

	// 1.方式一，使用 os.Open 直接提供对应的文件名即可（只读，无法修改）
	// 打开文件
	file1, err := os.Open("../test3/test3.txt")
	if os.IsNotExist(err) { // os内置的“文件不存在”参数
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件读取成功")
	}

	fmt.Println(file1) // 输出&{0xc0000746c8}
	// 文件file就是一个指针

	// 一定要记得关闭文件
	defer file1.Close()
}

func FileOpen2() {
	// 2.方式二，以读写模式打开一个文件，使用 os.OpenFile，权限为0666，表示为所有人都可以对该文件进行读写，且不存在时会自动创建。
	file2, err := os.OpenFile("OpenFile.txt", os.O_RDWR|os.O_CREATE, 0666)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file2.Name())
	}

	defer file2.Close()
}

func Stat() {
	// 只获取文件的一些信息，并不读取文件
	file3, err := os.Stat("Openfile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		res := fmt.Sprintf("%+v", file3)
		fmt.Println(res)
	}
}

func main() {
	FileOpen1()
	FileOpen2()
	Stat()
}
