package satumotu

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
	"testing"
)

type UserView struct {
	Controller
}

func (c *UserView) GET(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(map[string]string)
	for name, value := range params {
		fmt.Fprintf(w, "%s=%s,", name, value)
	}
	fmt.Fprintf(w, "\n")
}

func TestController(t *testing.T) {
	root := NewRouter(false)
	root.Get("/", hello)
	root.Get("/log", hello, Log())
	root.Get("/time", timeHandle)
	root.Get("/time/log", timeHandle, Log())
	root.Get("/name", printNameHandle)
	root.Get("/name/log", printNameHandle, Log())
	root.Get("/params/<string:name>/<int:id>", printParams)
	root.Get("/params/log/<string:action>", printParams, Log())
	root.Get("/params/<string:action>", printParams, Log())
	u := NewController((*UserView)(nil))
	root.AddRouter("/user/<int:id>", []string{"GET"}, u)
	http.ListenAndServe(":8080", &root)
}
