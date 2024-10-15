package main

import (
	"fmt"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/adapter"
	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

type Query struct {
	Take int
}

func main() {
	config.Env()
	PORT := os.Getenv("PORT")
	PORT = fmt.Sprintf(":%s", PORT)
	router := config.Router()
	router.Get("/test/:id", adapter.AdapterWithQuery(func(data map[string]string) utils.HttpResponse {
		take := data["take"]
		fmt.Println(take)
		return utils.HttpResponse{Code: 200, Body: "ola"}
	}, "/test/:id"))
	config.Server(PORT, router)
}
