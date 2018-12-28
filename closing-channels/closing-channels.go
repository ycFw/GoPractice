// 关闭一个channel意味着不能再向这个channel发送值了
//这个特性用来给这个channel的接收方传达工作已经完成的信息

package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				//more返回true，代表channel未关闭
				fmt.Println("received jobs", j)
			} else { //channel关闭
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//done未接收到值一直阻塞
	<-done
}
