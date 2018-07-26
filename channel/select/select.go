package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	var c1, c2 = generator(), generator()

	var worker = CreateWorker(0)

	var values []int

	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	//n := 0
	//hasValue := false

	for {

		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
			//exceed 800ms
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")

		case <-tick:
			fmt.Println("queeue len=", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
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
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

			out <- i
			i++
		}

	}()
	return out
}

func woker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d \n", id, n)
	}
}


//select使用
//定时器的使用
//在select中使用new channel




