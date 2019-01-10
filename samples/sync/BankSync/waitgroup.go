//等待组WaitGroup
/*
	type WaitGroup
	func (wg *WaitGroup) Add(delta int){}
	func (wg *WaitGroup) Done(){}
	func (wg *WaitGroup) Wait(){}
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func leader() {
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		go follower(&wg, i)
	}
	wg.Wait()

	fmt.Println("open the box together!!!")
}

func follower(wg *sync.WaitGroup, id int) {
	fmt.Printf("follower %d find key\n", id)
	wg.Done()
}

func main01() {
	leader()
}
