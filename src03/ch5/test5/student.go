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

type studentMgr struct { //管理学生集合的结构，提供增删改查方法
	allStudent []*student
}

/*
studentMgr 结构的设计好处

封装数据
将学生集合 (allStudent 切片) 和相关操作方法绑定在一起，符合面向对象设计原则。

集中管理
所有学生操作（增删改查）都通过 studentMgr 的方法完成，避免在 main 函数中散布业务逻辑。

容量预分配
通过 make([]*student, 0, 100) 预分配切片容量，减少内存重新分配次数。

扩展性
如需新增功能（如删除学生），只需在 studentMgr 中添加方法，无需修改主程序结构。

数据隔离
主程序 (main.go) 不直接操作学生数据，通过管理器接口访问，提高安全性。
*/

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
