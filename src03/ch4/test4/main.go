package main

import "fmt"

type animal struct {
	name string
}
type dog struct {
	feet int8
	*animal
}

func (a *animal) move() {
	fmt.Printf("%s会动\n", a.name)
}
func (d *dog) wang() {
	fmt.Printf("%s会叫\n", d.name)
}
func main() {
	d1 := dog{
		feet: 4,
		animal: &animal{
			name: "lucky",
		},
	}
	d1.move()
	d1.wang()
}
