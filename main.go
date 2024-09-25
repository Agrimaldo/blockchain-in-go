package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Agrimaldo/blockchain-in-go/api"
	"github.com/Agrimaldo/blockchain-in-go/config"
	"github.com/Agrimaldo/blockchain-in-go/models"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	config.EnvVariables = new(models.Enviroment)
	config.EnvVariables.Load()
	StartDBBun()
	api.Routes()
	//shell := os.Getenv("LOCAL_ENV")
	//fmt.Println("first Go Lang Program!")
	//fmt.Printf("TEST: %s\n", config.EnvVariables.DatabaseHost)
}

func StartDBBun() {

	pgconn := pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", config.EnvVariables.DatabaseHost, config.EnvVariables.DatabasePort)),
		//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(config.EnvVariables.DatabaseUser),
		pgdriver.WithPassword(config.EnvVariables.DatabasePass),
		pgdriver.WithDatabase(config.EnvVariables.DatabaseName),
		pgdriver.WithTimeout(5*time.Minute),
		pgdriver.WithDialTimeout(5*time.Minute),
		pgdriver.WithReadTimeout(5*time.Minute),
		pgdriver.WithWriteTimeout(5*time.Minute),
	)

	sqldb := sql.OpenDB(pgconn)
	config.BunDB = bun.NewDB(sqldb, pgdialect.New())
}
