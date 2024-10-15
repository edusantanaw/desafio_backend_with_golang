package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/cmd/config"
)

func main() {
	config.Env()
	PORT := os.Getenv("PORT")
	PORT = fmt.Sprintf(":%s", PORT)
	router := config.Router()
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
		w.WriteHeader(200)
	})
	router.Get("/test/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, with hello!"))
		w.WriteHeader(200)
	})
	router.Get("/test/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, with id!"))
		w.WriteHeader(200)
	})
	config.Server(PORT, router)
}
