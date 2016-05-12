package satumotu

import (
	"github.com/foodszhang/trie_router"
	"github.com/gorilla/context"
	"net/http"
	"sync"
)

type Router struct {
	*router.RouteNode
	hosts bool
	mu    sync.RWMutex
}

func NewRouter(hosts bool) Router {
	return Router{router.CreateRouteNode(), hosts, sync.RWMutex{}}
}

func (r *Router) AddRouter(url string, methods []string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, methods, handle, adapters...)
}
func (r *Router) Get(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"GET"}, handle, adapters...)
}
func (r *Router) Post(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"POST"}, handle, adapters...)
}
func (r *Router) Put(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"PUT"}, handle, adapters...)
}
func (r *Router) Delete(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"DELETE"}, handle, adapters...)
}
func (r *Router) Option(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"OPTION"}, handle, adapters...)
}
func (r *Router) Head(url string, handle http.HandlerFunc, adapters ...router.Adapter) error {
	return r.Insert(url, []string{"HEAD"}, handle, adapters...)
}
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	if request.RequestURI == "*" {
		w.Header().Set("Connection", "close")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h := r.Handler(request)
	h.ServeHTTP(w, request)
}
func (r *Router) Handler(req *http.Request) (h http.Handler) {
	if req.Method != "CONNECT" {
		if p := cleanPath(req.URL.Path); p != req.URL.Path {
			return http.RedirectHandler(p, http.StatusMovedPermanently)
		}
	}
	h, params := r.handler(req.Host, req.URL.Path, req.Method)
	param_dict := make(map[string]string)
	for _, p := range params {
		param_dict[p.Name] = p.Value
	}
	context.Set(req, "params", param_dict)
	return
}

func (r *Router) handler(host, path, method string) (h http.Handler, params []router.Param) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if r.hosts {
		_, h, params = r.Match(host+path, method)
	}
	if h == nil {
		_, h, params = r.Match(path, method)
	}
	if h == nil {
		h, params = http.NotFoundHandler(), []router.Param{}
	}
	return
}
