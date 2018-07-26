package main

import (
	"fmt"
	"time"
)

//func work(id int, c chan int) {
//	for {
//		n := <-c
//		fmt.Printf("worker %d recevied %c\n.", id, n)
//	}
//}

//<-chan int  收数据
//chan<- int 发数据

func CreateWorker(id int) chan<- int {

	c := make(chan int)

	go func() {

		for {
			//n := <-c
			fmt.Printf("worker %d recevied %c\n.", id, <-c)
		}

	}()

	return c
}

func chanDemo() {

	var channels [10]chan<- int

	for i := 0; i < 10; i++ {

		//channels[i] = make(chan int)

		channels[i] = CreateWorker(i)
		//go work(i, channels[i])

	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	//n := <-c
	//fmt.Println(n)

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

}

func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %d", id, n)
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferedChannel()
}
