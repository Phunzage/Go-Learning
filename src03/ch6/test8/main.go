package main

import (
	"fmt"
	"math/rand"

	// 引入sort包
	"sort"
)

// 定义学生结构体
type Student struct {
	Name string
	Age  int
}

// 定义学生结构体切片
type StuSli []Student

// 定义三个方法
// 为什么要定义三个方法：sort包已经整合了实现排序的内容（如何实现不用去管），我们只需满足sort包的要求即可使用sort包，sort的要求是定义一个切片长度的方法，一个比大小的方法，一个交换的方法，三个定义好后，sort会把这三个方法自动通过内部代码完成整合，输出排好序的切片
// Len方法
func (stu StuSli) Len() int {
	return len(stu)
}

// Less方法
// 当前面的age小于后面的age为真，返回真，不改变排序
// 否则返回假，会被排序
func (stu StuSli) Less(i, j int) bool {
	return stu[i].Age < stu[j].Age
}

// Swap方法
func (stu StuSli) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {

	// int切片排序演示
	var nums = []int{-2, 0, 5, 32, 17, 11}
	// 自带的排序
	sort.Ints(nums)
	fmt.Println(nums)
	/*
		输出[-2 0 5 11 17 32]
	*/

	// 自定义结构体类型排序
	// Class3是一个学生结构体切片变量
	var Class3 StuSli
	for i := 0; i < 10; i++ {
		// 为结构体的每个字段赋值
		student := Student{
			Name: fmt.Sprintf("Hero%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 把每个结构体添加进切片
		Class3 = append(Class3, student)
	}

	fmt.Println("--------排序前--------")
	for _, v := range Class3 {
		fmt.Println(v)
	}
	fmt.Println("--------排序后--------")
	sort.Sort(Class3)
	for _, v := range Class3 {
		fmt.Println(v)
	}
}
