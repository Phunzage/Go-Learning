// test3改进版（教师版）

package main

import (
	"fmt"
	"os"
)

// 打印选项
func ShowMenu() {
	fmt.Println(`
	1. 添加学员信息
	2. 编辑学员信息
	3. 展示所有学员信息
	4. 退出系统
	`)
}

// 获取用户输入，在student.go中，newStudent返回一个包含学生信息的地址
func getInput() *student {
	var (
		id          int
		name, class string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生班级：")
	fmt.Scanf("%s\n", &class)
	// 调用newStudent的构造函数，将包含学生的学生加入student结构体中，返回存有学生信息的地址
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	// sm是存储学生信息的切片，后续student都加入到里面
	sm := newStudentMgr()

	fmt.Println("欢迎来到学员信息管理系统")
	for {
		// 打印菜单系统
		ShowMenu()
		// 等待用户要执行的选项
		var input int
		fmt.Print("请输入你的选项: ")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			stu := getInput()
			// 接收用户输入的学生信息的地址，加入到切片
			sm.addStudent(stu)
			fmt.Println("添加成功")
		case 2:
			stu := getInput()
			// 通过用户输入的id比对，用新学生对象完全替换旧对象
			sm.editStudent(stu)
			fmt.Println("修改成功")
		case 3:
			fmt.Println("当前存储的学生信息有：")
			sm.showStudent()
		case 4:
			fmt.Println("您已退出")
			os.Exit(0)
		default:
			fmt.Println("您输入的选项有误，请重新输入")
			return
		}
	}
}
