package api

import (
	"log"

	"github.com/crypto-quant/exchange-api-gateway/config"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/builder"
)

var (
	RestApi   goex.API
	WsApi     goex.SpotWsApi
	WalletApi goex.WalletApi
)

func InitApi(config *config.PlatformConfig) {
	// init binance exchange
	httpClientConfig := &builder.HttpClientConfig{
		HttpTimeout:  config.HttpClientConfig.Timeout,
		MaxIdleConns: config.HttpClientConfig.MaxIdleConns}
	httpClientConfig.SetProxyUrl(config.HttpClientConfig.Proxy)
	builder := builder.NewAPIBuilder2(httpClientConfig).APIKey(config.Binance.ApiKey).APISecretkey(config.Binance.SecretKey)
	RestApi = builder.Build(goex.BINANCE)
	log.Printf("[%v] exchange api created", RestApi.GetExchangeName())

	var err error
	WsApi, err = builder.BuildSpotWs(goex.BINANCE)
	if err != nil {
		log.Fatalf("build spot ws api failed: %v", err.Error())
	}

	InitWsApiCallback()

	log.Printf("[%v] exchange spot ws api created", RestApi.GetExchangeName())

	WalletApi, err = builder.BuildWallet(goex.BINANCE)
	if err != nil {
		log.Fatalf("build wallet api failed: %v", err.Error())
	}
	log.Printf("[%v] exchange wallet api created", RestApi.GetExchangeName())
}

func InitWsApiCallback() {
	WsApi.TickerCallback(func(ticker *goex.Ticker) {
		log.Printf("ticker: %+v\n", ticker)
	})

	WsApi.DepthCallback(func(depth *goex.Depth) {
		log.Printf("depth: %+v\n", depth)
	})

	WsApi.TradeCallback(func(trade *goex.Trade) {
		log.Printf("trade: %+v\n", trade)
	})
}
