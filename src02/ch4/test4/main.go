//结构体案例
// 使结构体里的元素大写

package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName string
	lastName  string
}

func upPerson(p *person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//方法一
	var pers1 person
	pers1.firstName = "Eli"
	pers1.lastName = "David"
	upPerson(&pers1)
	fmt.Printf("The name of the person is \"%s %s\"\n", pers1.firstName, pers1.lastName)

	//方法二
	pers2 := new(person)
	pers2.firstName = "Eli"
	pers2.lastName = "David"
	upPerson(pers2)
	fmt.Printf("The name of the person is \"%s %s\"\n", pers2.firstName, pers2.lastName)

	//方法三
	//和第二种相同，底层仍会调用new()
	pers3 := &person{
		"Eli",
		"David",
	}
	upPerson(pers3)
	fmt.Printf("The name of the person is \"%s %s\"\n", pers3.firstName, pers3.lastName)
}
