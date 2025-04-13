// 结构体练习3
// 景区根据不同游客的年龄收取不同的门票
package main

import (
	"fmt"
	"os"
)

type visitor struct {
	name  string
	age   int
	money int
}

func (v *visitor) getNameAndAge() {
	fmt.Print("请输入姓名: ")
	fmt.Scan(&v.name)
	if v.name == "n" {
		fmt.Print("您已退出")
		os.Exit(0)
	}
	fmt.Print("请输入年龄: ")
	fmt.Scan(&v.age)
}

func (v *visitor) setMoney() {
	v.money = 20
	if v.age <= 18 {
		v.money = 0
	}
}

func main() {
	v1 := new(visitor)
	v1.getNameAndAge()
	v1.setMoney()
	if v1.money == 0 {
		fmt.Printf("%s的年龄为%d, 门票免费\n", v1.name, v1.age)
	} else {
		fmt.Printf("%s的年龄为%d, 门票价格为: 20\n", v1.name, v1.age)
	}

}
