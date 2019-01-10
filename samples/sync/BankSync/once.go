//once对象是用来存放1个无入参无返回值的函数，once可以确保这个函数只被执行1次。

package main

import (
	"fmt"
	"sync"
)

// once在10个协程中调用，但once中的函数onceBody()只执行了1次
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only Once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
