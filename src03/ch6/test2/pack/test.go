package test2

import "fmt"

type account struct {
	accountname string
	password    string
	salary      float64
}

//工厂模式
func NewAccount(accountname, password string, salary float64) *account {
	// 检测账号数据
	if len(accountname) < 6 || len(accountname) > 10 {
		fmt.Println("账号长度要在6-10之间")
		return nil
	}

	// 检测密码数据
	if len(password) != 6 {
		fmt.Println("密码必须是6位")
		return nil
	}

	//检测薪水数据
	if salary < 20 {
		fmt.Println("余额要大于20")
		return nil
	}

	//全部无误后赋值
	return &account{
		accountname: accountname,
		password:    password,
		salary:      salary,
	}
}

//使用Set和Get方法获取和读取account
func (a *account) SetAccountname(accountname string) {
	if len(accountname) < 6 || len(accountname) > 10 {
		fmt.Println("账号长度要在6-10之间")
		return
	}
	a.accountname = accountname
}

func (a *account) GetAccountAccount() string {
	return a.accountname
}

//使用Set和Get方法获取和读取password

func (a *account) SetAccountPassword(password string) {
	if len(password) != 6 {
		fmt.Println("密码必须是6位")
		return
	}
	a.password = password
}

func (a *account) GetAccountPassword() string {
	return a.password
}

//使用Set和Get方法获取和读取salary

func (a *account) SetAccountSalary(salary float64) {
	if salary < 20 {
		fmt.Println("余额要大于20")
		return
	}
	a.salary = salary
}

func (a *account) GetAccountSalary() float64 {
	return a.salary
}
