package main

import (
	"net/http"
	"fmt"
	"time"
	"runtime"
	"reflect"
	"log"
)

func main() {
	handle := myhandle{}
	handleHello := hello{}
	handleWorld := world{}
	s := &http.Server{
		Addr: ":7878",
		//Handler:        &handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/", &handle)
	http.Handle("/hello/", logHandler(&handleHello))
	http.Handle("/world", &handleWorld)
	log.Fatal(s.ListenAndServe())
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T", h)
		h.ServeHTTP(w, r)
	})
}

func logHTTP(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler func called -", name)
		h(w, req)

	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world...")
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello ...")
}
func worldHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "world ...")
}

type myhandle struct{}

func (h *myhandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world...")
}

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello ...")
}

type world struct{}

func (h *world) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "world...")
}
