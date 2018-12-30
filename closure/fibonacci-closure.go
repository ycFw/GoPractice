package main

import "fmt"

//fibonacci函数返回一个int类型的函数
func fibonacci() func() int {
	x1, x2 := 0, 1
	sum := 0
	return func() int {
		sum = x1 + x2
		x1, x2 = x2, sum
		return sum
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), ", ")
	}
	fmt.Println("...")
}
