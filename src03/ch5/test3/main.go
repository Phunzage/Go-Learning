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

// 存在的问题
// 1.student.go中全局切片 studentSlice，存在数据竞争风险且封装性差。
// 2.编辑时用索引 i-1，但ID可能不连续。
// 3.直接打印结构体，可读性差。
// 4.default 分支直接退出，应明确处理选项4。
// 5.若写在同一个包内可以不需要引入

// 改进版在 [test5](src03\ch5\test5\main.go)

func ShowMenu() {
	fmt.Println(`
	1. 添加学员信息
	2. 编辑学员信息
	3. 展示所有学员信息
	4. 退出系统
	`)
}

func main() {
	fmt.Println("欢迎来到学员信息管理系统")
	for {
		ShowMenu()
		var a int
		fmt.Print("请输入你的选项: ")
		fmt.Scan(&a)
		switch a {
		case 1:
			addStudent()
			fmt.Println("添加成功")
		case 2:
			editStudent()
			fmt.Println("修改成功")
		case 3:
			fmt.Println("当前存储的学生信息有：")
			fmt.Println("id name class")
			st.ShowStudent()
		default:
			fmt.Println("您已退出")
			os.Exit(0)
		}
	}
}

func addStudent() {
	var (
		id          int
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
	st.SaveStudent(stu)
}

func editStudent() {
	var (
		i           int
		name, class string
	)
	fmt.Print("请输入要修改的学生id：")
	fmt.Scan(&i)
	fmt.Printf("请输入%d号学生修改后的姓名：", i)
	fmt.Scan(&name)
	fmt.Printf("请输入%d号学生修改后的班级：", i)
	fmt.Scan(&class)
	st.ChangeStudent(i-1, name, class)

}
