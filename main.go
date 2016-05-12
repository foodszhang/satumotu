package satumotu

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("before")
	log.Println(time.Now())
	log.Println("end")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}

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
func main() {
	h := handler{"123"}
	http.Handle("/", Handle(timeHandle, hello))
	http.Handle("/adapter", Adapt(h, Log()))
	http.ListenAndServe(":8080", nil)
}
