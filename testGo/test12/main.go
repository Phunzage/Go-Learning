// reflect反射

package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
}

func main() {
	var i int8
	reflectType(i)
}
