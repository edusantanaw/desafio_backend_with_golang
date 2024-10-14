package config

import (
	"fmt"
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	method := r.Method
	fmt.Println(method, url)
	handler, err := getRouterHanlder(method, url)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	handler(w, r)
}

func Server(port string) {
	handler := http.HandlerFunc(httpHandler)
	fmt.Printf("Server running %s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
