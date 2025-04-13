// ASCII码

package main

import "fmt"

func main() {
	// ASCII码转换
	var a byte = 'a'

	fmt.Printf("%c对应的ASCII数字为：%d\n", a, a)

	var z rune = '庞'

	fmt.Printf("%c对应的ASCII数字为：%d", z, z)
}
