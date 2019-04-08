package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int){
	for n := range c{  //自动等待数据发完
		//n, ok := <-c  //ok判断是否还有值
		//if !ok{
		//	break
		//}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int)chan<-int{  // 加入<-表示send only， <-chan int 表示read only
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo(){
	//var c chan int // c == nil

	var channels [10]chan<- int
	for i:=0;i<10;i++{
		channels[i] = createWorker(i)
	}

	//c <- 1  //有人取走值后才能下一次的<-，否则阻塞
	//c <- 2
	for i:=0;i<10;i++{
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	c := make(chan int, 3)  //可以设置缓冲区，缓冲区为3，可以放入3个值
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}

func channelClose(){  //发送方进行CloseF
	c := make(chan int)  //可以设置缓冲区，缓冲区为3，可以放入3个值
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channle as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
