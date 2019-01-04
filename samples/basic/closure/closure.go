//非必要不要在程序中使用闭包
//闭包是一个函数值，他引用了函数体之外的变量，此函数可以对这个变量进行访问和赋值

package main

import "fmt"

func adder01() func(int) int {
	sum := 0
	//返回一个函数值，函数值引用了外部的sum，形成闭包
	return func(x int) int {
		sum += x
		return sum
	}
}

//外边的sum只是调用的时候执行一次
// pes和neg 每次循环给入的变量都在 return func(x int) int内累加
func adder() func(int) int {
	sum := 0
	fmt.Println("外边的sum=", sum)
	return func(x int) int {
		fmt.Print("里边的x=", x, " --> ")
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
