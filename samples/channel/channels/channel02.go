package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main01() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func main() {
	//设置channel缓冲，缓冲区满才会阻塞，缓冲区空的时候，接受操作会阻塞
	ch := make(chan int, 2)
	ch <- 100
	ch <- 5
	// ch <- 60
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)

}
