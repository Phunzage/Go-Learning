//二维数组的转置

package main

import "fmt"

type array2 [][]int

func main() {
	a := array2{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	a.trans()
	fmt.Println(a)
}

func (a array2) trans() {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}
}
