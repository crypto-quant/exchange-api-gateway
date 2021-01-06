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
