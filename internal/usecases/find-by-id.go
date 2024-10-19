package usecases

import (
	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/repository"
)

func FindCustomerById(id string) *entities.Customer {
	customerRepository := repository.GetRepository()
	customer := customerRepository.FindById(id)
	return customer
}
