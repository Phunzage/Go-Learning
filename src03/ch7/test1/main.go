// time库

package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前时间
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 返回年月日时分秒
	fmt.Println(year, month, day, hour, minute, second)

	// 时间戳，从1970.1.1到现在的一个秒数
	timeStamp := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp, timeStamp2)

	// 时间Add方法
	// time.Duration表示1纳秒，time.Secend表示1秒
	later := now.Add(time.Hour)
	fmt.Println(later)

	// 时间格式化
	// Go诞生时间：2006.01.02.15.04.05
	// 时间格式化输出
	t := now.Format("2006-01.02 15~04  05")
	fmt.Println(t) // 2025-04.14 14~12  06

	TickDemo()
}

// 定时器
// 每秒执行
func TickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}
