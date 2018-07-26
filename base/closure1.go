package main

import (
	"fmt"
)

func adder() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum

	}
}

func main() {
	adder := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(adder(i))
	}
}
