package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//Golang中最主要的状态管理方式是通过channel的沟通完成的
//worker-pool,sync/atomic
func main() {
	//计数器，无符号整数
	var ops uint64 = 0

	//为了模拟并发更新，启动50个协程，对计数器每1ms进行自增1操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 `AddUint64` 来让计数器自动增加，使用
				// `&` 语法来给出 `ops` 的内存地址。
				atomic.AddUint64(&ops, 1)

				//允许其他go协程的执行,让出时间片
				runtime.Gosched()
			}
		}()
	}

	//等待1s，让ops的自家操作执行一会
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)

}
