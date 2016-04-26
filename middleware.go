package satumatu

import (
	"net/http"
)

type handler struct {
	name string
}

// HandlerFunc is the interface for handle function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Adapter is a wrapper for handler
type Adapter func(http.Handler) http.Handler

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
