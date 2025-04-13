//结构体叠加

package main

import "fmt"

//值组合
func main() {
	type personStart struct {
		name string
		age  int8
	}

	type personMiddle struct {
		ps     personStart
		gender string
	}
	type personEnd struct {
		pm           personMiddle
		email, phone string
	}

	pe := personEnd{
		pm:    personMiddle{personStart{"xcmy", 18}, "man!!!"},
		email: "123@123.com",
		phone: "12345678",
	}

	fmt.Printf("%#v\n", pe)
}
