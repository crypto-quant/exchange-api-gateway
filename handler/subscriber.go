package handler

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/crypto-quant/exchange-api-gateway/common"
	"github.com/gin-gonic/gin"
)

type SubscribeParams struct {
	Pair string `json:"pair" binding:"required"`
}

// curl -X POST http://127.0.0.1:8080/subscribe_ticker -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
func SubscribeTicker(c *gin.Context) {
	var params SubscribeParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := api.WsApi.SubscribeTicker(common.ParseTradingPair(params.Pair))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": "success"})
}

// curl -X POST http://127.0.0.1:8080/subscribe_depth -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
func SubscribeDepth(c *gin.Context) {
	var params SubscribeParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := api.WsApi.SubscribeDepth(common.ParseTradingPair(params.Pair))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": "success"})
}

// curl -X POST http://127.0.0.1:8080/subscribe_account -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
// func SubscribeAccount(c *gin.Context) {
// 	var params SubscribeParams
// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := api.WsApi.SubscribeAccount(common.ParseTradingPair(params.Pair))
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"data": "success"})
// }

// func SubscribeOrder(c *gin.Context) {
// 	var params SubscribeParams
// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := api.WsApi.SubscribeOrder(common.ParseTradingPair(params.Pair))
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"data": "success"})
// }
