package test

import (
	"fmt"
	"github.com/foodszhang/trie_router"
	"github.com/gorilla/context"
	"net/http"
	"satumotu"
	"testing"
)

type UserView struct {
	satumotu.Controller
}

func (c *UserView) get(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").([]router.Param)
	for name, value := range params {
		fmt.Fprintf(w, "%s=%s,", name, value)
	}
	fmt.Fprintf(w, "\n")
}

func TestController(t *testing.T) {
	root := satumotu.NewRouter(false)
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
