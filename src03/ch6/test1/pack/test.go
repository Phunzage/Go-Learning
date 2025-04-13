// 一个封装案例
// 不能随意查看员工的年龄和工资等隐私，并对输入的年龄进行合理的验证

package test1

import "fmt"

type person struct { //外部包要使用工厂模式访问和赋值
	Name   string
	age    int //其他包不可以直接访问
	salary float64
}

// 工厂模式(相当于构造函数)，用于给Name赋值
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 使用Set和Get访问age和salary
func (p *person) SetPersonAge(age int) {
	//数据验证
	if age < 0 || age >= 130 {
		fmt.Println("请输入合适的年龄")
		return
	}
	p.age = age
}
func (p *person) GetPersonAge() int {
	return p.age
}

func (p *person) SetPersonSalary(salary float64) {
	//数据验证
	if salary < 3000 || salary >= 30000 {
		fmt.Println("请输入合适的薪水")
		return
	}
	p.salary = salary
}
func (p *person) GetPersonSalary() float64 {
	return p.salary
}
