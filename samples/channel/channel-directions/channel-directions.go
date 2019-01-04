package main

import "fmt"

//ping函数定义了一个只允许发送数据的channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong函数允许pings这个channel接收数据，同时pongs这个channel发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
