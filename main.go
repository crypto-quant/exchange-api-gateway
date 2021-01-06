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
	r.Run("127.0.0.1:8080")
}
