package model

type JSONRPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
}

type TradeVolumesResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Result  []TradeVolume `json:"result"`
	UsIn    int64         `json:"usIn"`
	UsOut   int64         `json:"usOut"`
	UsDiff  int           `json:"usDiff"`
	Testnet bool          `json:"testnet"`
}

type TradeVolume struct {
	SpotVolume    float64 `json:"spot_volume"`
	PutsVolume    float64 `json:"puts_volume"`
	FuturesVolume float64 `json:"futures_volume"`
	CurrencyPair  string  `json:"currency_pair"`
	Currency      string  `json:"currency"`
	CallsVolume   float64 `json:"calls_volume"`
}

type JSONRPCMessage struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  JSONRPCParams `json:"params"`
}

type JSONRPCParams struct {
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
}

type JSONRPCMessagePriceIndex struct {
	Jsonrpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  JSONRPCParamsPriceIndex `json:"params"`
}

type JSONRPCParamsPriceIndex struct {
	Channel string         `json:"channel"`
	Data    PriceIndexData `json:"data"`
}

type PriceIndexData struct {
	IndexName string  `json:"index_name"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type JSONRPCMessageTrade struct {
	Jsonrpc string             `json:"jsonrpc"`
	Method  string             `json:"method"`
	Params  JSONRPCParamsTrade `json:"params"`
}

type JSONRPCParamsTrade struct {
	Channel string      `json:"channel"`
	Data    []TradeData `json:"data"`
}

type TradeData struct {
	Contracts      float64 `json:"contracts"`
	TradeID        string  `json:"trade_id"`
	TickDirection  int     `json:"tick_direction"`
	MarkPrice      float64 `json:"mark_price"`
	TradeSeq       int     `json:"trade_seq"`
	Direction      string  `json:"direction"`
	InstrumentName string  `json:"instrument_name"`
	IndexPrice     float64 `json:"index_price"`
	Amount         float64 `json:"amount"`
	Price          float64 `json:"price"`
	IV             float64 `json:"iv"`
	Timestamp      int64   `json:"timestamp"`
}
