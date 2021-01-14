package main

import (
	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/crypto-quant/exchange-api-gateway/config"
	"github.com/crypto-quant/exchange-api-gateway/handler"
	"github.com/crypto-quant/exchange-api-gateway/logger"
	"github.com/crypto-quant/exchange-api-gateway/zmq"
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

	go zmq.InitPublisher("tcp://*:5563")
}

func main() {
	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.POST("/balance", handler.GetBalance)
	r.POST("/limit_order", handler.LimitOrder)
	r.POST("/market_order", handler.MarketOrder)
	r.POST("/cancel_order", handler.CancelOrder)
	r.POST("/cancel_all_orders", handler.CancelAllOrders)
	r.POST("/get_order", handler.GetOrder)
	r.POST("/get_order_history", handler.GetOrderHistory)
	r.POST("/get_unfinished_orders", handler.GetUnfinishedOrders)
	r.POST("/get_trading_pairs", handler.GetTradingPairs)
	r.POST("/get_ticker", handler.GetTicker)
	r.POST("/get_depth", handler.GetDepth)
	r.POST("/get_klines", handler.GetKLines)
	r.POST("/subscribe_ticker", handler.SubscribeTicker)
	r.POST("/subscribe_depth", handler.SubscribeDepth)
	r.POST("/subscribe_trade", handler.SubscribeTrade)
	r.POST("/subscribe_order", handler.SubscribeTrade)

	r.Run("127.0.0.1:8080")
}
