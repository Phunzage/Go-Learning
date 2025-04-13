// 结构体练习1
//定义一个学生结构体

package main

import "fmt"

type student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

//各字段组合方法
func (s *student) say() (res string) {
	res = fmt.Sprintf("name = %s, gender = %s, age = %d, id = %d, score = %.2f", s.name, s.gender, s.age, s.id, s.score)
	return
}

func main() {
	s1 := &student{"小明", "男", 20, 1, 90.8}
	res := s1.say()
	fmt.Println(res)
}
