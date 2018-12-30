//利用goroutine实现获取和格式化当前时间，通过channel返回到主函数并打印

package main

import (
	"fmt"
	"time"
)

//获取格式化当前时间
func timenow(ch chan string) {
	tn := time.Now().Format("2006年01月02日 15点04分05秒.0000000 时区-0700")
	ch <- tn
}

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go timenow(ch)
		fmt.Println(<-ch)
		time.Sleep(500 * time.Millisecond)
	}
}
