// 当使用 for 循环遍历切片时，每次迭代都会返回两个值。 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	for i, num := range nums {
		fmt.Printf("索引: %d, 值: %d\n", i, num)
		num *= 2 // 这只会修改副本，不会影响切片中的原始值
	}

	fmt.Println(nums) // 输出: [10 20 30] (未被修改)

	// 如果要修改原始切片中的值，需要通过索引
	for i := range nums {
		nums[i] *= 2
	}

	fmt.Println(nums) // 输出: [20 40 60] (已被修改)
}
