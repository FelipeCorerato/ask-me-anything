package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Erro ao carregar .env: %v", err))
	}

	fmt.Printf("DATABASE_PORT: %s\n", os.Getenv("DATABASE_PORT"))
	fmt.Printf("DATABASE_NAME: %s\n", os.Getenv("DATABASE_NAME"))
	fmt.Printf("DATABASE_USER: %s\n", os.Getenv("DATABASE_USER"))
	fmt.Printf("DATABASE_PASSWORD: %s\n", os.Getenv("DATABASE_PASSWORD"))
	fmt.Printf("DATABASE_HOST: %s\n", os.Getenv("DATABASE_HOST"))

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Erro ao executar comando tern migrate: %v\nSaída: %s", err, string(output)))
	}

	fmt.Println(string(output))
}
