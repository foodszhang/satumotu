package satumotu

import (
	"net/http"
	"reflect"
	"strings"
)

type Controller struct {
	methods map[string]bool
}

/*
func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	//do something with controller
}
*/

func (c Controller) Bind(url string, methods []string) http.HandlerFunc {
	v := reflect.ValueOf(c)
	for _, method := range methods {
		m := v.MethodByName(strings.ToLower(method))
		if m == reflect.Zero(m.Type()) {
			panic("dont't have this method")
		}
		c.methods[strings.ToLower(method)] = true
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if c.methods[strings.ToLower(r.Method)] {
				m := v.MethodByName(strings.ToLower(r.Method))
				m.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})

			}

		})

}
