package zmq

import (
	"encoding/json"

	. "github.com/crypto-quant/exchange-api-gateway/logger"
	zmq "github.com/pebbe/zmq4"
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

// TODO
// func PubPB(channel string) {

// }
