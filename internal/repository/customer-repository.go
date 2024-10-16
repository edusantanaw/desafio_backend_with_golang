package repository

import (
	"errors"
	"slices"

	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
)

type CustomerRepository struct {
	items []entities.Customer
}

var repository = &CustomerRepository{items: make([]entities.Customer, 0)}

func GetRepository() *CustomerRepository {
	return repository
}

type IFilter func(v string, e string) bool

func (r *CustomerRepository) FindAll() []entities.Customer {
	return r.items
}

func (r *CustomerRepository) FindById(id string) *entities.Customer {
	for _, customer := range r.items {
		if customer.Id == id {
			return &customer
		}
	}
	return nil
}

func (r *CustomerRepository) FindByEmail(email string) *entities.Customer {
	for _, customer := range r.items {
		if customer.Email == email {
			return &customer
		}
	}
	return nil
}

func (r *CustomerRepository) FindByCpfCnpj(cpfCnpj string) *entities.Customer {
	for _, customer := range r.items {
		if customer.CPF_CNPJ == cpfCnpj {
			return &customer
		}
	}
	return nil
}

func (r *CustomerRepository) Create(customer entities.Customer) *entities.Customer {
	r.items = append(r.items, customer)
	return &customer
}

func (r *CustomerRepository) Update(customer entities.Customer) (*entities.Customer, error) {
	exists := slices.Contains(r.items, customer)
	if !exists {
		return nil, errors.New("not found")
	}
	r.updateList(customer)
	return &customer, nil
}

func (r *CustomerRepository) updateList(customer entities.Customer) {
	updatedList := make([]entities.Customer, 0)
	for _, c := range r.items {
		if c.Id == customer.Id {
			updatedList = append(updatedList, customer)
			continue
		}
		updatedList = append(updatedList, c)
	}
	r.items = updatedList
}
