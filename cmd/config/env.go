package config

import (
	"fmt"
	"os"
	"strings"
)

func Env() {
	env, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Env não encontrado!")
		return
	}
	fmt.Println(env)
	envParsed := string(env)
	for _, enviroment := range strings.Split(envParsed, "\n") {
		splited := strings.Split(enviroment, "=")
		os.Setenv(splited[0], splited[1])
	}
}
