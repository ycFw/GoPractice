package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)

//Cond实现了一个条件变量，一个线程集合地，供线程等待或者宣布某事件的发生。
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知,阻塞当前goroutine
			fmt.Println(x)
			time.Sleep(time.Second * 1)

		}(i)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)
}
