package handler

import (
	"strings"

	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
)

func GetBalance(c *gin.Context) {
	assets := make(map[string]goex.SubAccount)
	account, err := api.RestApi.GetAccount()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for currency, subAccount := range account.SubAccounts {
		if subAccount.Amount > 0 {
			assets[strings.ToLower(currency.Symbol)] = subAccount
		}
	}

	c.JSON(200, gin.H{
		"data": assets,
	})
}
