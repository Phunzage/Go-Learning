// 接口的传值

package main

import "fmt"

type Usb interface {
	Say()
}
type Stu struct{}

func (s *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	// 方法传地址，此处使用&stu
	var u Usb = &stu
	u.Say()
	fmt.Println("here", u)
}
