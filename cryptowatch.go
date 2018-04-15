// Package that implements wrappers that request information
// from cryptowatch's public market rest api

package cryptowatch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var base = "https://api.cryptowat.ch/"

// cryptowatch indexes
var indexes = map[string]string{
	"Assets":              base + "assets",
	"Asset":               base + "assets/%v",
	"Pairs":               base + "pairs",
	"Pair":                base + "pairs/%v",
	"Exchanges":           base + "exchanges",
	"Exchange":            base + "exchanges/%v",
	"Markets":             base + "markets",
	"Market":              base + "markets/%v/%v",
	"MarketPrice":         base + "markets/%v/%v/price",
	"MarketSummary":       base + "markets/%v/%v/summary",
	"MarketTrades":        base + "markets/%v/%v/trades",
	"MarketOrderBook":     base + "markets/%v/%v/orderbook",
	"MarketOHLC":          base + "markets/%v/%v/ohlc",
	"AggregratePrices":    base + "markets/prices",
	"AggregrateSummaries": base + "markets/summaries",
}

// Assets returns all assets (in no particular order).
func Assets() ([]interface{}, error) {
	var tickers []interface{}
	res, err := request(indexes["Assets"])

	if res != nil {
		tickers = res.([]interface{})
	}
	return tickers, err
}

// AssetMarkets returns all markets which have this asset as a base or quote.
func AssetMarkets(asset string) (map[string]interface{}, error) {
	markets := make(map[string]interface{})
	url := fmt.Sprintf(indexes["Asset"], asset)
	res, err := request(url)

	if res != nil {
		markets = res.(map[string]interface{})
	}
	return markets, err
}

// Pairs returns all pairs (in no particular order).
func Pairs() ([]interface{}, error) {
	var pairs []interface{}
	res, err := request(indexes["Pairs"])

	if res != nil {
		pairs = res.([]interface{})
	}

	return pairs, err
}

// PairMarkets lists all markets for this pair.
func PairMarkets(pair string) (map[string]interface{}, error) {
	var markets map[string]interface{}
	url := fmt.Sprintf(indexes["Pair"], pair)
	res, err := request(url)

	if res != nil {
		markets = res.(map[string]interface{})
	}

	return markets, err
}

// Exchanges returns a list of all supported exchanges.
func Exchanges() ([]interface{}, error) {
	var exchanges []interface{}
	res, err := request(indexes["Exchanges"])

	if res != nil {
		exchanges = res.([]interface{})
	}

	return exchanges, err
}

// Exchange returns a single exchange, with associated routes.
func Exchange(name string) (map[string]interface{}, error) {
	var exchange map[string]interface{}
	url := fmt.Sprintf(indexes["Exchange"], name)
	res, err := request(url)

	if res != nil {
		exchange = res.(map[string]interface{})
	}

	return exchange, err
}

// Markets returns a list of all supported markets.
func Markets() (interface{}, error) {
	var markets interface{}
	res, err := request(indexes["Markets"])

	if res != nil {
		markets = res.([]interface{})
	}

	return markets, err
}

// Market returns a single market, with associated routes.
func Market(exchange, pair string) (map[string]interface{}, error) {
	var market map[string]interface{}
	url := fmt.Sprintf(indexes["Market"], exchange, pair)
	res, err := request(url)

	if res != nil {
		market = res.(map[string]interface{})
	}

	return market, err
}

// MarketPrice returns a market’s last price.
func MarketPrice(exchange, pair string) (map[string]interface{}, error) {
	var price map[string]interface{}
	url := fmt.Sprintf(indexes["MarketPrice"], exchange, pair)
	res, err := request(url)

	if res != nil {
		price = res.(map[string]interface{})
	}

	return price, err
}

// MarketSummary returns a market’s last price as well as other stats based on a 24-hour sliding window.
func MarketSummary(exchange, pair string) (map[string]interface{}, error) {
	var summary map[string]interface{}
	url := fmt.Sprintf(indexes["MarketSummary"], exchange, pair)
	res, err := request(url)

	if res != nil {
		summary = res.(map[string]interface{})
	}

	return summary, err
}

// Trades returns a market’s most recent trades, incrementing chronologically.
func Trades(exchange, pair string) ([]interface{}, error) {
	var trades []interface{}
	url := fmt.Sprintf(indexes["MarketTrades"], exchange, pair)
	res, err := request(url)

	if res != nil {
		trades = res.([]interface{})
	}

	return trades, err
}

// OrderBook returns a market’s order book.
func OrderBook(exchange, pair string) (map[string]interface{}, error) {
	var orderbook map[string]interface{}
	url := fmt.Sprintf(indexes["MarketOrderBook"], exchange, pair)
	res, err := request(url)

	if res != nil {
		orderbook = res.(map[string]interface{})
	}

	return orderbook, err
}

// Ohlc returns a market’s OHLC candlestick data. Returns data as lists of lists of numbers for each time period integer.
func Ohlc(exchange, pair string) (map[string]interface{}, error) {
	var ohlc map[string]interface{}
	url := fmt.Sprintf(indexes["MarketOHLC"], exchange, pair)
	res, err := request(url)

	if res != nil {
		ohlc = res.(map[string]interface{})
	}

	return ohlc, err
}

// AggregratePrices returns the current price for all supported markets. Some values may be out of date by a few seconds.
func AggregratePrices() (map[string]interface{}, error) {
	var prices map[string]interface{}
	res, err := request(indexes["AggregratePrices"])

	if res != nil {
		prices = res.(map[string]interface{})
	}

	return prices, err
}

// AggregrateSummaries returns the market summary for all supported markets. Some values may be out of date by a few seconds.
func AggregrateSummaries() (map[string]interface{}, error) {
	var summaries map[string]interface{}
	res, err := request(indexes["AggregrateSummaries"])

	if res != nil {
		summaries = res.(map[string]interface{})
	}

	return summaries, err
}

func request(url string) (interface{}, error) {
	var data interface{}
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	// convert the response to a usable format
	results := data.(map[string]interface{})

	switch {
	case resp.StatusCode == 429:
		ttr := 60 - time.Now().Minute()
		message := "Too Many Requests. Allowance resets in " + string(ttr) + " minutes."
		return nil, errors.New(message)
	case resp.StatusCode != 200:
		message := (results["error"]).(string)
		return nil, errors.New(message)
	default:
		return results["result"], err
	}
}
