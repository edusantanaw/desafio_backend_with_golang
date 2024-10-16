package customer

import (
	"github.com/edusantanaw/desafio_backend_with_golang/adapter"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"
)

func Create(ctx *adapter.AdapterContext) utils.HttpResponse {
	return utils.HttpResponse{Code: 200, Body: ctx.Body}
}
