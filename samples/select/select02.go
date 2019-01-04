package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	//非阻塞发送
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("bo message sent")
	}

	//在messages和signals这两个channel同时使用非阻塞的接收操作
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("bo activity")
	}

}
