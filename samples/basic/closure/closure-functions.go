package main

import "fmt"

//该函数返回一个匿名函数，该匿名函数使用闭包方式隐藏变量i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	//多次调用nextInt查看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//确认这个状态对于特定的函数来说是唯一的
	newInts := intSeq()
	fmt.Println(newInts())
}
