// 银行账户
package main

import (
	"fmt"
	"os"
)

// 定义银行账户
type Account struct {
	accountName string
	password    string
	blance      float64
}

// 存款
func (a *Account) Deposite() {
	fmt.Print("请输入密码: ")
	a.TestPwd()
	var bla float64
	fmt.Println("欢迎进入存款界面")
	fmt.Print("请输入存款金额: ")
	fmt.Scan(&bla)
	if bla < 0 {
		fmt.Println("无法存款")
	} else {
		a.blance += bla
		fmt.Println("存款成功")
	}

}

// 取款
func (a *Account) GetMoney() {
	fmt.Print("请输入密码: ")
	a.TestPwd()
	var bla float64
	fmt.Println("欢迎进入取款界面")
	fmt.Print("请输入取款金额: ")
	fmt.Scan(&bla)
	if bla < 0 || bla > a.blance {
		fmt.Println("无法取款")
	} else {
		a.blance -= bla
		fmt.Println("取款成功")
	}
}

// 查看余额
func (a *Account) ShowMoney() {
	fmt.Print("请输入密码: ")
	a.TestPwd()
	fmt.Println("您的余额为: ", a.blance)
}

// 密码检测
func (a *Account) TestPwd() {
	var pwd string
	for i := 3; i > 0; {
		fmt.Scan(&pwd)
		if pwd != a.password {
			fmt.Printf("密码错误, 您还有 %d 次机会\n", (i - 1))
			i--
			if i == 0 {
				fmt.Println("错误次数过多,已强制退出")
				os.Exit(0)
			}
		} else {
			fmt.Println("密码正确")
			return
		}
	}
}

// 修改密码
func (a *Account) ChangePwd() {
	fmt.Print("请输入原始密码: ")
	a.TestPwd()
	fmt.Println("请输入修改后的密码: ")
	fmt.Scan(&a.password)
}

func main() {

	//相当于数据库已存储用户信息
	ur1 := &Account{
		"phunzage",
		"123456",
		1000.0,
	}
	fmt.Printf("欢迎用户 %s 来到我行系统", ur1.accountName)
	var num, i int8
	for {
		fmt.Println(`
	1. 存款
	2. 取款
	3. 显示余额
	4. 修改密码
	5. 退出系统
	`)

		fmt.Print("请输入选项: ")
		fmt.Scan(&num)
		switch num {
		case 1:
			ur1.Deposite()
		case 2:
			ur1.GetMoney()
		case 3:
			ur1.ShowMoney()
		case 4:
			ur1.ChangePwd()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("请输入正确的选项")
		}
		fmt.Print("是否继续业务(1.是 2.否)")
		fmt.Scan(&i)
		if i != 1 {
			fmt.Println("您已退出")
			os.Exit(0)
		}
	}
}
