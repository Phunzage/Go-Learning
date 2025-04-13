//继承
// 在一个场景中，比如编写了两个结构体（小学生和大学生），两个结构体都有考试中方法，获取成绩方法和展示成绩方法，代码冗余————几个结构体的字段和方法几乎相同，但是我们却写了相同的代码
//此时出现的问题：代码冗余，不利于维护，也不利于功能的扩展

// 为保证简洁性,尽量避免使用多重继承

// 解决方法：继承
// 继承可以使代码复用

// 通过嵌套匿名结构体来实现继承
package main

import "fmt"

//定义父类 学生 结构体
type student struct {
	name  string
	age   int
	score float64
}

// 定义 小学生 结构体，继承 学生 结构体
type pupil struct {
	student
}

// 定义 大学生 结构体，继承 学生 结构体
type graduate struct {
	student
}

// 利用父类显示信息
func (stu *student) ShowInfo() {
	fmt.Printf("%s 的年龄是 %d 成绩是 %.2f\n", stu.name, stu.age, stu.score)
}

//利用父类传入成绩
func (stu *student) SetScore(score float64) {
	stu.score = score
}

// 小学生特有方法
func (p *pupil) Testing() {
	fmt.Println("小学生正在考试……")
}

// 大学生特有方法
func (g *graduate) Testing() {
	fmt.Println("大学生正在考试……")
}

func main() {
	//给小学生直接赋值（字段可见情况下）
	p1 := &pupil{
		student{
			name:  "Tom",
			age:   10,
			score: 90,
		},
	}
	// 利用继承给大学生赋值
	g2 := &graduate{}
	g2.student.name = "Eli"
	g2.student.age = 20

	// 小学生考试 传成绩 打印信息
	p1.Testing()
	p1.SetScore(90)
	p1.ShowInfo()
	// 大学生考试 传成绩 打印信息
	g2.Testing()
	g2.SetScore(95)
	g2.ShowInfo()

}
