// just for test

package main

import (
	"fmt"

	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_depth"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_order"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_ticker"
	zmq "github.com/pebbe/zmq4"
	"google.golang.org/protobuf/proto"
)

func main() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5563")
	subscriber.SetSubscribe("ticker")
	subscriber.SetSubscribe("depth")
	subscriber.SetSubscribe("trade")
	subscriber.SetSubscribe("order")

	for {
		//  Read envelope with address
		channel, _ := subscriber.Recv(0)
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		if channel == "ticker" {
			ticker := &pb_ticker.Ticker{}
			if err := proto.Unmarshal([]byte(contents), ticker); err != nil {
				fmt.Printf("Failed to parse ticker: %v\n", err)
			} else {
				fmt.Println(ticker)
			}
		} else if channel == "depth" {
			depth := &pb_depth.Depth{}
			if err := proto.Unmarshal([]byte(contents), depth); err != nil {
				fmt.Printf("Failed to parse depth: %v\n", err)
			} else {
				fmt.Println(depth)
			}
		} else if channel == "order" {
			order := &pb_order.Order{}
			if err := proto.Unmarshal([]byte(contents), order); err != nil {
				fmt.Printf("Failed to parse order: %v\n", err)
			} else {
				fmt.Printf("status: %v\n", order.Status)
				fmt.Printf("%+v\n", order)
			}
		} else {
			fmt.Println(channel)
		}

	}
}
