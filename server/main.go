package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yashpatel20/bitcoinflow/server/model"
)

// TODO add eth support
var subscribtions = map[string]interface{}{
	"jsonrpc": "2.0",
	"id":      1,
	"method":  "public/subscribe",
	"params": map[string][]string{
		"channels": {
			"deribit_price_index.btc_usd",
			"deribit_price_index.eth_usd",
			"trades.option.BTC.100ms",
			"trades.option.ETH.100ms",
		},
	},
}

var dataChan = make(chan []byte)
var feedChan = make(chan []byte)

//var volumeData = make(map[string]float64)

func fetchData() {
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial("wss://www.deribit.com/ws/api/v2", nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	err = conn.WriteJSON(subscribtions)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	//TODO fix volume request
	volumeRequest := model.JSONRPCRequest{
		Jsonrpc: "2.0",
		Id:      2,
		Method:  "public/get_trade_volumes",
	}
	err = conn.WriteJSON(volumeRequest)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				volumeRequest := model.JSONRPCRequest{
					Jsonrpc: "2.0",
					Id:      2,
					Method:  "public/get_trade_volumes",
				}
				err := conn.WriteJSON(volumeRequest)
				if err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Println(string(message))
		dataChan <- message
	}
}

type OptionDetails struct {
	Expiry      string
	StrikePrice string
	OptionType  string
}

func ParseOptionIdentifier(identifier string, index string) (OptionDetails, error) {
	var details OptionDetails

	pattern := fmt.Sprintf("^(%s)-(\\d{1,2}[A-Z]{3}\\d{2})-(\\d+)-(P|C)$", index)
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(identifier)

	if len(matches) != 5 {
		return details, fmt.Errorf("invalid option identifier format")
	}

	var optionType string
	if matches[4] == "C" {
		optionType = "CALL"
	} else {
		optionType = "PUT"
	}
	details = OptionDetails{
		Expiry:      matches[2],
		StrikePrice: matches[3],
		OptionType:  optionType,
	}

	return details, nil
}

