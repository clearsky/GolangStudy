package main

import (
	"fmt"
	"sync"
	"time"
)

type automicInt struct {
	value int
	lock sync.Mutex
}

func (a *automicInt) increment(){
	fmt.Println("safe increment")
	//a.lock.Lock()
	//defer a.lock.Unlock()
	//a.value++
	func(){  // 将mutex控制在一段代码里面
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *automicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	var a automicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
