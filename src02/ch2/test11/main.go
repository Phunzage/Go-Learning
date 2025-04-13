//利用闭包检测文件后缀
//检测文件是否有指定后缀, 有的话返回文件名, 没有的话加上后缀返回

package main

import (
	"fmt"
	"strings"
)

func main() {

	f := makeSuffix(".png")

	fmt.Println(f("vscode.png"))

}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
