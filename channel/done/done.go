package main

import (
	"fmt"
	"sync"
)

//<-chan int  收数据
//chan<- int 发数据

type worker struct {
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}

func doworker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {

		fmt.Printf("worker %d received %c\n", id, n)

		wg.Done()

	}
}

func CreateWorker(id int, wg *sync.WaitGroup) worker {

	w := worker{in: make(chan int), wg: wg,}

	go doworker(id, w.in, wg)
	return w

}

func chanDemo() {

	var wg sync.WaitGroup

	var workers [10]worker

	for i := 0; i < 10; i++ {

		workers[i] = CreateWorker(i, &wg)

	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done

	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	wg.Done()
	wg.Wait()

}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

}

func main() {
	chanDemo()
	//bufferedChannel()
}
