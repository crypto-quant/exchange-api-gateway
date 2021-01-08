package handler

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
)

type GetTickerParams struct {
	Pair string `json:"pair" binding:"required"`
}

// curl -X GET http://127.0.0.1:8080/get_ticker -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
func GetTicker(c *gin.Context) {
	var quoteParams GetTickerParams
	if err := c.ShouldBindJSON(&quoteParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ticker, err := api.RestApi.GetTicker(goex.NewCurrencyPair3(quoteParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": ticker})
}

type GetDepthParams struct {
	Pair string `json:"pair" binding:"required"`
	Size int    `json:"size" binding:"required"`
}

// curl -X GET http://127.0.0.1:8080/get_depth -H 'content-type: application/json' -d '{ "pair": "BTC-USDT", "size": 100 }'
func GetDepth(c *gin.Context) {
	var quoteParams GetDepthParams
	if err := c.ShouldBindJSON(&quoteParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	depth, err := api.RestApi.GetDepth(quoteParams.Size, goex.NewCurrencyPair3(quoteParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": depth})
}
