package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gengerator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		   out <- i
		   i ++
		}
	}()
	return out
}

func worker(id int, c chan int){
	for n := range c{  //自动等待数据发完
		//n, ok := <-c  //ok判断是否还有值
		//if !ok{
		//	break
		//}
		time.Sleep(1 * time.Second)  // 消耗数据太慢会被覆盖
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int)chan<-int{  // 加入<-表示send only， <-chan int 表示read only
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var c1, c2 chan int  // c1 and c2 = nil
	c1, c2 := gengerator(), gengerator()
	worker := createWorker(0)
	n := 0
	//hasValue := false  //为了让n刷新
	var values []int  //缓冲数据
	tm := time.After(10 * time.Second) //10s后tm会有数据
	tick := time.Tick(time.Second)  //每隔一段时间送一个值
	for{
		var activeWorker chan<- int
		var activeValue int
		//if hasValue{
		//	activeWorker = worker
		//}
		if len(values) > 0{
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <- c1:
			//hasValue = true
			values = append(values, n)
		case n = <- c2:
			//hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			//hasValue = false
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 以每个select计算时间
			fmt.Println("time out")
		case <-tick:
			fmt.Println("queue len=", len(values))
		case <- tm:
			fmt.Println("Bye")
			return
		}
	}

}
