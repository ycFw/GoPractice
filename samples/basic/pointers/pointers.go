package main

import "fmt"

//zeroval在main函数中不能改变i的值，但是zeroptr可以
//因为zeroptr有这个变量的内存地址的引用

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i) //1

	zeroval(i)
	fmt.Println("zeroval: ", i) //1

	zeroptr(&i)
	fmt.Println("zeroptr: ", i) //0

	fmt.Println("pointer: ", &i)
}
