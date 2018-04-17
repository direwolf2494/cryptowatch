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

// Assets returns all assets (in no particular order).
func Assets() ([]Asset, error) {
	var assets []Asset
	res, err := request(indexes["Assets"])

	if res != nil {
		json.Unmarshal(res, &assets)
	}
	return assets, err
}

// AssetMarkets returns all markets which have this asset as a base or quote.
func AssetMarkets(asset string) (DetailedAsset, error) {
	var markets DetailedAsset
	url := fmt.Sprintf(indexes["Asset"], asset)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &markets)
	}
	return markets, err
}

// Pairs returns all pairs (in no particular order).
func Pairs() ([]Pair, error) {
	var pairs []Pair
	res, err := request(indexes["Pairs"])

	if res != nil {
		err = json.Unmarshal(res, &pairs)
	}

	return pairs, err
}

// PairMarkets lists all markets for this pair.
func PairMarkets(pair string) (PairMarket, error) {
	var markets PairMarket
	url := fmt.Sprintf(indexes["Pair"], pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &markets)
	}

	return markets, err
}

// Exchanges returns a list of all supported exchanges.
func Exchanges() ([]GeneralExchange, error) {
	var exchanges []GeneralExchange
	res, err := request(indexes["Exchanges"])

	if res != nil {
		err = json.Unmarshal(res, &exchanges)
	}

	return exchanges, err
}

// Exchange returns a single exchange, with associated routes.
func Exchange(name string) (DetailedExchange, error) {
	var exchange DetailedExchange
	url := fmt.Sprintf(indexes["Exchange"], name)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &exchange)
	}

	return exchange, err
}

// Markets returns a list of all supported markets.
func Markets() ([]GeneralMarket, error) {
	var markets []GeneralMarket
	res, err := request(indexes["Markets"])

	if res != nil {
		err = json.Unmarshal(res, &markets)
	}

	return markets, err
}

// Market returns a single market, with associated routes.
func Market(exchange, pair string) (DetailedMarket, error) {
	var market DetailedMarket
	url := fmt.Sprintf(indexes["Market"], exchange, pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &market)
	}

	return market, err
}

// MarketPrice returns a market’s last price.
func MarketPrice(exchange, pair string) (float64, error) {
	var price float64
	url := fmt.Sprintf(indexes["MarketPrice"], exchange, pair)
	res, err := request(url)

	if res != nil {
		var resp map[string]float64
		err = json.Unmarshal(res, &resp)

		if err == nil {
			price = resp["price"]
		}
	}

	return price, err
}

// MarketSummary returns a market’s last price as well as other stats based on a 24-hour sliding window.
func MarketSummary(exchange, pair string) (Summary, error) {
	var summary Summary
	url := fmt.Sprintf(indexes["MarketSummary"], exchange, pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &summary)
	}

	return summary, err
}

// Trades returns a market’s most recent trades, incrementing chronologically.
func Trades(exchange, pair string) ([]Trade, error) {
	var trades []Trade
	url := fmt.Sprintf(indexes["MarketTrades"], exchange, pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &trades)
	}

	return trades, err
}

// OrderBook returns a market’s order book.
func OrderBook(exchange, pair string) (MarketOrderBook, error) {
	var orderbook MarketOrderBook
	url := fmt.Sprintf(indexes["MarketOrderBook"], exchange, pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &orderbook)
	}

	return orderbook, err
}

// Ohlc returns a market’s OHLC candlestick data. Returns data as lists of lists of numbers for each time period integer.
func Ohlc(exchange, pair string) (OHLC, error) {
	var ohlc OHLC
	url := fmt.Sprintf(indexes["MarketOHLC"], exchange, pair)
	res, err := request(url)

	if res != nil {
		err = json.Unmarshal(res, &ohlc)
	}

	return ohlc, err
}

// AggregratePrices returns the current price for all supported markets. Some values may be out of date by a few seconds.
func AggregratePrices() (AggregratePrice, error) {
	var prices AggregratePrice
	res, err := request(indexes["AggregratePrices"])

	if res != nil {
		err = json.Unmarshal(res, &prices)
	}

	return prices, err
}

// AggregrateSummaries returns the market summary for all supported markets. Some values may be out of date by a few seconds.
func AggregrateSummaries() (AggregrateSummary, error) {
	var summaries AggregrateSummary
	res, err := request(indexes["AggregrateSummaries"])

	if res != nil {
		err = json.Unmarshal(res, &summaries)
	}

	return summaries, err
}

func request(url string) ([]byte, error) {
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
		results, err := json.Marshal(results["result"])
		return results, err
	}
}
