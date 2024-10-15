package main

import (
	"fmt"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/routes"
)

func main() {
	config.Env()
	PORT := os.Getenv("PORT")
	PORT = fmt.Sprintf(":%s", PORT)
	router := routes.MainRouter()
	config.Server(PORT, router)
}
