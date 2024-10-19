package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
	Pool *pgxpool.Pool
}

var connection *Connection

func Connect() *Connection {
	connStr := os.Getenv("DATABASE_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newConn, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println(err.Error())
		panic("erro ao conectar com o banco de dados")
	}
	if err = newConn.Ping(ctx); err != nil {
		fmt.Println(err.Error())
		panic("erro ao conectar com o banco de dados")
	}
	connection = &Connection{Pool: newConn}
	return connection
}

func GetConnection() *Connection {
	return connection
}
