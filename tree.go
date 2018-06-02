package hrouter

import (
	"net/http"
	"sync"
)

type (
	routes struct {
		// method and mounted path-handler
		routes map[string]*node
		mu     sync.RWMutex
	}
	node struct {
		// path and handler
		paths map[string]http.HandlerFunc
		mu    sync.RWMutex
	}
)

func newNode() *node {
	return &node{
		paths: make(map[string]http.HandlerFunc),
	}
}

func (n *node) get(path string) http.HandlerFunc {
	n.mu.RLock()
	h := http.NotFound
	if hh, ok := n.paths[path]; ok {
		h = hh
	}
	n.mu.RUnlock()
	return h
}

func (n *node) add(path string, h http.HandlerFunc) {
	n.mu.Lock()
	if _, ok := n.paths[path]; ok {
		panic("err: path " + path + " already registered")
	}
	n.paths[path] = h
	n.mu.Unlock()
}

func newRoutes() *routes {
	return &routes{
		routes: make(map[string]*node),
	}
}

func (r *routes) add(m string, p string, h http.HandlerFunc) {
	r.mu.Lock()
	if p == "" {
		p = "/"
	}
	if r.routes[m] == nil {
		r.routes[m] = newNode()
	}
	r.routes[m].add(p, h)
	r.mu.Unlock()
}

func (r *routes) get(req *http.Request) http.HandlerFunc {
	h := http.NotFound
	r.mu.RLock()
	if n, ok := r.routes[req.Method]; ok {
		h = n.get(req.URL.Path)
	}
	r.mu.RUnlock()
	return h
}
