package main

import (
	"io/ioutil"
	"log"

	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_example"
	"github.com/crypto-quant/exchange-api-gateway/pb/out/pb_ticker"
	"google.golang.org/protobuf/proto"
)

func ticker_message() {
	t := &pb_ticker.Ticker{
		Last: 1,
		Buy:  2,
		Sell: 3,
		High: 4,
		Low:  5,
		Vol:  6,
		Date: 7,
	}
	_, err := proto.Marshal(t)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
}

func main() {
	p := &pb_example.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb_example.Person_PhoneNumber{
			{Number: "555-4321", Type: pb_example.Person_HOME},
		},
	}
	fname := "test.pb"
	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb_example.Person{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	log.Printf("%+v\n", book)
}
