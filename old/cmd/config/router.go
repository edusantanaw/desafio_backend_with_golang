package config

import (
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	method  string
	route   string
	handler http.HandlerFunc
}

type Routers struct {
	routers []Route
}

func (r *Routers) getRouterHandler(method string, route string) (http.HandlerFunc, error) {
	for _, router := range r.routers {
		if router.method != method {
			continue
		}
		paths := strings.Split(router.route, "/")
		routePaths := strings.Split(route, "/")
		if len(paths) != len(routePaths) {
			continue
		}
		diff := false
		for i := range paths {
			if strings.HasPrefix(paths[i], ":") {
				continue
			}
			if paths[i] != routePaths[i] {
				diff = true
			}
		}
		if !diff {
			return router.handler, nil
		}
	}
	return nil, fmt.Errorf("route not found: %s %s", method, route)
}

func Router() *Routers {
	var routers = Routers{routers: []Route{}}
	return &routers
}

func (r *Routers) register(method string, path string, handler http.HandlerFunc) {
	router := Route{method: method, route: path, handler: handler}
	r.routers = append(r.routers, router)
}

func (r *Routers) Get(path string, handler http.HandlerFunc) {
	r.register("GET", path, handler)
}

func (r *Routers) POST(path string, handler http.HandlerFunc) {
	r.register("POST", path, handler)
}

func (r *Routers) PUT(path string, handler http.HandlerFunc) {
	r.register("PUT", path, handler)
}

func (r *Routers) DELETE(path string, handler http.HandlerFunc) {
	r.register("DELETE", path, handler)
}

func (r *Routers) PATCH(path string, handler http.HandlerFunc) {
	r.register("PATCH", path, handler)
}
