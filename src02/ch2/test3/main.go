// 空白符_

package main

import "fmt"

func main() {

	//丢弃不需要的值
	i, _, k := sum(5, 10)

	fmt.Println(i, k)

}

func sum(num1, num2 int) (sum1, sum2, sum3 int) {
	sum1 = num1 + num2
	sum2 = num1 * num2
	sum3 = num1 + num2 + 2
	return
}
