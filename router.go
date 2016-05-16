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

type Adapter func(http.Handler) http.Handler

func (a Adapter) Adapt(h http.Handler) http.Handler {
	return a(h)
}

func (r *Router) AddRouter(url string, methods []string, handle http.Handler, adapters ...router.Adapter) error {
	return r.Insert(url, methods, handle, adapters...)
}
func (r *Router) Get(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"GET"}, handle, adps...)
}
func (r *Router) Post(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"POST"}, handle, adps...)
}
func (r *Router) Put(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"PUT"}, handle, adps...)
}
func (r *Router) Delete(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"DELETE"}, handle, adps...)
}
func (r *Router) Option(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"OPTION"}, handle, adps...)
}
func (r *Router) Head(url string, handle http.HandlerFunc, adapters ...Adapter) error {
	adps := make([]router.Adapter, len(adapters))
	for i, v := range adapters {
		adps[i] = v
	}
	return r.Insert(url, []string{"HEAD"}, handle, adps...)
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

// Router 作为分发请求的手段，根据request分发给不同的路由，在路由执行ServerHTTP时，在外层包裹中间层，里面执行绑定的HandleFunc
// 通过context传输经过中间层所设置的数据
// RouterGroup 一个组有共同的前缀，共同的中间层，每一个组可以直接拿来当做外部接口，也可以附在一个父组上作为子节点，一个组包括一系列
// Router, 每个Router也可以作为一个RouterGroup附在别的RouterGroup上。每次附在别的group时，router会保存一个绝对路径
// 分发时直接通过绝对路径进行匹配， 最终绑定工作都交给router来做
// server过程 ServerHTTP->RouterGroup->匹配->Router->中间件->HandleFunc
// 初始化 HandlerFunc绑定到Router->Router绑定到RouterGroup->RouterGroup记录绝对路径->RouterGroup绑定上层RouterGroup->上层RouterGroup记录绝对路径....顶层
// 所以RouterGroup 里面两个属性 []groups， []routers，不含正则的直接通过map[string]router记录，然后直接用绝对路径匹配,
// 含正则的通过map[string]groups，对前缀进行匹配，每次相当于缩短了一部分路径
//1：遍历所有router得到结果，2:分层遍历,根据前缀分层遍历,
// 可以用trie树实现,我这里就懒得用了
// 最长公共前缀进行匹配， 要求group的前缀必须完全匹配 以'/'为结尾
