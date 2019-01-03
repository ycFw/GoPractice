- 这个点之前一直没弄透彻，今天弄懂了，写个md防止自己忘了。。。

- sync中的状态量cond不是开箱即用的，需要用sync.NewCond函数创建它的指针值
  
- 比如代码中例子，分别初始化sendCond和recvCond

        sendCond := sync.NewCond(&lock)
        recvCond := sync.NewCond(lock.RLocker())

- newcond传入的参数其实是Locker接口，该接口中定义了lock()和unlock()方法。只要某类型或参数实现了该方法，就说明实现了Locker接口

        type Locker() interface {
            Lock()
            Unlock()
         }
 - sync.mutex 和 sync.rwmutex 这两个类型都有lock()和unlock()方法，但是他们都是指针方法

        func(m *Mutex) Lock()
        func(m *RWMutex) Lock() ...
- 因此，sync.mutex和sync.rwmutex类型的指针类型才是sync.locker接口的实现类型

        //需要加上地址符取地址
        var lock sync.RWMutex   
        //把基于lock变量的指针值传给了sync.NewCond
        func (rw *RWMutex) RLocker() Locker 
                
