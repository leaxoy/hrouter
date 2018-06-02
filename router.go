package hrouter

import (
	"net/http"
)

type (
	Middleware func(next http.HandlerFunc) http.HandlerFunc
	Router struct {
		middlewares []Middleware
		route       *routes
	}
	G struct {
		r           *Router
		middlewares []Middleware
		prefix      string
	}
)

func New() *Router {
	return &Router{
		route: newRoutes(),
	}
}

func (r *Router) Use(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *Router) Get(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodGet, path, h, middlewares...)
}

func (r *Router) Post(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodPost, path, h, middlewares...)
}

func (r *Router) Put(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodPut, path, h, middlewares...)
}

func (r *Router) Delete(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodDelete, path, h, middlewares...)
}

func (r *Router) Connect(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodConnect, path, h, middlewares...)
}

func (r *Router) Head(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodHead, path, h, middlewares...)
}

func (r *Router) Patch(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodPatch, path, h, middlewares...)
}

func (r *Router) Options(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodOptions, path, h, middlewares...)
}

func (r *Router) Trace(path string, h http.HandlerFunc, middlewares ...Middleware) {
	r.Handle(http.MethodTrace, path, h, middlewares...)
}

func (r *Router) Handle(method string, path string, h http.HandlerFunc, middlewares ...Middleware) {
	middlewares = append(r.middlewares, middlewares...)
	for _, m := range middlewares {
		h = m(h)
	}
	r.route.add(method, path, h)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h := r.route.get(request)
	for _, m := range r.middlewares {
		h = m(h)
	}
	h(writer, request)
}

func (r *Router) G(path string, middlewares ...Middleware) *G {
	if path == "" {
		panic("path can'routes be empty")
	}
	return &G{
		r:           r,
		prefix:      path,
		middlewares: middlewares,
	}
}

func (g *G) Use(middlewares ...Middleware) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *G) Get(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodGet, path, h, middlewares...)
}

func (g *G) Post(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodPost, path, h, middlewares...)
}

func (g *G) Put(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodPut, path, h, middlewares...)
}

func (g *G) Delete(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodDelete, path, h, middlewares...)
}

func (g *G) Connect(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodConnect, path, h, middlewares...)
}

func (g *G) Head(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodHead, path, h, middlewares...)
}

func (g *G) Patch(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodPatch, path, h, middlewares...)
}

func (g *G) Options(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodOptions, path, h, middlewares...)
}

func (g *G) Trace(path string, h http.HandlerFunc, middlewares ...Middleware) {
	g.Handle(http.MethodTrace, path, h, middlewares...)
}

func (g *G) Handle(method string, path string, h http.HandlerFunc, middlewares ...Middleware) {
	middlewares = append(g.middlewares, middlewares...)
	g.r.Handle(method, g.prefix+path, h, middlewares...)
}
