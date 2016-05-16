package satumotu

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
	"testing"
)

func printNameHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Fprintf(w, "Hello, %s!\n", query.Get("name"))
}
func printParams(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(map[string]string)
	for name, value := range params {
		fmt.Fprintf(w, "%s=%s,", name, value)
	}
	fmt.Fprintf(w, "\n")
}

func TestRouter(t *testing.T) {
	root := NewRouter(false)
	root.Get("/", hello)
	root.Get("/log", hello, Log())
	root.Get("/time", timeHandle)
	root.Get("/time/log", timeHandle, Log())
	root.Get("/name", printNameHandle)
	root.Get("/name/log", printNameHandle, Log())
	root.Get("/params/<string:name>/<int:id>", printParams)
	root.Get("/params/log/<string:action>", printParams, Log())
	http.ListenAndServe(":8080", &root)
}
