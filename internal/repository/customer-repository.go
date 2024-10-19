package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/edusantanaw/desafio_backend_with_golang/cmd/db"
	"github.com/edusantanaw/desafio_backend_with_golang/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func GetCustomerRepository() *Repository {
	conn := db.GetConnection()
	repository := &Repository{conn: conn.Pool}
	return repository
}

func (r *Repository) Create(data *entities.Customer) (*entities.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.conn.Exec(ctx, "INSERT INTO customers(id, name, email, cpf_cnpj) VALUES ($1, $2, $3, $4);", data.Id, data.Name, data.Email, data.CPF_CNPJ)
	if err != nil {
		return nil, fmt.Errorf("error ao criar usario!")
	}
	return data, nil
}

func (r *Repository) FindByEmail(email string) *entities.Customer {

	return nil
}

func (r *Repository) FindByCpfCnpj(cpfCnpj string) *entities.Customer {

	return nil
}
