package common

import (
	"fmt"
	"strings"

	"github.com/nntaoli-project/goex"
)

const Separater = "-"

func ParseTradingPair(pair string) goex.CurrencyPair {
	return goex.NewCurrencyPair3(pair, Separater)
}

func AdaptSymbolToTradingPair(symbol string) string {
	symbol = strings.ToUpper(symbol)

	if strings.HasSuffix(symbol, "USD") {
		return fmt.Sprintf("%s-USD", strings.TrimSuffix(symbol, "USD"))
	}

	if strings.HasSuffix(symbol, "USDT") {
		return fmt.Sprintf("%s-USDT", strings.TrimSuffix(symbol, "USDT"))
	}

	if strings.HasSuffix(symbol, "PAX") {
		return fmt.Sprintf("%s-PAX", strings.TrimSuffix(symbol, "PAX"))
	}

	if strings.HasSuffix(symbol, "BTC") {
		return fmt.Sprintf("%s-BTC", strings.TrimSuffix(symbol, "BTC"))
	}

	return ""
}


