package zmq

import (
	"encoding/json"
	"time"

	"github.com/crypto-quant/exchange-api-gateway/common"
	. "github.com/crypto-quant/exchange-api-gateway/logger"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_depth"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_ticker"
	"github.com/nntaoli-project/goex"
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
	Publisher.Send(channel, zmq.SNDMORE)

	jsonData, err := json.Marshal(data)
	if err != nil {
		Logger.Error(err)
		return
	}
	Publisher.Send(string(jsonData), 0)
}

func PubPBTicker(ticker *goex.Ticker) {
	Publisher.Send("ticker", zmq.SNDMORE)
	t := &pb_ticker.Ticker{
		Pair: ticker.Pair.ToSymbol(common.Separater),
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
	Publisher.Send(string(pbData), 0)
}

func PubPBDepth(depth *goex.Depth) {
	Publisher.Send("depth", zmq.SNDMORE)
	t := &pb_depth.Depth{
		Pair: depth.Pair.ToSymbol(common.Separater),
		Date: int64(time.Nanosecond) * depth.UTime.UnixNano() / int64(time.Millisecond),
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
	Publisher.Send(string(pbData), 0)
}
