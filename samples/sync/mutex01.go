//为了保证同一时刻只能有一个goroutine访问共享变量，需要用到互斥和互斥锁

package main

import (
	"fmt"
	"sync"
	"time"
)

//安全计数器，使用sync.Mutex保证安全
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//增加计数器 key 的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	//Lock之后，同一时刻只要一个goroutine可以访问c.v
	c.v[key]++
	c.mux.Unlock()
}

//取得计数器当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	//使用defer语句保证互斥锁一定会被解锁
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	// fmt.Println(c)
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
		time.Sleep(100 * time.Millisecond)
		fmt.Println(c.Value("somekey"))
	}
	fmt.Println(c.Value("somekey"))
}
