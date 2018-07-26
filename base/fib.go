package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fib() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func main() {

	f := Fib()

	PrintFileContents(f)

	suf := ".jpg"
	cc := makeSuffix(suf)
	fmt.Println(cc("kkgo"))
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {

	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix

		}
		return name
	}
}
