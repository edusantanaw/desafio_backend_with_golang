package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const dotEnvFilePath = ".env"

func Env() {
	file, err := os.Open(dotEnvFilePath)
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar %s: %v", dotEnvFilePath, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Linha mal formatada: %s\n", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("Erro ao definir variável %s: %v\n", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Erro ao ler %s: %v", dotEnvFilePath, err))
	}
}
