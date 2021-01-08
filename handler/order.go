package handler

import (
	"math"
	"strconv"
	"time"

	"github.com/crypto-quant/exchange-api-gateway/api"
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
)

// Binding from JSON
type LimitOrderParams struct {
	Side   goex.TradeSide `json:"side" binding:"required"`
	Pair   string         `json:"pair" binding:"required"`
	Price  string         `json:"price" binding:"required"`
	Amount string         `json:"amount" binding:"required"`
}

// curl -X POST http://localhost:8080/limit_order -H 'content-type: application/json' -d '{ "side": 2, "pair": "BTC-USDT", "price": "45000", "amount": "0.01"}'
func LimitOrder(c *gin.Context) {
	var orderParams LimitOrderParams
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

type MarketOrderParams struct {
	Side   goex.TradeSide `json:"side" binding:"required"`
	Pair   string         `json:"pair" binding:"required"`
	Amount string         `json:"amount" binding:"required"`
}

func MarketOrder(c *gin.Context) {
	var orderParams MarketOrderParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if orderParams.Side == goex.BUY {
		order, err := api.RestApi.MarketBuy(orderParams.Amount, "0", goex.NewCurrencyPair3(orderParams.Pair, "-"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": order})
	} else {
		order, err := api.RestApi.MarketSell(orderParams.Amount, "0", goex.NewCurrencyPair3(orderParams.Pair, "-"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": order})
	}
}

type CancelOrderParams struct {
	OrderId string `json:"order_id" binding:"required"`
	Pair    string `json:"pair" binding:"required"`
}

func CancelOrder(c *gin.Context) {
	var orderParams CancelOrderParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := api.RestApi.CancelOrder(orderParams.OrderId, goex.NewCurrencyPair3(orderParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": result})
}

type GetOrderParams struct {
	OrderId string `json:"order_id" binding:"required"`
	Pair    string `json:"pair" binding:"required"`
}

func GetOrder(c *gin.Context) {
	var orderParams GetOrderParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := api.RestApi.GetOneOrder(orderParams.OrderId, goex.NewCurrencyPair3(orderParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if math.IsNaN(order.AvgPrice) {
		order.AvgPrice = 0
	}
	c.JSON(200, gin.H{"data": order})
}

type GetUnfinishedOrdersParams struct {
	Pair string `json:"pair" binding:"required"`
}

func GetUnfinishedOrders(c *gin.Context) {
	var orderParams GetUnfinishedOrdersParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	orders, err := api.RestApi.GetUnfinishOrders(goex.NewCurrencyPair3(orderParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for index, order := range orders {
		if math.IsNaN(order.AvgPrice) {
			orders[index].AvgPrice = 0
		}
	}
	c.JSON(200, gin.H{"data": orders})
}

func CancelAllOrders(c *gin.Context) {
	var orderParams GetUnfinishedOrdersParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	orders, err := api.RestApi.GetUnfinishOrders(goex.NewCurrencyPair3(orderParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := make([]string, 0)
	for _, order := range orders {
		orderId := strconv.Itoa(order.OrderID)
		_, err := api.RestApi.CancelOrder(orderId, order.Currency)
		if err == nil {
			result = append(result, orderId)
			time.Sleep(120 * time.Millisecond)
		}
	}
	c.JSON(200, gin.H{"data": result})
}

func GetOrderHistory(c *gin.Context) {
	var orderParams GetUnfinishedOrdersParams
	if err := c.ShouldBindJSON(&orderParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	orders, err := api.RestApi.GetOrderHistorys(goex.NewCurrencyPair3(orderParams.Pair, "-"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for index, order := range orders {
		if math.IsNaN(order.AvgPrice) {
			orders[index].AvgPrice = 0
		}
	}

	c.JSON(200, gin.H{"data": orders})
}
