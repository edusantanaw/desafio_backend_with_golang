package main

import (
	"fmt"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/adapter"
	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

func main() {
	config.Env()
	PORT := os.Getenv("PORT")
	PORT = fmt.Sprintf(":%s", PORT)
	router := config.Router()
	router.Get("/test/:id", adapter.AdapterWithQuery(func(ctx *adapter.AdapterContext) utils.HttpResponse {
		fmt.Println(ctx.Params["id"])
		return utils.HttpResponse{Code: 200, Body: "ola"}
	}, "/test/:id"))
	config.Server(PORT, router)
}
