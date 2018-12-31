package main

import (
	"fmt"
	"time"
)

//实现一个爆破过程，有读秒、爆炸和过程
//每半秒钟读秒一次；0.25秒一个记录过程；最后爆炸

func main() {
	//定义两个定时器，分别设定读秒时长和爆炸时长
	tick := time.Tick(500 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
