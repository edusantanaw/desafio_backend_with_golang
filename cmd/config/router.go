package config

import (
	"fmt"
	"net/http"
)

type Route struct {
	method  string
	route   string
	handler http.HandlerFunc
}

type Routers struct {
	routers []Route
}

var routers = Routers{routers: []Route{}}

func getRouterHanlder(method string, route string) (http.HandlerFunc, error) {
	for _, router := range routers.routers {
		if router.method == method && router.route == route {
			return router.handler, nil
		}
	}
	return nil, fmt.Errorf("route not found: %s %s", method, route)
}

func Router() *Routers {
	return &routers
}

func (r *Routers) Get(path string, handler http.HandlerFunc) {
	router := Route{method: "GET", route: path, handler: handler}
	routers.routers = append(routers.routers, router)
}
