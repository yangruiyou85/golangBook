package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {

	c := make(chan os.Signal)

	fmt.Println(os.Getpid())

	signal.Notify(c, syscall.SIGTERM, os.Kill, os.Interrupt)
	fmt.Println("boot......")
	for {
		select {
		case s := <-c:
			println(s.String())
			switch s.String() {
			case "terminated":
				fmt.Println("term...")
			case "interrupt":
				fmt.Println("interrupt...")
				continue

			}
		}
	}
	fmt.Println("exit...")
}
