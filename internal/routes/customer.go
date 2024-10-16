package routes

import (
	"github.com/edusantanaw/desafio_backend_with_golang/adapter"
	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/customer"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
)

func CustomerRouter(router *config.Routers) {
	router.POST("/api/customer", adapter.AdapterWithBody(customer.Create, schema.CustomerSchema{}, "/api/customer"))
}
