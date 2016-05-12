package test

import (
	"fmt"
	"github.com/foodszhang/trie_router"
	"log"
	"net/http"
	"satumotu"
	"testing"
	"time"
)

func timeHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("before time", time.Now())
	log.Println(time.Now())
	log.Println("end time", time.Now())
	fmt.Fprintf(w, "%s", time.Now())
}

type handler struct {
	name string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// Log give some log
func Log() router.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			end := time.Now()
			//[LOG][Date][Method][URL][in xx msec]
			log.Printf("[LOG][%s][%s][%s][in %d msec]", end, r.Method, r.URL.String(), (end.UnixNano()-start.UnixNano())/1000)
		})
	}
}
func TestMiddleWare(t *testing.T) {
	h := handler{"123"}
	http.Handle("/", satumotu.Handle(timeHandle, hello))
	http.Handle("/adapter", router.Adapt(h, Log()))
	//http.ListenAndServe(":8080", nil)
}
