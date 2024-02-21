package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strings"
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
func (i *Identity) Buy() {

}
func (i *Identity) Sell() {

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
	Quantity    int     `json:"quantity"`
	RequestId   string  `json:"requestId"`
	TimeInForce string  `json:"timeInForce"`
}

func (i *Identity) ExecuteOrder(Method int, Quantity int, Price float64, symbol string, instrumentId int) {
	var executePayload ExecutePayload
	executePayload.Legs[0].Symbol = symbol
	executePayload.Legs[0].InstrumentId = instrumentId
	executePayload.Legs[0].PositionEffect = "OPEN"
	executePayload.Legs[0].RatioQuantity = 1
	switch Price {
	case 0:
		executePayload.DirectExchange = true
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

	url := "https://dxtrade.ftmo.com/api/orders/single"
	method := "POST"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "dxtrade.ftmo.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	//req.Header.Add("dnt", "1")
	//req.Header.Add("origin", "https://dxtrade.ftmo.com")
	//req.Header.Add("pragma", "no-cache")
	//req.Header.Add("referer", "https://dxtrade.ftmo.com/")
	//req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"121\", \"Not A(Brand\";v=\"99\"")
	//req.Header.Add("sec-ch-ua-mobile", "?0")
	//req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	//req.Header.Add("sec-fetch-dest", "empty")
	//req.Header.Add("sec-fetch-mode", "cors")
	//req.Header.Add("sec-fetch-site", "same-origin")
	//req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	//req.Header.Add("x-atmosphere-tracking-id", "5ca790a8-5d3e-4fd0-a409-c9cf674e8d84")
	//req.Header.Add("x-csrf-token", "a9779deb-1889-42a7-af60-9abcedd1a435")
	//req.Header.Add("x-requested-with", "XMLHttpRequest")

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
	fmt.Println(string(body))
}

//DXTFID="4a5c1792438ca392"; JSESSIONID=D158AB6886F36A59CEFB4FE770D215EF.jvmroute

func main() {
	identity := Identity{
		Username: "1210003069",
		Password: "2K2=WJ3^6rj5",
		Server:   "ftmo",
	}
	identity.login()
	positions := identity.GetTransactions()
	for _, v := range positions.Body {
		// Divide everything by 1000
		fmt.Println(v.Uid)
	}
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
		Quantity     int         `json:"quantity"`
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
