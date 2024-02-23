package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Identity struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"vendor"`
	Cookies  map[string]string
}

const (
	US30 = 3351
)
const (
	BUY = iota
	SELL
	MARKET = -1
)

func (i *Identity) login() {

	url := "https://dxtrade.ftmo.com/api/auth/login"
	method := "POST"

	payload, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "dxtrade.ftmo.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	i.Cookies = make(map[string]string)
	fmt.Println(string(body), res.Status)
	for _, cookie := range res.Cookies() {
		i.Cookies[cookie.Name] = cookie.Value
	}
	i.EstablishHandshake()
}
func (i *Identity) GetTransactions() *Positions {
	inc_msg := strings.Split(i.EstablishHandshake("POSITIONS"), "|")
	if len(inc_msg) < 2 {
		return nil
	}
	inc_msg2 := inc_msg[1]
	var positions *Positions
	err := json.Unmarshal([]byte(inc_msg2), &positions)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return positions

}
func (i *Identity) EstablishHandshake(kill_msg ...string) string {
	// Websocket
	var kill string
	if len(kill_msg) > 0 {
		kill = kill_msg[0]
	}
	dialer := websocket.Dialer{}
	headers := http.Header{}
	headers.Add("Cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	conn, _, err := dialer.Dial("wss://dxtrade.ftmo.com/client/connector?X-Atmosphere-tracking-id=0&X-Atmosphere-Framework=2.3.2-javascript&X-Atmosphere-Transport=websocket&X-Atmosphere-TrackMessageSize=true&Content-Type=text/x-gwt-rpc;%20charset=UTF-8&X-atmo-protocol=true&sessionState=dx-new&guest-mode=false", headers)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer conn.Close()
	// Send handshake
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return ""
		}
		if strings.Contains(string(message), kill) {
			return string(message)
		}
	}
	return ""
}
func (i *Identity) Buy(Quantity float64, Price float64, symbol string, instrumentId int) {
	i.ExecuteOrder(BUY, Quantity, Price, symbol, instrumentId)
}
func (i *Identity) Sell(Quantity float64, Price float64, symbol string, instrumentId int) {
	i.ExecuteOrder(SELL, Quantity, Price, symbol, instrumentId)
}

type ExecutePayload struct {
	DirectExchange bool `json:"directExchange"`
	Legs           []struct {
		InstrumentId   int    `json:"instrumentId"`
		PositionEffect string `json:"positionEffect"`
		RatioQuantity  int    `json:"ratioQuantity"`
		Symbol         string `json:"symbol"`
	} `json:"legs"`
	LimitPrice  float64 `json:"limitPrice"`
	OrderSide   string  `json:"orderSide"`
	OrderType   string  `json:"orderType"`
	Quantity    float64 `json:"quantity"`
	RequestId   string  `json:"requestId"`
	TimeInForce string  `json:"timeInForce"`
}

func (i *Identity) CloseAllPositions() {
	positions := i.GetTransactions()
	for _, position := range positions.Body {
		i.ClosePosition(position.PositionKey.PositionCode, position.Quantity, 0, position.PositionKey.PositionCode, position.PositionKey.InstrumentId)
	}
}
func (i *Identity) ClosePosition(PositionId string, Quantity float64, Price float64, symbol string, instrumentId int) {
	url := "https://dxtrade.ftmo.com/api/positions/close"
	method := "POST"
	var payload ClosePosition
	legs := make([]struct {
		InstrumentId   int    `json:"instrumentId"`
		PositionCode   string `json:"positionCode"`
		PositionEffect string `json:"positionEffect"`
		RatioQuantity  int    `json:"ratioQuantity"`
		Symbol         string `json:"symbol"`
	}, 1)
	legs[0].InstrumentId = instrumentId
	legs[0].PositionCode = PositionId
	legs[0].PositionEffect = "CLOSING"
	legs[0].RatioQuantity = 1
	legs[0].Symbol = symbol
	payload.Legs = legs
	payload.LimitPrice = 0
	payload.OrderType = "MARKET"
	payload.Quantity = -Quantity
	payload.TimeInForce = "GTC"
	client := &http.Client{}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadJson))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	req.Header.Add("x-csrf-token", i.FetchCSRF())
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res.Status)
}
func (i *Identity) ExecuteOrder(Method int, Quantity float64, Price float64, symbol string, instrumentId int) {
	var executePayload ExecutePayload
	executePayload.DirectExchange = false
	executePayload.Legs = make([]struct {
		InstrumentId   int    `json:"instrumentId"`
		PositionEffect string `json:"positionEffect"`
		RatioQuantity  int    `json:"ratioQuantity"`
		Symbol         string `json:"symbol"`
	}, 1)

	// The generated request ID gwt-uid-931-08b3a3e1-5e92-4db9-9b32-049777c03e17
	executePayload.Legs[0].Symbol = symbol
	executePayload.Legs[0].InstrumentId = instrumentId
	executePayload.Legs[0].PositionEffect = "OPENING"
	executePayload.Legs[0].RatioQuantity = 1
	switch Price {
	case -1:
		executePayload.DirectExchange = false
		executePayload.LimitPrice = Price
	default:
		executePayload.DirectExchange = false
		executePayload.LimitPrice = 0
	}
	executePayload.OrderType = "MARKET"
	switch Method {
	case BUY:
		executePayload.OrderSide = "BUY"
	case SELL:
		executePayload.OrderSide = "SELL"
	}
	executePayload.Quantity = Quantity
	executePayload.TimeInForce = "GTC"
	//931-08b3a3e1-5e92-4db9-9b32-049777c03e17
	executePayload.RequestId = "gwt-uid-931-" + uuid.New().String()
	url := "https://dxtrade.ftmo.com/api/orders/single"
	method := "POST"

	payload, err := json.Marshal(executePayload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	req.Header.Add("x-csrf-token", i.FetchCSRF())
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	// Fetch csrf document.querySelector('meta[name="csrf"]'))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
}
func (i *Identity) FetchCSRF() string {
	url := "https://dxtrade.ftmo.com/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "dxtrade.ftmo.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	// GET THIS     <meta id="csrf-token" name="csrf" content="2813b206-da5f-4271-8385-51e5c427f47b">
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if strings.Contains(string(body), "name=\"csrf\" content=\"") {
		csrf := strings.Split(string(body), "name=\"csrf\" content=\"")
		csrf = strings.Split(csrf[1], "\">")
		return csrf[0]
	}
	return ""
}

//DXTFID="4a5c1792438ca392"; JSESSIONID=D158AB6886F36A59CEFB4FE770D215EF.jvmroute

func main() {
	identity := Identity{
		Username: "1210003069",
		Password: "2K2=WJ3^6rj5",
		Server:   "ftmo",
	}
	identity.login()
	identity.ExecuteOrder(BUY, 0.01, MARKET, "US30.cash", US30)
	time.Sleep(5 * time.Second)
	identity.CloseAllPositions()
}

type ClosePosition struct {
	Legs []struct {
		InstrumentId   int    `json:"instrumentId"`
		PositionCode   string `json:"positionCode"`
		PositionEffect string `json:"positionEffect"`
		RatioQuantity  int    `json:"ratioQuantity"`
		Symbol         string `json:"symbol"`
	} `json:"legs"`
	LimitPrice  int     `json:"limitPrice"`
	OrderType   string  `json:"orderType"`
	Quantity    float64 `json:"quantity"`
	TimeInForce string  `json:"timeInForce"`
}
type Positions struct {
	AccountId string `json:"accountId"`
	Body      []struct {
		Uid         string `json:"uid"`
		AccountId   string `json:"accountId"`
		PositionKey struct {
			InstrumentId int    `json:"instrumentId"`
			PositionCode string `json:"positionCode"`
		} `json:"positionKey"`
		Quantity     float64     `json:"quantity"`
		Cost         float64     `json:"cost"`
		CostBasis    float64     `json:"costBasis"`
		MarginRate   float64     `json:"marginRate"`
		Time         int64       `json:"time"`
		ModifiedTime int64       `json:"modifiedTime"`
		UserLogin    string      `json:"userLogin"`
		TakeProfit   interface{} `json:"takeProfit"`
		StopLoss     interface{} `json:"stopLoss"`
	} `json:"body"`
	Type string `json:"type"`
}
