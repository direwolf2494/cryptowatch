# Cryptowatch
This go package acts as a wrapper to the various Cryptowatch REST API endpoints found (here | https://cryptowat.ch/docs/api). The returned data is captured and parsed as structs that can be used in any other go package.


## Exported Functions

### Assets
This function returns an array of all crytowatch assets in no particular order.

- Arguments: None
- Returns: []Asset, error
- Invocation:
```go
assets, err := Assets()
```

Asset Definition:
```go
type Asset struct {
    Symbol string
    Name   string
    Fiat   bool
    Route  string
}
```

### AssetMarkets
This function returns information on an asset and all the markets the asset belongs to.

- Arguments: `asset string`
- Returns: DetailedAsset, error
- Invocation: 
```go
name := "btc"
asset, err := AssetMarkets(name)
```

- DetailedAsset Definition:
```go
type DetailedAsset struct {
    ID      int
    Symbol  string
    Name    string
    Fiat    bool
    Markets struct {
        Base  []AssetMarket
	Quote []AssetMarket
    }
}

type AssetMarket struct {
    Exchange string
    Pair     string
    Active   bool
    Route    string
}
```

### Pairs
Returns an array of all pairs in no particular order.

- Arguments: None
- Returns: []Pair, error
- Invocation:
```go
pairs, err := Pairs()
```

- Pair Definition:
```go
type Pair struct {
    Symbol string
    ID     int
    Base   PairData
    Quote  PairData
    Route  string
}

type PairData struct {
    Symbol string
    Name   string
    IsFiat bool 
    Route  string
}

```

### PairMarkets
Returns all the markets that this pair belongs to.

- Argruments: `pair string`
- Returns: PairMarket, error
- Invocation:
```go
pair := "ethbtc"
markets, err := PairMarkets(pair)
```

- PairMarket Definition:
```go
type PairMarket struct {
    Symbol  string
    ID      int  
    Base    PairData
    Quote   PairData
    Route   string
    Markets []AssetMarket
}
```

### Exchanges
Returns all supported exchanges.

- Argruments: None
- Returns: []GeneralExchange, error
- Invocation:
```go
exchanges, err := Exchanges()
```

- Definition:
```go
type GeneralExchange struct {
    Symbol string
    Name   string
    Active bool
    Route  string
}
```


### Exchange
Returns a single exchange, with associated routes.

- Argruments: `name string`
- Returns: DetailedExchange, error
- Invocation:
```go
name := gdax
exchange, err := Exchange(name)
```

- Definition:
```go
type DetailedExchange struct {
    ID     int
    Name   string
    Active bool
    Routes struct {
	Markets string
    }
}
```


### Markets
Returns all supported markets.

- Argruments: None
- Returns: []GeneralMarket, error
- Invocation:
```go
markets, err := Markets()
```

- GeneralMarket Definition:
```go
type GeneralMarket struct {
    Exchange string
    Pair     string
    Active   bool   
    Route    string
}
```


### Market
Returns detailed information for a single market.

- Argruments: `exch, pair string`
- Returns: DetailedMarket, error
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
market, err := Market(exch, pair)
```

- DetailedMarket Definition:
```go
type DetailedMarket struct {
    Exchange string
    Pair     string 
    Active   bool
    Routes   struct {
	Price     string
	Summary   string
	Orderbook string
	Trades    string
	Ohlc      string
    }
}
```


### MarketPrice
Returns the last price for a market.

- Argruments: `exch, pair string`
- Returns: float64, error
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
price, err := MarketPrice(exch, pair)
```

### MarketSummary
Returns a market’s last price as well as other stats based on a 24-hour sliding window.

- Argruments: `exch, pair string`
- Returns: Summary, error
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
summary, err := MarketSummary(exch, pair)
```

- Summary Definition:
```go
type Summary struct {
    Price struct {
	Last   float64
	High   float64
	Low    float64
	Change struct {
	    Percentage float64	
	    Absolute   float64
	}
    }
    Volume float64
}
```


### Trades
Returns a market’s most recent trades, incrementing chronologically. Each Trade consists of a slice of length four (4). The attributes of each index is shown below:

* `[ ID, Timestamp, Price, Amount ]`

- Argruments: `exch, pair string`
- Returns: []Trade
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
trades, err := Trades(exch, pair)
```

- Trade Definition:
```go
type Trade []float64
```

### OrderBook
Returns a market’s order book. Each Ask/Bid consists of a slice of length two (2). The attribute for each index is shown below:
* `[ Price, Amount ]`

- Arguments: `exch, pair string`
- Returns: MarketOrderBook, error
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
orderbook, err := OrderBook(exch, pair)
```

- MarketOrderBook Defintion:
```go
type MarketOrderBook struct {
    Asks [][]float64
    Bids [][]float64
}
```

### Ohlc
returns a market’s Open, High, Low, Close candlestick data. Returns data as lists of lists of numbers for each time period integer.

- Arguments: `exch, pair string`
- Returns: OHLC, error
- Invocation:
```go
exch, pair := "gdax", "ethbtc"
ohlc, err := Ohlc(exch, pair)
```

- Defintion:
```go
type OHLC map[string][][]float64
```

### AggregratePrices
Returns the current price for all supported markets. Some values may be out of date by a few seconds.

- Arguments: None
- Returns: AggregratePrice, error
- Invocation:
```go
prices, err := AggregratePrices(exch, pair)
```

- Defintion:
```go
type AggregratePrice map[string]float64
```

### AggregrateSummaries

- Arguments: None
- Returns: AggregrateSummary, error
- Invocation:
```go
summaries, err := AggregrateSummaries(exch, pair)
```

- Defintion:
```go
type AggregrateSummary map[string]Summary
```

*N.B.* This project is licensed under the terms of the MIT license.
