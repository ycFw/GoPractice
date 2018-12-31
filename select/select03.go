//使用goroutine和select实现fibonacci数列

package main

import "fmt"

func fibonacci(ch, sig chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-sig:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	sig := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%v, ", <-ch)
		}
		fmt.Println("...")
		sig <- 0
	}()
	fibonacci(ch, sig)
}
