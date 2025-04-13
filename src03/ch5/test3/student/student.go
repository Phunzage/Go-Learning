package st

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

var studentSlice []student

func SaveStudent(x *student) {
	studentSlice = append(studentSlice, *x)
}

func ShowStudent() {
	for _, v := range studentSlice {
		fmt.Println(v)
	}
}

func ChangeStudent(i int, name, class string) {
	studentSlice[i].name = name
	studentSlice[i].class = class
}

func NewStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}
