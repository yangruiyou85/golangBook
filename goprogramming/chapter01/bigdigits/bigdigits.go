package main

import (
	"os"
	"fmt"
	"path/filepath"
	"gopkg.in/go-module/log.v1"
)

func main() {
	var bigDigits = [][]string{
		{"  000  ",
			" 0    0 ",},
	}

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal(line)
			}
		}
		fmt.Println(line)
	}
}
