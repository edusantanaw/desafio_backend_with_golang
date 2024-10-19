package routes

import (
	"net/http"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/customer"
)

func CustomerRouter(router *http.ServeMux) {
	router.HandleFunc("POST /api/customer", customer.Create)
	router.HandleFunc("/api/customer", customer.FindAll)
	router.HandleFunc("/api/customer/{id}", customer.FinById)
}
