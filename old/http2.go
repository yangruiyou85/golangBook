package main

import (
	"sync"
)

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool
}
type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		w.Header().Set("Connection", "close")
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServerHTTP(w, r)
}


func (mux *ServeMux)Hanler(r *Request )(h Handler,pattern string){

}

func (mux *ServeMux)handler(host,path string)(h Handler,pattern string){

}