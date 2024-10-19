package customer

import (
	"fmt"
	"net/http"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/usecases"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var customerBody schema.CustomerSchema
	if err := utils.ParseJSON(req, &customerBody); err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}
	customer, err := usecases.CreateCustomer(customerBody)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}
	if err = utils.WriteJson(res, 201, customer); err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}
}

func FindAll(res http.ResponseWriter, req *http.Request) {
	customers := usecases.FindAllCustomer()
	if err := utils.WriteJson(res, 200, customers); err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
	}
}

func FinById(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	customer := usecases.FindCustomerById(id)
	if customer == nil {
		utils.WriteError(res, http.StatusNotFound, fmt.Errorf("client não encontrado"))
		return
	}
	if err := utils.WriteJson(res, 200, customer); err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
	}
}
