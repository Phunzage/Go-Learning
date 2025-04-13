// 多态数组

package main

import "fmt"

type Animaler interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }
type Bird struct{ Name string }

func (d Dog) Speak() string {
	return d.Name + ": 汪汪"
}

func (c Cat) Speak() string {
	return c.Name + ": 喵喵"
}

func (b Bird) Speak() string {
	return b.Name + ": 啾啾"
}

func main() {
	zoo := [3]Animaler{
		Dog{Name: "xcdog"},
		Cat{Name: "xccat"},
		Bird{Name: "xcbird"},
	}

	for _, v := range zoo {
		fmt.Println(v.Speak())
	}

	// 输出
	// 	xcdog: 汪汪
	// xccat: 喵喵
	// xcbird: 啾啾
}

/*
内部结构

[]Animal 数组
+---+---------------------+
| 0 | 值: Dog{Name:"阿黄"} |
|   | 类型: *Dog          |
+---+---------------------+
| 1 | 值: Cat{Name:"咪咪"} |
|   | 类型: *Cat          |
+---+---------------------+

*/
