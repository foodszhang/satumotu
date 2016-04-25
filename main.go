package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type handler struct {
	name string
}

// HandlerFunc is the interface for handle function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Adapter is a wrapper for handler
type Adapter func(http.Handler) http.Handler

// Log give some log
func Log() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before")
			defer log.Println("end")
			h.ServeHTTP(w, r)
		})
	}
}

// Adapt wrap all adaters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for i, _ := range adapters {
		h = adapters[len(adapters)-i-1](h)
	}
	return h

}

// Handle combine handlers to one handler
func Handle(handlers ...HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, handler := range handlers {
				handler(w, r)
			}
		})
}
func timeHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("before")
	log.Println(time.Now())
	log.Println("end")
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func main() {
	h := handler{"123"}
	http.Handle("/", Handle(timeHandle, hello))
	http.Handle("/adapter", Adapt(h, Log()))
	http.ListenAndServe(":8080", nil)
}
