package services

import (
	"context"

	"github.com/Agrimaldo/blockchain-in-go/config"
	"github.com/Agrimaldo/blockchain-in-go/models"
)

func Initialize() {
	ctx := context.Background()
	blockCount, blockError := config.BunDB.NewSelect().Model((*models.Block)(nil)).Count(ctx)
	if blockError != nil {
		panic(blockError)
	}

	if blockCount < 1 {

	}

}

func GetGenesisBlock() models.Block {
	return models.Block{Id: 0, Height: 0}
}
