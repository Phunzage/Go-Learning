// 数组与切片

package main

import (
	"fmt"
)

func main() {

	// 数组定义方法
	// 方法1：var声明，零值初始化
	var arr1 [5]int // 5个元素的int数组，初始化为[0,0,0,0,0]

	// 方法2：字面量初始化
	arr2 := [3]int{1, 2, 3}   // 显式指定长度并初始化
	arr3 := [...]int{1, 2, 3} // 编译器推导长度，结果是[3]int

	// 方法3：指定索引初始化
	arr4 := [5]int{1: 10, 3: 30} // [0, 10, 0, 30, 0]

	fmt.Println("-----------数组------------")
	fmt.Println(arr1, "\n", arr2, "\n", arr3, "\n", arr4)
	// 注意：定义数组时，必须指定其长度或者大小
	var person = [3]string{"som", "tom", "john"}
	// 数组定义后可以给已有元素赋其他值，但长度不可变
	person[2] = "poter"
	fmt.Println(person)

	// 在此基础上引入切片（Slice），切片（Slice）相较于数组更灵活，因为在声明切片后其长度是可变的

	// 定义一个切片

	// 方法1：var声明，初始为nil切片
	var slice1 []int // nil切片，长度和容量为0

	// 方法2：字面量初始化
	slice2 := []int{1, 2, 3} // 长度和容量都为3的切片

	// 方法3：make创建
	slice3 := make([]int, 5)    // 长度和容量都为5
	slice4 := make([]int, 3, 5) // 长度3，容量5

	// 方法4：从数组或切片切割
	arr := [5]int{1, 2, 3, 4, 5}
	// 容量是指从索引1开始到底层数组末尾
	slice5 := arr[1:4] // [2,3,4]，长度3，容量4

	fmt.Println("-----------切片------------")
	fmt.Println(slice1, "\n", slice2, "\n", slice3, "\n", slice4, "\n", slice5)

	var nameSlice []string
	ageSlice := make([]int, 5)

	fmt.Println(nameSlice)
	// 通过append添加切片元素
	nameSlice = append(nameSlice, "zhangsan", "lisi")
	fmt.Println(nameSlice)

	// 遍历切片并赋值
	for i := range ageSlice {
		ageSlice[i] = i + 1
	}
	fmt.Println(ageSlice)
}
