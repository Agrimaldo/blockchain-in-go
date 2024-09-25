package api

import (
	"net/http"

	"github.com/Agrimaldo/blockchain-in-go/services"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/genesis-block", getGenesisBlock)
	router.Run("localhost:8080")
}

func getGenesisBlock(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, services.GetGenesisBlock())
}
