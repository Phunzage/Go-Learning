// 文件操作5

package main

import (
	"fmt"
	"os"
)

// 重命名文件
func rename() {
	err := os.Rename("MV.txt", "mv.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("重命名成功")
	}
}

// 删除单个文件
func remove() {
	// 删除当前目录下所有的文件与子目录
	err := os.Remove("move.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功")
	}
}

// 删除目录
func removeall() {
	// 删除当前目录下所有的文件与子目录
	err := os.RemoveAll(".")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除成功")
	}
}

func main() {
	rename()
	remove()
}
