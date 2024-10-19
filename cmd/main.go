package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/routes"
)

func main() {
	config.Env()
	PORT := os.Getenv("PORT")
	PORT = fmt.Sprintf(":%s", PORT)
	routes.MainRouter()
	server := http.Server{
		Addr:    PORT,
		Handler: config.Router(),
	}
	fmt.Printf("Server running %s\n", PORT)
	server.ListenAndServe()
}
