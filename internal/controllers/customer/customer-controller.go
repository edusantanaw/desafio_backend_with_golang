package customer

import (
	"github.com/edusantanaw/desafio_backend_with_golang/adapter"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/usecases"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

func Create(ctx *adapter.AdapterContext[schema.CustomerSchema]) utils.HttpResponse {
	customer, err := usecases.CreateCustomer(ctx.Body)
	if err != nil {
		return utils.HttpResponse{Code: 400, Body: err.Error()}
	}
	return utils.HttpResponse{Code: 200, Body: customer}
}

func FindAll(ctc *adapter.GetAdapterContext) utils.HttpResponse {
	customers := usecases.FindAllCustomer()
	return utils.HttpResponse{Code: 200, Body: customers}
}
