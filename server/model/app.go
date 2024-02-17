package model

type AppFeed struct {
	FeedType string      `json:"feedType"`
	Data     interface{} `json:"data"`
}

type Option struct {
	Id   interface{} `json:"id"`
	Time interface{} `json:"time"`
	//Tick   interface{} `json:"tick"`
	Dir    interface{} `json:"dir"`
	Cp     interface{} `json:"cp"`
	Expiry interface{} `json:"expiry"`
	Strike interface{} `json:"strike"`
	Spot   interface{} `json:"spot"`
	Price  interface{} `json:"price"`
	Size   interface{} `json:"size"`
	Prem   interface{} `json:"prem"`
	Iv     interface{} `json:"iv"`
}

type Price struct {
	IndexName  string  `json:"indexName"`
	IndexPrice float64 `json:"indexPrice"`
}

type Volume struct {
	CallVolume   float64 `json:"callVolume"`
	PutVolume    float64 `json:"putVolume"`
	PutCallRatio float64 `json:"putCallRatio"`
}
