package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	s := `asdfasdf
asdfasdf


adsfsdfasdf
dkjf打卡 跨世纪的反馈
ddd`
	PrintFileContents(strings.NewReader(s))

}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	PrintFileContents(file)

}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
