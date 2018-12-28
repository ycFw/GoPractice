package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob", age: 20})

	fmt.Println(person{name: "Fred"})

	// `&` 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的
	sp.age = 51
	fmt.Println(sp.age)
}
