package handler

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
)

// Binding from JSON
type OrderParams struct {
	Side   goex.TradeSide `json:"side" binding:"required"`
	Pair   string         `json:"pair" binding:"required"`
	Price  string         `json:"price" binding:"required"`
	Amount string         `json:"amount" binding:"required"`
}

// curl -X POST http://localhost:8080/limit_order -H 'content-type: application/json' -d '{ "side": 2, "pair": "BTC-USDT", "price": "45000", "amount": "0.01"}'
func LimitOrder(c *gin.Context) {
	var orderParams OrderParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if orderParams.Side == goex.BUY {
		order, err := api.RestApi.LimitBuy(orderParams.Amount, orderParams.Price, goex.NewCurrencyPair3(orderParams.Pair, "-"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": order})
	} else {
		order, err := api.RestApi.LimitSell(orderParams.Amount, orderParams.Price, goex.NewCurrencyPair3(orderParams.Pair, "-"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": order})
	}
}
