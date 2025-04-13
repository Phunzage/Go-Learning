//封装练习
// 使用工厂方法和Set Get方法编写账户创建封装案例

package main

import (
	"fmt"
	test2 "test/src03/ch6/test2/pack"
)

func main() {
	//工厂模式
	acc := test2.NewAccount("DavidEli", "135790", 666)
	if acc != nil {
		fmt.Println("创建成功", acc)
	} else {
		fmt.Println("创建失败")
	}

	//使用Set, Get方法为单独字段赋值
	acc1 := test2.NewAccount("TomJohn", "222333", 5000)
	fmt.Printf("%#v\n", acc1) //输出&test2.Account{account:"TomJohn", password:"222333", salary:5000}

	//用Set方法为account赋值
	acc1.SetAccountname("phunzage")
	fmt.Printf("%#v\n", acc1) //输出&test2.Account{account:"phunzage", password:"222333", salary:5000}

	// 用Set方法为possword和salary赋值
	acc1.SetAccountPassword("123456")
	fmt.Println(acc1.GetAccountPassword()) //123456
	acc1.SetAccountSalary(3000)
	fmt.Println(acc1.GetAccountSalary()) //3000

	// 输出全部赋值完后的acc1
	fmt.Printf("%#v", acc1) //输出&test2.Account{account:"phunzage", password:"123456", salary:3000}

}
