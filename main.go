package main

import (
	"fmt"

	"github.com/Agrimaldo/blockchain-in-go/config"
	"github.com/Agrimaldo/blockchain-in-go/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	config.EnvVariables = new(models.Enviroment)

	config.EnvVariables.Load()

	//shell := os.Getenv("LOCAL_ENV")
	//fmt.Println("first Go Lang Program!")
	fmt.Printf("TEST: %s\n", config.EnvVariables.DatabaseHost)
}
