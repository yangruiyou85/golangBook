package main

import (
	"fmt"

	"time"

	"github.com/hello/retriever/mock"
	"github.com/hello/retriever/real"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

func Post(poster Poster) string {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func Download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func Inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever1:
		fmt.Println("contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}

	fmt.Printf("%T %v\n", r, r)

}

func main() {

	var r Retriever
	r = mock.Retriever1{"this is a fake imooc.com"}

	//fmt.Printf("%T %v\n", r, r)
	Inspect(r)
	r = &real.Retriever{

		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	Inspect(r)
	//fmt.Printf("%T  %v\n", r, r)

	//fmt.Println(Download(r))

	if mockRetriever, ok := r.(mock.Retriever1); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever.")
	}

}
