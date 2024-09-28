package services

import (
	"context"
	"fmt"

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
		genesisBlock := new(models.Block)
		_, insertGenesisBlockErr := config.BunDB.NewInsert().Model(&genesisBlock).Exec(ctx)
		if insertGenesisBlockErr != nil {
			panic(fmt.Sprint("Error on Insert Genesis Block: ", insertGenesisBlockErr))
		}
	}

}

func GetGenesisBlock() models.Block {
	return models.Block{Id: 0, Height: 0}
}
