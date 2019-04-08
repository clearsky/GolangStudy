package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker){
	for n := range w.in{  //自动等待数据发完
		//n, ok := <-c  //ok判断是否还有值
		//if !ok{
		//	break
		//}
		fmt.Printf("Worker %d received %d\n", id, n)
		//go func() {
		//	done <- true
		//}()
		w.done()

	}
}

type worker struct {
	in chan int
	//done chan bool
	done func()
}


func createWorker(id int, wg *sync.WaitGroup)worker{  // 加入<-表示send only， <-chan int 表示read only
	w := worker{
		in : make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo(){
	//var c chan int // c == nil

	var wg sync.WaitGroup
	var workers [10]worker
	for i:=0;i<10;i++{
		workers[i] = createWorker(i, &wg)
	}
	//c <- 1  //有人取走值后才能下一次的<-，否则阻塞
	//c <- 2
	for i, worker := range workers{
		worker.in <- 'a' + i
		wg.Add(1)
	}
	wg.Wait()

	//// wait for all of them
	//for _, worker := range workers{  //类似python的join
	//	<-worker.done
	//}
}


func main() {
	chanDemo()
}
