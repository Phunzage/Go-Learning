package st

type Student struct {
	Id    int8
	Name  string
	Class string
}

func NewStudent(id int8, name, class string) *Student {
	return &Student{
		Id:    id,
		Name:  name,
		Class: class,
	}
}
