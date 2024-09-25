package config

import (
	"github.com/Agrimaldo/blockchain-in-go/models"
	"github.com/uptrace/bun"
)

var (
	EnvVariables *models.Enviroment
	BunDB        *bun.DB
)
