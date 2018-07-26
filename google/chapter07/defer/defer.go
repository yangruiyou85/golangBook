package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hello/google/chapter07/defer/fib"
)

func tryDefer() {
	defer fmt.Println("hello")
	defer fmt.Println("hello golang")

	fmt.Println(2)

}

func WriteFile(fn string) {
	file, err := os.Create(fn)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())

	}
}

func main() {

	//tryDefer()

	WriteFile("fib.txt")

}
