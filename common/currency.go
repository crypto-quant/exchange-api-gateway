package common

import "github.com/nntaoli-project/goex"

const Separater = "-"

func ParseTradingPair(pair string) goex.CurrencyPair {
	return goex.NewCurrencyPair3(pair, Separater)
}
