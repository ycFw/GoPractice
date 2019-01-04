//定时器timer，打点器ticker

package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器等待2s
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer1 expired")

	//单纯的等待，需要使用time.sleep
	//定时器是有用原因之一就是你可以在定时器失效之前，取消这个定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()

	//在没失效之前就停止了
	stop2 := timer2.Stop()
	if stop2 { //true,停止成功
		fmt.Println("Timer2 stopped")
	}
}
