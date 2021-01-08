### api index

[GIN-debug] GET /ping --> main.main.func1 (3 handlers)
[GIN-debug] POST /balance --> github.com/crypto-quant/exchange-api-gateway/handler.GetBalance (3 handlers)
[GIN-debug] POST /limit_order --> github.com/crypto-quant/exchange-api-gateway/handler.LimitOrder (3 handlers)
[GIN-debug] POST /market_order --> github.com/crypto-quant/exchange-api-gateway/handler.MarketOrder (3 handlers)
[GIN-debug] POST /cancel_order --> github.com/crypto-quant/exchange-api-gateway/handler.CancelOrder (3 handlers)
[GIN-debug] POST /cancel_all_orders --> github.com/crypto-quant/exchange-api-gateway/handler.CancelAllOrders (3 handlers)
[GIN-debug] POST /get_order --> github.com/crypto-quant/exchange-api-gateway/handler.GetOrder (3 handlers)
[GIN-debug] POST /get_order_history --> github.com/crypto-quant/exchange-api-gateway/handler.GetOrderHistory (3 handlers)
[GIN-debug] POST /get_unfinished_order --> github.com/crypto-quant/exchange-api-gateway/handler.GetUnfinishedOrders (3 handlers)
[GIN-debug] POST /get_trading_pairs --> github.com/crypto-quant/exchange-api-gateway/handler.GetTradingPairs (3 handlers)
[GIN-debug] POST /get_ticker --> github.com/crypto-quant/exchange-api-gateway/handler.GetTicker (3 handlers)
[GIN-debug] POST /get_depth --> github.com/crypto-quant/exchange-api-gateway/handler.GetDepth (3 handlers)

## limit order

```shell
## side: 1 buy, 2 sell
curl -X POST http://localhost:8080/limit_order -H 'content-type: application/json' -d '{ "side": 2, "pair": "BTC-USDT", "price": "45000", "amount": "0.01"}'
```

```json
{
  "data": {
    "Price": 50000,
    "Amount": 0.01,
    "AvgPrice": 0,
    "DealAmount": 0,
    "Fee": 0,
    "Cid": "",
    "OrderID2": "4180096286",
    "OrderID": 4180096286,
    "Status": 0,
    "Currency": {
      "CurrencyA": {
        "Symbol": "BTC",
        "Desc": "https://bitcoin.org/"
      },
      "CurrencyB": {
        "Symbol": "USDT",
        "Desc": ""
      },
      "AmountTickSize": 0,
      "PriceTickSize": 0
    },
    "Side": 2,
    "Type": "",
    "OrderType": 0,
    "OrderTime": 1609947571469,
    "FinishedTime": 0
  }
}
```

```json
{
  "error": "Key: 'Order.Vol' Error:Field validation for 'Vol' failed on the 'required' tag"
}
```

### cancel order

```shell
curl -X POST http://localhost:8080/cancel_order  -H 'content-type: application/json' -d '{ "order_id": "4180096286", "pair": "BTC-USDT" }'
```

```json
{
  "data": true
}
```

### get order

```shell
curl -X POST http://localhost:8080/get_order  -H 'content-type: application/json' -d '{ "order_id": "4180096286", "pair": "BTC-USDT" }'
```

```json
{
  "data": {
    "Price": 50000,
    "Amount": 0.01,
    "AvgPrice": 0,
    "DealAmount": 0,
    "Fee": 0,
    "Cid": "CUQgF5VBGUTjzgDt8LlmVG",
    "OrderID2": "4.180096286e+09",
    "OrderID": 4180096286,
    "Status": 3,
    "Currency": {
      "CurrencyA": {
        "Symbol": "BTC",
        "Desc": "https://bitcoin.org/"
      },
      "CurrencyB": {
        "Symbol": "USDT",
        "Desc": ""
      },
      "AmountTickSize": 0,
      "PriceTickSize": 0
    },
    "Side": 2,
    "Type": "",
    "OrderType": 0,
    "OrderTime": 1609947571469,
    "FinishedTime": 1609948546316
  }
}
```

### get unfinished orders

```shell
curl -X POST http://localhost:8080/get_unfinished_order  -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
```

```json
{
  "data": [
    {
      "Price": 45000,
      "Amount": 0.01,
      "AvgPrice": 0,
      "DealAmount": 0,
      "Fee": 0,
      "Cid": "mtBXJq8W3hUDp6uo3JoMr1",
      "OrderID2": "4.180071516e+09",
      "OrderID": 4180071516,
      "Status": 0,
      "Currency": {
        "CurrencyA": {
          "Symbol": "BTC",
          "Desc": "https://bitcoin.org/"
        },
        "CurrencyB": {
          "Symbol": "USDT",
          "Desc": ""
        },
        "AmountTickSize": 0,
        "PriceTickSize": 0
      },
      "Side": 2,
      "Type": "",
      "OrderType": 0,
      "OrderTime": 1609947390627,
      "FinishedTime": 1609947390627
    },
    {
      "Price": 38000,
      "Amount": 0.1,
      "AvgPrice": 0,
      "DealAmount": 0,
      "Fee": 0,
      "Cid": "jbTM8KGKai9SOFBAPzBg80",
      "OrderID2": "4.180079391e+09",
      "OrderID": 4180079391,
      "Status": 0,
      "Currency": {
        "CurrencyA": {
          "Symbol": "BTC",
          "Desc": "https://bitcoin.org/"
        },
        "CurrencyB": {
          "Symbol": "USDT",
          "Desc": ""
        },
        "AmountTickSize": 0,
        "PriceTickSize": 0
      },
      "Side": 2,
      "Type": "",
      "OrderType": 0,
      "OrderTime": 1609947437804,
      "FinishedTime": 1609947437804
    }
  ]
}
```

