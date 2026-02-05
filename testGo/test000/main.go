// 临时运行一些无关紧要的测试

package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	// for i := range 7 {
	//     fmt.Scan(&a[i])
	// }
	fmt.Println(a[:3], a[3:6])
}
