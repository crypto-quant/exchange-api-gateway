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
