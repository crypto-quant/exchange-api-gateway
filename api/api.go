package api

import (
	"log"

	"github.com/crypto-quant/exchange-api-gateway/config"
	"github.com/crypto-quant/exchange-api-gateway/zmq"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/binance"
	"github.com/nntaoli-project/goex/builder"
)

var (
	RestApi              goex.API
	WsApi                goex.SpotWsApi
	WalletApi            goex.WalletApi
	BinanceUserDataWsApi *binance.UserDataWs
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

	WalletApi, err = builder.BuildWallet(goex.BINANCE)
	if err != nil {
		log.Fatalf("build wallet api failed: %v", err.Error())
	}
	log.Printf("[%v] exchange wallet api created", RestApi.GetExchangeName())

	binanceRestApi, ok := RestApi.(*binance.Binance)
	if ok {
		BinanceUserDataWsApi = binance.NewUserDataWs(binanceRestApi)
	}

	InitWsApiCallback()

	log.Printf("[%v] exchange spot ws api created", RestApi.GetExchangeName())

}

func InitWsApiCallback() {
	WsApi.TickerCallback(func(ticker *goex.Ticker) {
		// log.Printf("ticker: %+v\n", ticker)
		// zmq.PubJson("ticker", ticker)
		zmq.PubPBTicker(ticker)
	})

	WsApi.DepthCallback(func(depth *goex.Depth) {
		// log.Printf("depth: %+v\n", depth)
		// zmq.PubJson("depth", depth)
		zmq.PubPBDepth(depth)
	})

	WsApi.TradeCallback(func(trade *goex.Trade) {
		// log.Printf("trade: %+v\n", trade)
		// zmq.PubJson("trade", trade)
	})

	if BinanceUserDataWsApi != nil {
		BinanceUserDataWsApi.AccountInfoCallback(func(accountInfo *binance.AccountInfo) {
			log.Printf("account info: %+v\n", accountInfo)
		})

		BinanceUserDataWsApi.AccountPositionCallback(func(accountPosition *binance.AccountPosition) {
			log.Printf("account position: %+v\n", accountPosition)
			zmq.PubPBAccountPosition(accountPosition)
		})

		BinanceUserDataWsApi.ListStatusCallback(func(listStatus *binance.ListStatus) {
			log.Printf("list status : %+v\n", listStatus)
		})

		BinanceUserDataWsApi.BalanceUpdateCallback(func(balanceUpdate *binance.BalanceUpdate) {
			log.Printf("balance update: %+v\n", balanceUpdate)
		})

		BinanceUserDataWsApi.OrderUpdateCallback(func(orderUpdate *binance.OrderUpdate) {
			log.Printf("order update: %+v\n", orderUpdate)
			zmq.PubPBOrderUpdate(orderUpdate)
		})
	}
}