### Cancel all orders

```shell
curl -X POST http://localhost:8080/cancel_all_orders  -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
```

```json
{
  "data": ["4180071516", "4180079391"]
}
```

### Get trading pairs

```shell
curl http://localhost:8080/get_trading_pairs
```

check doc/trading_pairs

### Get ticker

```shell
curl -X POST http://127.0.0.1:8080/get_ticker -H 'content-type: application/json' -d '{ "pair": "BTC-USDT" }'
```

```json
{
  "data": {
    "omitempty": {
      "CurrencyA": {
        "Symbol": "BTC",
        "Desc": "https://bitcoin.org/"
      },
      "CurrencyB": {
        "Symbol": "USDT",
        "Desc": ""
      },
      "AmountTickSize": 0,
      "PriceTickSize": 0
    },
    "last": "38884.02",
    "buy": "38884.02",
    "sell": "38884.03",
    "high": "40365",
    "low": "36500",
    "vol": "135714.770006",
    "date": 1610099152
  }
}
```

### Get depth

```shell
curl -X POST http://127.0.0.1:8080/get_depth -H 'content-type: application/json' -d '{ "pair": "BTC-USDT", "size": 100 }'
```

```json
{
  "data": {
    "Pair": {
      "CurrencyA": {
        "Symbol": "BTC",
        "Desc": "https://bitcoin.org/"
      },
      "CurrencyB": {
        "Symbol": "USDT",
        "Desc": ""
      },
      "AmountTickSize": 0,
      "PriceTickSize": 0
    },
    "UTime": "2021-01-08T18:19:26.113863+08:00",
    "AskList": [
      {
        "Price": 39664.99,
        "Amount": 0.210001
      },
      {
        "Price": 39664.87,
        "Amount": 0.030253
      },
      {
        "Price": 39662.55,
        "Amount": 0.031497
      },
      {
        "Price": 39662.54,
        "Amount": 0.006615
      },
      {
        "Price": 39661.59,
        "Amount": 0.33849
      }
    ],
    "BidList": [
      {
        "Price": 39661.58,
        "Amount": 0.741277
      },
      {
        "Price": 39660.31,
        "Amount": 0.336257
      },
      {
        "Price": 39660.3,
        "Amount": 1.997482
      },
      {
        "Price": 39660.17,
        "Amount": 0.487192
      },
      {
        "Price": 39658.88,
        "Amount": 0.048207
      }
    ]
  }
}
```

### Get klines

```shell
curl -X POST http://127.0.0.1:8080/get_klines -H 'content-type: application/json' -d '{ "pair": "BTC-USDT", "size": 2, "period": 13, "start_time": 1577887978 }'
```

```json
{
  "data": [
    {
      "Pair": {
        "CurrencyA": {
          "Symbol": "BTC",
          "Desc": "https://bitcoin.org/"
        },
        "CurrencyB": {
          "Symbol": "USDT",
          "Desc": ""
        },
        "AmountTickSize": 0,
        "PriceTickSize": 0
      },
      "Timestamp": 1502928000,
      "Open": 4261.48,
      "Close": 4427.3,
      "High": 4485.39,
      "Low": 4261.32,
      "Vol": 145.708747
    },
    {
      "Pair": {
        "CurrencyA": {
          "Symbol": "BTC",
          "Desc": "https://bitcoin.org/"
        },
        "CurrencyB": {
          "Symbol": "USDT",
          "Desc": ""
        },
        "AmountTickSize": 0,
        "PriceTickSize": 0
      },
      "Timestamp": 1502971200,
      "Open": 4436.06,
      "Close": 4285.08,
      "High": 4485.39,
      "Low": 4200.74,
      "Vol": 649.44163
    }
  ]
}
```

```go
// k线周期
const (
	KLINE_PERIOD_1MIN = 1 + iota
	KLINE_PERIOD_3MIN
	KLINE_PERIOD_5MIN
	KLINE_PERIOD_15MIN
	KLINE_PERIOD_30MIN
	KLINE_PERIOD_60MIN
	KLINE_PERIOD_1H
	KLINE_PERIOD_2H
	KLINE_PERIOD_3H
	KLINE_PERIOD_4H
	KLINE_PERIOD_6H
	KLINE_PERIOD_8H
	KLINE_PERIOD_12H
	KLINE_PERIOD_1DAY = 13
	KLINE_PERIOD_3DAY
	KLINE_PERIOD_1WEEK
	KLINE_PERIOD_1MONTH
	KLINE_PERIOD_1YEAR
)

// binanace支持的周期类型
1m
3m
5m
15m
30m
1h
2h
4h
6h
8h
12h
1d
3d
1w
1M
```
