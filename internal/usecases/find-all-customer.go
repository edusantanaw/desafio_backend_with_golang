package usecases

import (
	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/repository"
)

func FindAllCustomer() []entities.Customer {
	customerRepository := repository.GetRepositoryInMemory()
	customers := customerRepository.FindAll()
	return customers
}
