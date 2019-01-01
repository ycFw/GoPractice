package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello there")
	}()
}
