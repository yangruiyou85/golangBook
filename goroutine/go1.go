package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {

	var a [10]int

	for i := 0; i < 10; i++ {
		go func() {
			for {

				a[i]++
				//fmt.Printf("hello from "+" groutine %d\n", i)
				runtime.Gosched()
			}
		}()

	}
	time.Sleep(time.Millisecond)

	fmt.Println(a)
}
