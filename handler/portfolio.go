package handler

import (
	"strings"

	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	assets := make(map[string]float64)
	account, err := api.RestApi.GetAccount()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for currency, subAccount := range account.SubAccounts {
		if subAccount.Amount > 0 {
			assets[strings.ToLower(currency.Symbol)] = subAccount.Amount
		}
	}

	c.JSON(200, gin.H{
		"data": assets,
	})
}
