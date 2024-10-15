package routes

import "github.com/edusantanaw/desafio_backend_with_golang/cmd/config"

func MainRouter() *config.Routers {
	router := config.Router()
	CustomerRouter(router)
	return router
}
