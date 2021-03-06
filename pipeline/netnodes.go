package pipeline

import (
	"net"
	"bufio"
)

func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()

		WriterSink(writer, in)

	}()

}

func NetworkSource(addr string, in <-chan int) {
	out := make(chan int, 1024)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)

		}
		r := ReaderSource(bufio.NewReader(conn), -1)

		for v := range r {
			out <- v

		}
		close(out)
	}()
}
