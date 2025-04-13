// 学员信息管理系统
package main

import (
	"fmt"
	"os"
	st "test/src03/ch5/test3/student"
)

//需求
/*
1.添加学员信息
2.编辑学员信息
3.展示所有学员信息
*/

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
}

func main() {
	showMenu()
	var a int8
	fmt.Print("请输入你的选项: ")
	fmt.Scan(&a)
	switch a {
	case 1:
		addStudent()
	case 2:
		editStudent()
	case 3:
		showStudent()
	default:
		fmt.Println("您已退出")
		os.Exit(0)
	}
}

func addStudent() int8 {
	var (
		id          int8
		name, class string
	)
	fmt.Println("请输入学员信息")
	fmt.Print("请输入学员id: ")
	fmt.Scan(&id)
	fmt.Print("请输入学员name: ")
	fmt.Scan(&name)
	fmt.Print("请输入学员class: ")
	fmt.Scan(&class)
	stu := st.NewStudent(id, name, class)
	fmt.Println(*stu)
	return 0
}

func editStudent() {
	//编辑学员信息

}
func showStudent() {

}
