package config

import "net/http"

var router = http.NewServeMux()

func Router() *http.ServeMux {
	return router
}
