package test

//此处结构体名小写,可使用工厂模式调用
type student struct {
	Name  string //若此处Name和Age首字母小写,外部仍然不可调用,比如main中的stu.age不可访问
	Age   int8
	class string
}
type Cat struct {
	Color string
	age   int8
}

func NewStudent(name, class string, age int8) *student {
	return &student{
		Name:  name,
		Age:   age,
		class: class,
	}
}

// 针对结构体内部首字母小写情况,使用Get方法获取
// 核心思路: class不可访问,在同一包内定义一个方法,此方法可被外部访问,使用此方法获取class值
func (s *student) GetClass() string {
	return s.class
}
