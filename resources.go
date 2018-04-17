package cryptowatch

const base = "https://api.cryptowat.ch/"

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

// AssetMarket contains details for a quote or base for a market
type AssetMarket struct {
	Exchange string `json:"exchange"`
	Pair     string `json:"pair"`
	Active   bool   `json:"active,omitempty"`
	Route    string `json:"route"`
}

// PairData contains details for a quote or pair for a specific pair
type PairData struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	IsFiat bool   `json:"isFiat"`
	Route  string `json:"route"`
}

// Asset holds the general data for a cryptowatch asset
type Asset struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Fiat   bool   `json:"fiat"`
	Route  string `json:"route"`
}

// DetailedAsset contains addition information on an asset (Markets)
type DetailedAsset struct {
	ID      int    `json:"id,omitempty"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Fiat    bool   `json:"fiat"`
	Markets struct {
		Base  []AssetMarket `json:"base"`
		Quote []AssetMarket `json:"quote"`
	} `json:"markets"`
}

// Pair contains general information on a crytowatch pair
type Pair struct {
	Symbol string   `json:"symbol"`
	ID     int      `json:"id"`
	Base   PairData `json:"base"`
	Quote  PairData `json:"quote"`
	Route  string   `json:"route"`
}

// PairMarket contains market data for a cryptowatch pair
type PairMarket struct {
	Symbol  string        `json:"symbol"`
	ID      int           `json:"id"`
	Base    PairData      `json:"base"`
	Quote   PairData      `json:"quote"`
	Route   string        `json:"route"`
	Markets []AssetMarket `json:"markets"`
}

// GeneralExchange contains general information on an exchange
type GeneralExchange struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Route  string `json:"route"`
}

// DetailedExchange contains additional information on an exchange
type DetailedExchange struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Routes struct {
		Markets string `json:"markets"`
	} `json:"routes"`
}

// GeneralMarket contains general information for a single market
type GeneralMarket struct {
	Exchange string `json:"exchange"`
	Pair     string `json:"pair"`
	Active   bool   `json:"active"`
	Route    string `json:"route"`
}

// DetailedMarket contains addition routing information for a market
type DetailedMarket struct {
	Exchange string `json:"exchange"`
	Pair     string `json:"pair"`
	Active   bool   `json:"active"`
	Routes   struct {
		Price     string `json:"price"`
		Summary   string `json:"summary"`
		Orderbook string `json:"orderbook"`
		Trades    string `json:"trades"`
		Ohlc      string `json:"ohlc"`
	} `json:"routes"`
}

// Summary contains summary information for a market
type Summary struct {
	Price struct {
		Last   float64 `json:"last"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Change struct {
			Percentage float64 `json:"percentage"`
			Absolute   float64 `json:"absolute"`
		} `json:"change"`
	} `json:"price"`
	Volume float64 `json:"volume"`
}

// Trade contains trading information for an asset
type Trade []float64

// MarketOrderBook contains the ask/bid prices for a market
type MarketOrderBook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// OHLC contains open-high-low-close info for a market
type OHLC map[string][][]float64

// AggregratePrice contains prices on all markets
type AggregratePrice map[string]float64

// AggregrateSummary contains summary for all markets
type AggregrateSummary map[string]Summary
