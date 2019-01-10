//读写锁同一时间可以允许多个协程读数据
//但是有且只有一个协程写数据

/*
	type RWMutex
	func (rw *RWMutex) Lock(){}
	func (rw *RWMutex) RLock(){}
	func (rw *RWMutex) RLocker() Locker{}
	func (rw *RWMutex) RUnlock(){}
	func (rw *RWMutex) Unlock(){}
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

//银行的锁换成读写锁，存取钱是写操作，查询余额是读操作
type Bank struct {
	sync.RWMutex
	saving map[string]int //每个账户的存款余额
}

func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

// Deposit 存款
func (b *Bank) Deposit(name string, amount int) {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

// Withdraw 取款，返回实际取到的金额
func (b *Bank) Withdraw(name string, amount int) int {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}
	if b.saving[name] < amount {
		amount = b.saving[name]
	}
	b.saving[name] -= amount

	return amount
}

//Query查询余额
func (b *Bank) Query(name string) int {
	b.RLock()
	defer b.RUnlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}

	return b.saving[name]
}

func main() {

	b := NewBank()
	go b.Deposit("xiaoming", 100)
	go b.Withdraw("xiaoming", 20)
	go b.Deposit("xiaogang", 2000)

	time.Sleep(time.Second)
	print := func(name string) {
		fmt.Printf("%s has: %d\n", name, b.Query(name))
	}

	nameList := []string{"xiaoming", "xiaogang", "xiaohong", "xiaozhang "}
	for _, name := range nameList {
		go print(name)
	}

	time.Sleep(time.Second)
}
