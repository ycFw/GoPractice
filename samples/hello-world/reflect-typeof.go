package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf("test reflect"))
	fmt.Println(reflect.TypeOf([]int{1, 2, 3, 4}))
	fmt.Println(reflect.TypeOf(1.23))
}
