package main

import (
	"fmt"
	"time"
)

func timenow() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午好")
	case t.Hour() < 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}
