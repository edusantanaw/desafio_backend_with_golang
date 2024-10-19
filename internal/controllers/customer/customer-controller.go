package customer

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/usecases"
)

func Create(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		http.Error(res, "Invalid or empty request body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	var customerBody schema.CustomerSchema
	if err := json.Unmarshal(body, &customerBody); err != nil {
		http.Error(res, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	customer, err := usecases.CreateCustomer(customerBody)
	if err != nil {
		res.Write([]byte(err.Error()))
		res.WriteHeader(400)
		return
	}
	response, err := json.Marshal(customer)
	if err != nil {
		res.Write([]byte(err.Error()))
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(response)
}

func FindAll(res http.ResponseWriter, req *http.Request) {
	customers := usecases.FindAllCustomer()
	body, err := json.Marshal(customers)
	if err != nil {
		res.Write([]byte(err.Error()))
		res.WriteHeader(500)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(body)
}
