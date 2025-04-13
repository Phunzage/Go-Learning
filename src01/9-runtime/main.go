package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOARCH
	// switch os {
	// case "darwin":
	// 	fmt.Println("macOS")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Println(os)
	// }
	fmt.Println(os)
}
