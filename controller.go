package satumotu

import (
	"fmt"
	"net/http"
	"reflect"
)

type Controller struct {
}
type ControllerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)     //method=GET的处理
	Post(w http.ResponseWriter, r *http.Request)    //method=POST的处理
	Delete(w http.ResponseWriter, r *http.Request)  //method=DELETE的处理
	Put(w http.ResponseWriter, r *http.Request)     //method=PUT的处理
	Head(w http.ResponseWriter, r *http.Request)    //method=HEAD的处理
	Patch(w http.ResponseWriter, r *http.Request)   //method=PATCH的处理
	Options(w http.ResponseWriter, r *http.Request) //method=OPTIONS的处理
}

func (c *Controller) GET(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) POST(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) DELETE(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) PUT(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) HEAD(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) PATCH(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

func (c *Controller) OPTIONS(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method Not Allowed", 405)
}

/*
func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	//do something with controller
}
*/
func NewController(v interface{}) ControllerRouter {
	return ControllerRouter{reflect.TypeOf(v)}
}

type ControllerRouter struct {
	ControllerType reflect.Type
}

func (c ControllerRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c.ControllerType)
	vc := reflect.New(c.ControllerType)
	vc = vc.Elem()
	fmt.Println(vc.Type())
	fmt.Println(vc, r.Method)
	m := vc.MethodByName(r.Method)
	fmt.Println(m)
	m.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})

}
