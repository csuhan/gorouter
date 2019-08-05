package gorouter

import "net/http"

type Router struct {
	ServeMux *http.ServeMux
	route    map[string]string
}

// New a gorouter
func New() *Router {
	return &Router{ServeMux: http.NewServeMux()}
}

func (router *Router) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("GET", path, handler)
}

func (router *Router) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("POST", path, handler)
}

func (router *Router) Head(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("HEAD", path, handler)
}

func (router *Router) Delete(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("DELETE", path, handler)
}
func (router *Router) Handle(method string, path string, handler http.Handler) {
	router.ServeMux.Handle(path, handler)
}

func (router *Router) HandleFunc(method string, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	if router.route == nil {
		router.route = make(map[string]string)
	}
	router.route[path] = method
	router.ServeMux.HandleFunc(path, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := router.route[r.URL.Path]; ok {
		if router.route[r.URL.Path] == r.Method {
			router.ServeMux.ServeHTTP(w, r)
		}
	}
}
