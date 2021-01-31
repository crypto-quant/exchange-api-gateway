package zmq

import (
	"encoding/json"

	"github.com/crypto-quant/exchange-api-gateway/common"
	. "github.com/crypto-quant/exchange-api-gateway/logger"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_account"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_depth"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_order"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_ticker"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/binance"
	zmq "github.com/pebbe/zmq4"
	"google.golang.org/protobuf/proto"
)

var (
	Publisher *zmq.Socket
)

// endpoint: "tcp://*:5563"
func InitPublisher(endpoint string) {
	Publisher, _ = zmq.NewSocket(zmq.PUB)
	Publisher.Bind(endpoint)
}

func PubJson(channel string, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		Logger.Error(err)
		return
	}
	PubData(channel, jsonData)
}

func PubPBTicker(ticker *goex.Ticker) {
	t := &pb_ticker.Ticker{
		Pair: ticker.Pair.ToSymbol(common.Separater),
		Open: ticker.Open,
		Last: ticker.Last,
		Buy:  ticker.Buy,
		Sell: ticker.Sell,
		High: ticker.High,
		Low:  ticker.Low,
		Vol:  ticker.Vol,
		Date: int64(ticker.Date),
	}
	pbData, err := proto.Marshal(t)
	if err != nil {
		Logger.Error(err)
	}
	PubData("ticker", pbData)
}

func PubPBDepth(depth *goex.Depth) {
	t := &pb_depth.Depth{
		Pair: depth.Pair.ToSymbol(common.Separater),
		Date: common.NsToMs(depth.UTime.UnixNano()),
	}
	t.Asks = make([]*pb_depth.Depth_PriceLevel, 0)
	for _, ask := range depth.AskList {
		t.Asks = append(t.Asks, &pb_depth.Depth_PriceLevel{
			Price:  ask.Price,
			Amount: ask.Amount,
		})
	}
	for _, bid := range depth.BidList {
		t.Bids = append(t.Bids, &pb_depth.Depth_PriceLevel{
			Price:  bid.Price,
			Amount: bid.Amount,
		})
	}
	pbData, err := proto.Marshal(t)
	if err != nil {
		Logger.Error(err)
	}
	PubData("depth", pbData)
}

func PubPBOrderUpdate(order *binance.OrderUpdate) {
	t := &pb_order.Order{
		Pair:         common.AdaptSymbolToTradingPair(order.Data.Symbol),
		Price:        order.Data.Price,
		Amount:       order.Data.Quantity,
		DealAmount:   order.Data.CumulativeFilledQuantity,
		Side:         common.AdaptTradeSide(order.Data.Side),
		Status:       common.AdaptOrderStatus(order.Data.OrderStatus),
		Cid:          order.Data.ClientOrderID,
		OrderId:      order.Data.OrderID,
		TimeInForce:  order.Data.TimeInForce,
		OrderType:    order.Data.OrderType,
		OrderTime:    common.NsToMs(order.Data.OrderCreationTime),
		FinishedTime: common.NsToMs(order.Data.TransactionTime),
		TradeId:      order.Data.TradeID,
	}
	if t.Cid == "" {
		t.Cid = order.Data.CancelledClientOrderID
	}
	pbData, err := proto.Marshal(t)
	if err != nil {
		Logger.Error(err)
	}
	PubData("order", pbData)
}

func PubPBAccountPosition(account *binance.AccountPosition) {
	t := &pb_account.Account{
		Date:      common.NsToMs(account.Data.EventTime),
		EventType: account.Data.EventType,
	}

	t.Assets = make([]*pb_account.Account_Asset, 0)
	for _, asset := range account.Data.Currencies {
		t.Assets = append(t.Assets, &pb_account.Account_Asset{
			Symbol:    asset.Asset,
			Available: asset.Available,
			Locked:    asset.Locked,
		})
	}
	pbData, err := proto.Marshal(t)
	if err != nil {
		Logger.Error(err)
	}
	PubData("account", pbData)
}

func PubData(channel string, bytecode []byte) {
	Publisher.Send(channel, zmq.SNDMORE)
	Publisher.Send(string(bytecode), 0)
}
