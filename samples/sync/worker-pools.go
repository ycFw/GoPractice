//通过go协程和channel实现一个worker-pool

package main

import (
	"fmt"
	"time"
)

//在多个并发实例中支持的任务，从jobs接收任务，从results发送对应的结果
//每个任务间隔1s来模仿一个耗时的任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(3 * time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动3个worker，初始是阻塞的，因为还没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//发送9个jobs，然后close这些channel来表示这就是所有的任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	//最后收集所有任务的返回值
	for a := 1; a <= 9; a++ {
		<-results
	}

}
