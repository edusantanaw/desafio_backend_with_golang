package usecases

import (
	"errors"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/repository"

	"github.com/google/uuid"
)

func CreateCustomer(data schema.CustomerSchema) (*entities.Customer, error) {
	customerRepository := repository.GetRepository()
	verifyEmail := customerRepository.FindByEmail(data.Email)
	if verifyEmail != nil {
		return nil, errors.New("email ja esta em uso")
	}
	verifyCpfCnpj := customerRepository.FindByCpfCnpj(data.CPF_CNPJ)
	if verifyCpfCnpj != nil {
		return nil, errors.New("cpf/cnpj ja esta em uso")
	}
	customer := &entities.Customer{Name: data.Name, Email: data.Email, Id: uuid.New().String()}
	customerRepository.Create(*customer)
	return customer, nil
}
