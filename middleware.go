package satumotu

import (
	"net/http"
)

// Adapt wrap all adaters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for i := range adapters {
		h = adapters[len(adapters)-i-1](h)
	}
	return h

}

// Handle combine handlers to one handler
func Handle(handlers ...http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, handler := range handlers {
				handler(w, r)
			}
		})
}
