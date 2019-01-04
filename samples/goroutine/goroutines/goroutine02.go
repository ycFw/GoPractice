package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	//同步调用
	f("direct")
	//go协程并发执行该函数
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("golang")

	//两个go协程异步运行
	fmt.Scanln()
	fmt.Println("done")
}
