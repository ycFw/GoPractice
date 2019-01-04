//前面的例子，使用原子操作来管理简单的计数器，现在使用互斥锁来在go协程间安全的访问数据
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	//这里的mutex将同步对state的访问
	var mutex = &sync.Mutex{} //sync.Mutex本身是一个struct

	//ops记录对state的操作次数
	var ops int64 = 0

	//运行100个go协程来重复读取state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				//让出时间片
				runtime.Gosched()
			}
		}()
	}

	//同样的，运行10个go协程来模拟写入操作，使用和读取相同的模式
	for w := 0; w < 10; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}

	//让这10个go协程对state和mutex的操作运行1s
	time.Sleep(time.Second)

	//获取并输出最终的操作计数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)

	//对state使用一个最终的锁，显示他是如何结束的
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
