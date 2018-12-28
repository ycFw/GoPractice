package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "pings"
	}()

	msg := <-messages
	fmt.Println(msg)
}
