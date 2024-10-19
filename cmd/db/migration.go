package db

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

func RunMigrations() {
	files, err := os.ReadDir("migrations")
	if err != nil {
		fmt.Println("erro ao ler a pasta de migrations")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn := GetConnection()
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.Split(file.Name(), ".")[1] == "sql" {
			data, err := os.ReadFile(fmt.Sprintf("migrations/%s", file.Name()))
			if err != nil {
				fmt.Printf("erro ao ler arquivo: %s, error: %s", file.Name(), err.Error())
				continue
			}
			parsed := string(data)
			println(parsed)
			pg, err := conn.Pool.Exec(ctx, parsed)
			fmt.Print(pg.String())
			if err != nil {
				fmt.Print(err.Error())
			}
		}
	}
}
