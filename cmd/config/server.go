package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	method := r.Method
	fmt.Println(method, url)
	if url == "/test" && method == "GET" {
		w.WriteHeader(200)
		w.Write([]byte("Hello, world222222!"))
		return
	}
	fmt.Println(url)
	w.Write([]byte("Hello, world!"))
	w.WriteHeader(400)
}

func Server() {
	PORT := os.Getenv("PORT")
	handler := http.HandlerFunc(httpHandler)
	fmt.Printf("Server running %s\n", PORT)
	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}
}
