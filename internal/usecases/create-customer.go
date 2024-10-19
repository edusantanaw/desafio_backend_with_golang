package usecases

import (
	"errors"
	"os"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/controllers/schema"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/repository"
	"github.com/edusantanaw/desafio_backend_with_golang/pkg/utils"

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
	encrypterPass, err := utils.Encrypt(data.Pass, os.Getenv("SECRET"))
	if err != nil {
		return nil, errors.New("encrypter failed")
	}
	customer := &entities.Customer{Name: data.Name, Email: data.Email, Id: uuid.New().String(), CPF_CNPJ: data.CPF_CNPJ}
	customer.SetPassword(encrypterPass)
	customerRepository.Create(*customer)
	return customer, nil
}
