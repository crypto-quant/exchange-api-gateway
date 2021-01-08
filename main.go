package main

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/crypto-quant/exchange-api-gateway/config"
	"github.com/crypto-quant/exchange-api-gateway/handler"
	"github.com/crypto-quant/exchange-api-gateway/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

var (
	platformConfig *config.PlatformConfig

	log *logrus.Logger
)

func init() {
	log = logger.InitLogger()

	platformConfig = config.ParseConfigFile("config.yml")
	api.InitApi(platformConfig)

}

func main() {
	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.GET("/balance", handler.GetBalance)
	r.POST("/limit_order", handler.LimitOrder)
	r.POST("/market_order", handler.MarketOrder)
	r.POST("/cancel_order", handler.CancelOrder)
	r.POST("/cancel_all_orders", handler.CancelAllOrders)
	r.POST("/get_order", handler.GetOrder)
	r.POST("/get_unfinished_order", handler.GetUnfinishedOrders)
	r.GET("/get_trading_pairs", handler.GetTradingPairs)
	r.GET("/get_ticker", handler.GetTicker)
	r.GET("/get_depth", handler.GetDepth)

	r.Run("127.0.0.1:8080")
}
