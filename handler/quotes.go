package handler

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/crypto-quant/exchange-api-gateway/common"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
)

type GetTickerParams struct {
	Pair string `json:"pair" binding:"required"`
}

// curl -X POST http://127.0.0.1:8080/get_ticker -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
func GetTicker(c *gin.Context) {
	var quoteParams GetTickerParams
	if err := c.ShouldBindJSON(&quoteParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ticker, err := api.RestApi.GetTicker(common.ParseTradingPair(quoteParams.Pair))
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

// curl -X POST http://127.0.0.1:8080/get_depth -H 'content-type: application/json' -d '{ "pair": "BTC-USDT", "size": 100 }'
func GetDepth(c *gin.Context) {
	var quoteParams GetDepthParams
	if err := c.ShouldBindJSON(&quoteParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	depth, err := api.RestApi.GetDepth(quoteParams.Size, common.ParseTradingPair(quoteParams.Pair))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": depth})
}

type GetKlinesParams struct {
	Pair      string           `json:"pair" binding:"required"`
	Period    goex.KlinePeriod `json:"period" binding:"required"`
	Size      int              `json:"size" binding:"required"`
	StartTime int              `json:"start_time" binding:"omitempty"`
	EndTime   int              `json:"end_time" binding:"omitempty"`
}

// https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#klinecandlestick-data
// date -j -f %Y%m%d 20200101  +%s
// curl -X POST http://127.0.0.1:8080/get_klines -H 'content-type: application/json' -d '{ "pair": "BTC-USDT", "size": 2, "period": 13, "start_time": 1577887978 }'
func GetKLines(c *gin.Context) {
	var quoteParams GetKlinesParams
	if err := c.ShouldBindJSON(&quoteParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	optionalParameter := goex.OptionalParameter{}
	if quoteParams.StartTime > 0 {
		optionalParameter.Optional("startTime", quoteParams.StartTime)
	}
	if quoteParams.EndTime > 0 {
		optionalParameter.Optional("endTime", quoteParams.EndTime)
	}
	bars, err := api.RestApi.GetKlineRecords(common.ParseTradingPair(quoteParams.Pair), quoteParams.Period, quoteParams.Size, optionalParameter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": bars})
}
