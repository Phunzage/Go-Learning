package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {

	p1 := person{"Tom", 20}
	// p1.add("John", 30) //方法周期结束后被释放
	fmt.Printf("name = %s, age = %d\n", p1.name, p1.age)
	//输出：name = Tom, age = 20
}

//值接收组合，不会修改值，传递的是副本
// func (p person) add(name string, age int8) {
// 	p.name = name
// 	p.age = age
// }

/*
什么时候使用指针类型：
1.需要修改接收者的值
2.接受这是拷贝代价比较大的大对象
3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
*/
