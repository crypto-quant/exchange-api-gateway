package common

import (
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_common"
)

func AdaptOrderStatus(status string) pb_common.OrderStatus {
	var tradeStatus pb_common.OrderStatus
	switch status {
	case "NEW":
		tradeStatus = pb_common.OrderStatus_ORDER_NEW
	case "FILLED":
		tradeStatus = pb_common.OrderStatus_ORDER_FILLED
	case "PARTIALLY_FILLED":
		tradeStatus = pb_common.OrderStatus_ORDER_PARTIALLY_FILLED
	case "CANCELED":
		tradeStatus = pb_common.OrderStatus_ORDER_CANCELLED
	case "PENDING_CANCEL":
		tradeStatus = pb_common.OrderStatus_ORDER_PENDING_CANCEL
	case "REJECTED":
		tradeStatus = pb_common.OrderStatus_ORDER_REJECTED
	}
	return tradeStatus
}

func AdaptTradeSide(side string) pb_common.Side {
	var tradeSide pb_common.Side
	switch side {
	case "BUY":
		tradeSide = pb_common.Side_BUY
	case "SELL":
		tradeSide = pb_common.Side_SELL
	case "BUY_MARKET":
		tradeSide = pb_common.Side_BUY_MARKET
	case "SELL_MARKET":
		tradeSide = pb_common.Side_SELL_MARKET
	}
	return tradeSide
}
