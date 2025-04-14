package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

// newStudent 是student类型的构造函数，用于创建学生信息，返回存放学生信息的地址
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理的类型，切片类型，用于存储每个学生信息的地址
type studentMgr struct {
	allStudent []*student
}

// newStudentMgr 是studentMgr类型的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudent: make([]*student, 0, 100),
	}
}

// 1.添加学生，接收用户输入的学生信息的地址，加入到切片
func (s *studentMgr) addStudent(newstu *student) {
	s.allStudent = append(s.allStudent, newstu)
}

// 2.编辑学生，通过id比对，在切片找到学生后修改信息
func (s *studentMgr) editStudent(newstu *student) {
	for i, v := range s.allStudent {
		// 用户输入的id与存储的id对比
		if newstu.id == v.id {
			s.allStudent[i] = newstu
			return
		}
	}
	fmt.Printf("您输入的信息有误，系统中没有学号为：%d的学生\n", newstu.id)
}

// 3.展示学生，格式化输出，有语法糖
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学生学号：%d，姓名：%s，班级：%s\n", v.id, v.name, v.class)
	}
}