func transformData() {
	for {
		dataMsg := <-dataChan
		var rawMsg map[string]interface{}
		err := json.Unmarshal(dataMsg, &rawMsg)
		if err != nil {
			log.Fatalf("Error unmarshaling to map: %v", err)
		}
		if rawMsg["id"] == 2.0 {
			var tradeVolumeMsg model.TradeVolumesResponse
			err := json.Unmarshal(dataMsg, &tradeVolumeMsg)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}
			var btcVolData model.TradeVolume
			var ethVolData model.TradeVolume
			for _, x := range tradeVolumeMsg.Result {
				if x.Currency == "BTC" {
					btcVolData = x
				}
				if x.Currency == "ETH" {
					ethVolData = x
				}
			}

			btcVolFeed := model.Volume{
				CallVolume:   btcVolData.CallsVolume,
				PutVolume:    btcVolData.PutsVolume,
				PutCallRatio: math.Round(btcVolData.PutsVolume/btcVolData.CallsVolume*1000) / 1000,
			}
			btcVolAppfeed := model.AppFeed{}
			btcVolAppfeed.Index = "BTC"
			btcVolAppfeed.FeedType = "VOLUME"
			btcVolAppfeed.Data = btcVolFeed
			btcBytes, err := json.Marshal(btcVolAppfeed)
			if err != nil {
				log.Fatalf("Error serializing app feed to []byte: %v", err)
			}
			feedChan <- btcBytes

			ethVolFeed := model.Volume{
				CallVolume:   ethVolData.CallsVolume,
				PutVolume:    ethVolData.PutsVolume,
				PutCallRatio: math.Round(ethVolData.PutsVolume/ethVolData.CallsVolume*1000) / 1000,
			}
			ethVolAppfeed := model.AppFeed{}
			ethVolAppfeed.Index = "ETH"
			ethVolAppfeed.FeedType = "VOLUME"
			ethVolAppfeed.Data = ethVolFeed
			ethBytes, err := json.Marshal(ethVolAppfeed)
			if err != nil {
				log.Fatalf("Error serializing eth app feed to []byte: %v", err)
			}
			feedChan <- ethBytes
		} else {
			var msg model.JSONRPCMessage
			err := json.Unmarshal(dataMsg, &msg)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}

			switch msg.Params.Channel {
			case "deribit_price_index.btc_usd":
				var priceIndexMsg model.JSONRPCMessagePriceIndex
				err := json.Unmarshal(dataMsg, &priceIndexMsg)
				if err != nil {
					log.Fatalf("Error unmarshalling JSON: %v", err)
				}
				indexData := priceIndexMsg.Params.Data
				appfeed := model.AppFeed{}
				appfeed.Index = "BTC"
				indexPrice := model.Price{
					IndexName:  indexData.IndexName,
					IndexPrice: indexData.Price,
				}
				appfeed.FeedType = "INDEX_PRICE"
				appfeed.Data = indexPrice
				bytes, err := json.Marshal(appfeed)
				if err != nil {
					log.Fatalf("Error serializing app feed to []byte: %v", err)
				}
				feedChan <- bytes
			case "trades.option.BTC.100ms":
				var tradeMsg model.JSONRPCMessageTrade
				err := json.Unmarshal(dataMsg, &tradeMsg)
				if err != nil {
					log.Fatalf("Error unmarshalling JSON: %v", err)
				}
				trades := tradeMsg.Params.Data
				for _, option := range trades {
					appfeed := model.AppFeed{}
					index := "BTC"
					appfeed.Index = index
					optionDetails, _ := ParseOptionIdentifier(option.InstrumentName, index)
					var dir string
					if option.Direction == "buy" {
						dir = "BUY"
					} else {
						dir = "SELL"
					}
					msg := model.Option{
						Id:   option.TradeID,
						Time: option.Timestamp,
						//Tick:   option.TickDirection,
						Dir:    dir,
						Cp:     optionDetails.OptionType,
						Expiry: optionDetails.Expiry,
						Strike: optionDetails.StrikePrice,
						Spot:   option.IndexPrice,
						Price:  option.Price,
						Size:   option.Contracts,
						Prem:   math.Round(option.Price*option.Contracts*10000) / 10000,
						Iv:     option.IV,
					}
					if msg.Prem == 0.0 {
						continue
					}
					appfeed.FeedType = "OPTION"
					appfeed.Data = msg
					bytes, err := json.Marshal(appfeed)
					if err != nil {
						log.Fatalf("Error serializing app feed to []byte: %v", err)
					}
					feedChan <- bytes
				}
			case "deribit_price_index.eth_usd":
				var priceIndexMsg model.JSONRPCMessagePriceIndex
				err := json.Unmarshal(dataMsg, &priceIndexMsg)
				if err != nil {
					log.Fatalf("Error unmarshalling JSON: %v", err)
				}
				indexData := priceIndexMsg.Params.Data
				appfeed := model.AppFeed{}
				appfeed.Index = "ETH"
				indexPrice := model.Price{
					IndexName:  indexData.IndexName,
					IndexPrice: indexData.Price,
				}
				appfeed.FeedType = "INDEX_PRICE"
				appfeed.Data = indexPrice
				bytes, err := json.Marshal(appfeed)
				if err != nil {
					log.Fatalf("Error serializing app feed to []byte: %v", err)
				}
				feedChan <- bytes
			case "trades.option.ETH.100ms":
				var tradeMsg model.JSONRPCMessageTrade
				err := json.Unmarshal(dataMsg, &tradeMsg)
				if err != nil {
					log.Fatalf("Error unmarshalling JSON: %v", err)
				}
				trades := tradeMsg.Params.Data
				for _, option := range trades {
					appfeed := model.AppFeed{}
					index := "ETH"
					appfeed.Index = index
					optionDetails, _ := ParseOptionIdentifier(option.InstrumentName, index)
					var dir string
					if option.Direction == "buy" {
						dir = "BUY"
					} else {
						dir = "SELL"
					}
					msg := model.Option{
						Id:   option.TradeID,
						Time: option.Timestamp,
						//Tick:   option.TickDirection,
						Dir:    dir,
						Cp:     optionDetails.OptionType,
						Expiry: optionDetails.Expiry,
						Strike: optionDetails.StrikePrice,
						Spot:   option.IndexPrice,
						Price:  option.Price,
						Size:   option.Contracts,
						Prem:   math.Round(option.Price*option.Contracts*10000) / 10000,
						Iv:     option.IV,
					}
					if msg.Prem == 0.0 {
						continue
					}
					appfeed.FeedType = "OPTION"
					appfeed.Data = msg
					bytes, err := json.Marshal(appfeed)
					if err != nil {
						log.Fatalf("Error serializing app feed to []byte: %v", err)
					}
					feedChan <- bytes
				}
			}
		}
	}

}

// TODO Fix Error reading message: read tcp 100.64.100.6:63004->193.72.79.190:443: read: operation timed out
func setupServer() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error during connection upgrade:", err)
			return
		}

		defer conn.Close()
		for {
			msg := <-feedChan
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	go fetchData()
	go transformData()
	setupServer()
}
