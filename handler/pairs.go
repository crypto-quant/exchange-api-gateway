package handler

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex/binance"
)

func GetTradingPairs(c *gin.Context) {
	binanceRestApi, ok := api.RestApi.(*binance.Binance)
	if !ok {
		c.JSON(400, gin.H{
			"error": "binance api not found",
		})
		return
	}
	exchangInfo, err := binanceRestApi.GetExchangeInfo()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": exchangInfo.Symbols,
	})

}
