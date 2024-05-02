package dx // dx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Identity struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Server    string `json:"vendor"`
	AccountId int    `json:"accountId"`
	Cookies   map[string]string
	Csrf      string
}

const (
	ETHUSD     = 3443
	BTCUSD     = 3425
	XRPUSD     = 3404
	DOGEUSD    = 3410
	EURUSD     = 3438
	US30CASH   = 3351
	NEOUSD     = 3370
	AUDCAD     = 3436
	XAUUSD     = 3406
	DASHUSD    = 3445
	USD        = 1201
	CNY        = 1301
	XMRUSD     = 3396
	LTC        = 1401
	USDILS     = 3405
	XAUEUR     = 3412
	MAR        = 1605
	XPTUSD     = 3444
	USDSGD     = 3421
	ADA        = 3367
	USOILCASH  = 3354
	USDTRY     = 3400
	YHOO       = 1515
	MS         = 1584
	AUDNZD     = 3446
	CADJPY     = 3391
	EURCAD     = 3399
	EURRUB     = 3394
	USDHKD     = 3441
	VOD        = 1597
	GBPUSD     = 3440
	KRW        = 801
	THB        = 501
	USDCHF     = 3390
	GBPNZD     = 3418
	TRIP       = 1627
	DXF        = 3374
	ADAUSD     = 3369
	GBP        = 751
	USTN10F    = 3376
	DOGE       = 1567
	PLN        = 551
	USDNOK     = 3417
	NZD        = 901
	LTCUSD     = 3409
	GBPAUD     = 3430
	AAPL       = 1570
	US100CASH  = 3352
	ETFC       = 1556
	IBE        = 3317
	XAUAUD     = 3449
	AMZN       = 1569
	NZDUSD     = 3398
	INR        = 1101
	RL         = 1501
	GBPCHF     = 3432
	CNH        = 1251
	USDJPY     = 3427
	NWS        = 1614
	EURJPY     = 3392
	V          = 3316
	UKOILCASH  = 3357
	YELP       = 1559
	US500CASH  = 3363
	JPM        = 1526
	FRA40CASH  = 3358
	T          = 1640
	GBPJPY     = 3397
	EOS        = 1600
	AUD        = 701
	PEP        = 1633
	NVDA       = 3318
	NATGASF    = 3377
	EURPLN     = 3401
	BLK        = 1631
	USDRUB     = 3429
	MXN        = 851
	WHEAT      = 4954
	SEK        = 401
	NKE        = 1523
	BTC        = 1351
	HUF        = 601
	AUDCHF     = 3395
	EURNZD     = 3414
	USDSEK     = 3423
	CS         = 1646
	US2000CASH = 3356
	HKD        = 251
	LVMH       = 3324
	XPT        = 3308
	TMUS       = 3385
	HK50CASH   = 3362
	YNDX       = 1576
	USDPLN     = 3434
	NZDCAD     = 3428
	UPS        = 1518
	QCOM       = 3381
	NFLX       = 1649
	TWTR       = 1563
	ADBE       = 3383
	TSLA       = 1645
	USDZAR     = 3435
	GBPCAD     = 3403
	DIS        = 1527
	USDMXN     = 3439
	XPDUSD     = 3408
	SBUX       = 1642
	COCOAC     = 4956
	DASH       = 1529
	XLM        = 1508
	RACE       = 3311
	ERBNF      = 3378
	PYPL       = 3382
	SOYBEAN    = 4955
	WFM        = 1606
	NEO        = 3368
	DOT        = 3366
	ATVI       = 3380
	BAYGN      = 3313
	SKK        = 951
	DPZ        = 1598
	EURGBP     = 3419
	PFE        = 1534
	JPY        = 301
	BCH        = 1528
	ABNB       = 3386
	JNJ        = 1503
	ETC        = 1620
	DB         = 1594
	XAGAUD     = 3393
	GER40CASH  = 3365
	PCG        = 1654
	CADCHF     = 3424
	RUB        = 1151
	ZM         = 3320
	EURCHF     = 3426
	DOTUSD     = 3371
	UK100CASH  = 3355
	EU50CASH   = 3359
	CHFJPY     = 3407
	VZ         = 1544
	EURHUF     = 3402
	CHF        = 101
	EURNOK     = 3447
	MU         = 3384
	USDCAD     = 3433
	USDCZK     = 3448
	SPN35CASH  = 3360
	INTC       = 1562
	VOWG_P     = 3319
	EXPE       = 1647
	COFFEEC    = 4959
	XAGEUR     = 3416
	XRP        = 1521
	EUR        = 1
	NZDJPY     = 3415
	AIRF       = 3312
	CZK        = 151
	ETH        = 1451
	XAGUSD     = 3413
	USDHUF     = 3437
	USDILSIS   = 3379
	DBKGN      = 3314
	SGD        = 451
	ILS        = 1051
	SOYBEANC   = 4957
	LVS        = 1609
	XMR        = 1623
	COCOA      = 4951
	AUDJPY     = 3431
	ALVG       = 3323
	NOK        = 351
	UAH        = 1575
	PM         = 1545
	COFFEE     = 4952
	HOG        = 1564
	TROW       = 1512
	XDG        = 1603
	RBAG       = 3325
	TRY        = 651
	CORN       = 4953
	BAT_TST    = 1504
	AUS200CASH = 3361
	MSFT       = 3310
	USDT       = 1610
	CORNC      = 4958
	GE         = 1644
	META       = 3373
	BAC        = 1524
	BABA       = 1542
	WHEATC     = 4960
	CAD        = 51
	DEVX       = 1516
	NZDCHF     = 3442
	WMT        = 1519
	JP225CASH  = 3364
	DKK        = 201
	GOOG       = 3315
	AUDUSD     = 3411
	PG         = 1593
	EURCZK     = 3420
	EURAUD     = 3422
	TWC        = 1557
	ZAR        = 1001
)

const (
	BUY = iota
	SELL
	MARKET = -1
)

func (i *Identity) Login() {

	url := "https://dxtrade." + i.Server + ".com/api/auth/login"
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
	req.Header.Add("authority", "dxtrade."+i.Server+".com")
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
	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	i.Cookies = make(map[string]string)
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
	conn, _, err := dialer.Dial("wss://dxtrade."+i.Server+".com/client/connector?X-Atmosphere-tracking-id=0&X-Atmosphere-Framework=2.3.2-javascript&X-Atmosphere-Transport=websocket&X-Atmosphere-TrackMessageSize=true&Content-Type=text/x-gwt-rpc;%20charset=UTF-8&X-atmo-protocol=true&sessionState=dx-new&guest-mode=false", headers)
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
		fmt.Println(string(message))
		if strings.Contains(string(message), kill) {
			return string(message)
		}
	}
	return ""
}
func (i *Identity) Buy(Quantity, Price, TakeProfit, StopLoss float64, symbol string, instrumentId int) {
	i.ExecuteOrder(BUY, Quantity, Price, TakeProfit, StopLoss, symbol, instrumentId)
}
func (i *Identity) Sell(Quantity, Price, TakeProfit, StopLoss float64, symbol string, instrumentId int) {
	i.ExecuteOrder(SELL, -Quantity, Price, TakeProfit, StopLoss, symbol, instrumentId)
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
	StopLoss    struct {
		FixedOffset           int     `json:"fixedOffset"`
		FixedPrice            float64 `json:"fixedPrice"`
		OrderType             string  `json:"orderType"`
		PriceFixed            bool    `json:"priceFixed"`
		QuantityForProtection float64 `json:"quantityForProtection"`
		Removed               bool    `json:"removed"`
	} `json:"stopLoss"`
	TakeProfit struct {
		FixedOffset           int     `json:"fixedOffset"`
		FixedPrice            float64 `json:"fixedPrice"`
		OrderType             string  `json:"orderType"`
		PriceFixed            bool    `json:"priceFixed"`
		QuantityForProtection float64 `json:"quantityForProtection"`
		Removed               bool    `json:"removed"`
	} `json:"takeProfit"`
}

func (i *Identity) CloseAllPositions() {
	positions := i.GetTransactions()
	for _, position := range positions.Body {
		i.ClosePosition(position.PositionKey.PositionCode, position.Quantity, 0, position.PositionKey.PositionCode, position.PositionKey.InstrumentId)
	}
}
func (i *Identity) ClosePosition(PositionId string, Quantity float64, Price float64, symbol string, instrumentId int) {
	url := "https://dxtrade." + i.Server + ".com/api/positions/close"
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
func (i *Identity) ExecuteOrder(Method int, Quantity, Price, TakeProfit, StopLoss float64, symbol string, instrumentId int) {
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
		executePayload.LimitPrice = 0
		executePayload.OrderType = "MARKET"
	default:
		executePayload.DirectExchange = false
		executePayload.LimitPrice = Price
		executePayload.OrderType = "LIMIT"
	}
	switch Method {
	case BUY:
		executePayload.OrderSide = "BUY"
	case SELL:
		executePayload.OrderSide = "SELL"
	}
	if TakeProfit != 0 {
		executePayload.TakeProfit.FixedOffset = 0
		executePayload.TakeProfit.FixedPrice = TakeProfit
		executePayload.TakeProfit.OrderType = "LIMIT"
		executePayload.TakeProfit.PriceFixed = true
		executePayload.TakeProfit.QuantityForProtection = Quantity
		executePayload.TakeProfit.Removed = false
	}
	if StopLoss != 0 {
		executePayload.StopLoss.FixedOffset = 0
		executePayload.StopLoss.FixedPrice = StopLoss
		executePayload.StopLoss.OrderType = "STOP"
		executePayload.StopLoss.PriceFixed = true
		executePayload.StopLoss.QuantityForProtection = Quantity
		executePayload.StopLoss.Removed = false
	}
	executePayload.Quantity = Quantity
	executePayload.TimeInForce = "GTC"
	//931-08b3a3e1-5e92-4db9-9b32-049777c03e17
	executePayload.RequestId = "gwt-uid-931-" + uuid.New().String()
	url := "https://dxtrade." + i.Server + ".com/api/orders/single"
	method := "POST"

	payload, err := json.Marshal(executePayload)
	if StopLoss == 0 && TakeProfit == 0 {
		payload = []byte(strings.ReplaceAll(string(payload), ",\"stopLoss\":{\"fixedOffset\":0,\"fixedPrice\":0,\"orderType\":\"\",\"priceFixed\":false,\"quantityForProtection\":0,\"removed\":false},\"takeProfit\":{\"fixedOffset\":0,\"fixedPrice\":0,\"orderType\":\"\",\"priceFixed\":false,\"quantityForProtection\":0,\"removed\":false}", ""))
	}
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
	if res.StatusCode != http.StatusOK {
		fmt.Println(req.Header)
		fmt.Println(string(payload), res.Status)
	}
	defer res.Body.Close()
}
func (i *Identity) FetchCSRF() string {
	url := "https://dxtrade." + i.Server + ".com/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "dxtrade."+i.Server+".com")
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
	// GET THIS     <meta id="csrf-token" name="csrf" content="">
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
func (i *Identity) GetOrders() *Order {
	order_str := strings.Split(i.EstablishHandshake("ORDERS"), "|")[1]
	var orders *Order
	err := json.Unmarshal([]byte(order_str), &orders)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return orders
}
func (i *Identity) GetInstruments() *GetInstruments {
	instrument_str := strings.Split(i.EstablishHandshake("Euro vs United States Dollar"), "|")[1]
	var instruments *GetInstruments
	err := json.Unmarshal([]byte(instrument_str), &instruments)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return instruments
}

func (i *Identity) GetAccountMetrix() *AccountData {
	inc_msg := strings.Split(i.EstablishHandshake("ACCOUNT_METRICS"), "|")
	if len(inc_msg) < 2 {
		return nil
	}
	inc_msg2 := inc_msg[1]
	var accountData *AccountData
	err := json.Unmarshal([]byte(inc_msg2), &accountData)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return accountData
}
func (i *Identity) PositionMetrix() *PositionMetrix {
	inc_msg := strings.Split(i.EstablishHandshake("POSITION_METRICS"), "|")
	if len(inc_msg) < 2 {
		return nil
	}
	inc_msg2 := inc_msg[1]
	var positionMetrix *PositionMetrix
	err := json.Unmarshal([]byte(inc_msg2), &positionMetrix)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return positionMetrix
}
func (i *Identity) CancelOrder(OrderId int) bool {
	url := fmt.Sprintf("https://dxtrade.%s.com/api/orders/cancel?accountId=%d&orderChainId=%d", i.Server, i.AccountId, OrderId)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Cookie", fmt.Sprintf("DXTFID=%s; JSESSIONID=%s", i.Cookies["DXTFID"], i.Cookies["JSESSIONID"]))
	req.Header.Add("X-CSRF-Token", i.FetchCSRF())
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false
	}

	return true
}
func (i *Identity) CancelAllOrders() {
	orders := i.GetOrders()
	for _, order := range orders.Body {
		i.CancelOrder(order.OrderId)
	}

}
func (i *Identity) TradeHistory() []TradeHistory {
	url := "https://dxtrade." + i.Server + ".com/api/history?from=1708664400000&to=1714708799999&orderId="
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	req.Header.Add("x-csrf-token", i.FetchCSRF())
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	var tradeHistory []TradeHistory
	err = json.NewDecoder(res.Body).Decode(&tradeHistory)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tradeHistory
}
func (i *Identity) CandleStickProcess(symbol string) (*CandleStickData, error) {
retry:
	var dataStr string
	timeStart := time.Now()
	go func() {
		dataStr = i.EstablishHandshake("BigChartComponentPresenter")
		if strings.Contains(dataStr, "BigChartComponentPresenter") {
			dataStr = strings.Split(dataStr, "|")[1]
		} else {
			dataStr = "err"
		}
	}()
	i.EstablishHandshake("INSTRUMENT_METRICS")
	url := "https://dxtrade.ftmo.com/api/instruments/subscribeInstrumentSymbols"
	method := "PUT"

	payload := strings.NewReader(`{"instruments":["` + symbol + `"]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	req.Header.Add("x-csrf-token", i.FetchCSRF())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	url = "https://dxtrade.ftmo.com/api/charts"
	method = "PUT"

	payload = strings.NewReader(`{"chartIds":[],"requests":[{"aggregationPeriodSeconds":300,"extendedSession":true,"forexPriceField":"bid","id":0,"maxBarsCount":3500,"range":345600,"studySubscription":[],"subtopic":"BigChartComponentPresenter-4","symbol":"` + symbol + `"}]}`)

	client = &http.Client{}
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
	req.Header.Add("x-csrf-token", i.FetchCSRF())
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	for {
		if dataStr != "" {
			if strings.Contains(dataStr, "BigChartComponentPresenter") {
				var candleData *CandleStickData
				err := json.Unmarshal([]byte(dataStr), &candleData)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				return candleData, nil
			} else if dataStr == "err" {
				return nil, errors.New("error fetching data")
			}
			break
		}
		if time.Since(timeStart) > 6500*time.Millisecond {
			log.Println("TIMEOUT")
			goto retry
		}
		time.Sleep(1 * time.Second)
	}
	return nil, errors.New("error fetching data")
}
func (i *Identity) GetCandleStickData(sym string) *CandleStickData {
	var retrievedData *CandleStickData
	for interval := 0; interval < 10; interval++ {
		go func() {
			data, err := i.CandleStickProcess(sym)
			if err != nil {
				fmt.Println(err)
			} else {
				retrievedData = data
			}

		}()
	}
	for {
		if retrievedData != nil {
			return retrievedData
		}
		time.Sleep(100 * time.Millisecond)
	}
	/*
		Candlestick data is in a slice under the data key ranging from the oldest to the newest data available by candle
		I'm not sure what the timeframe is
	*/

}

// Identical to the one above but with a different return type

// Couldnt get it to work, some weird api update but if you want to give it a shot you need to be able to find the refOrderChainId
//
//	func (i *Identity) ModifyOrder(orderId int, newLimit, newTp, newSl float64) bool {
//		type Payload struct {
//			DirectExchange bool `json:"directExchange"`
//			Legs           []struct {
//				InstrumentId   int    `json:"instrumentId"`
//				PositionEffect string `json:"positionEffect"`
//				RatioQuantity  int    `json:"ratioQuantity"`
//				Symbol         string `json:"symbol"`
//			} `json:"legs"`
//			LimitPrice      float64 `json:"limitPrice"`
//			OrderSide       string  `json:"orderSide"`
//			OrderType       string  `json:"orderType"`
//			Quantity        float64 `json:"quantity"`
//			RefOrderChainId string  `json:"refOrderChainId"`
//			RequestId       string  `json:"requestId"`
//			StopLoss        struct {
//				FixedOffset           int     `json:"fixedOffset"`
//				FixedPrice            float64 `json:"fixedPrice"`
//				OrderType             string  `json:"orderType"`
//				PriceFixed            bool    `json:"priceFixed"`
//				QuantityForProtection float64 `json:"quantityForProtection"`
//				RefOrderChainId       string  `json:"refOrderChainId"`
//				Removed               bool    `json:"removed"`
//			} `json:"stopLoss"`
//			TakeProfit struct {
//				FixedOffset           int     `json:"fixedOffset"`
//				FixedPrice            float64 `json:"fixedPrice"`
//				OrderType             string  `json:"orderType"`
//				PriceFixed            bool    `json:"priceFixed"`
//				QuantityForProtection float64 `json:"quantityForProtection"`
//				RefOrderChainId       string  `json:"refOrderChainId"`
//				Removed               bool    `json:"removed"`
//			} `json:"takeProfit"`
//			TimeInForce string `json:"timeInForce"`
//		}
//		var x int
//		for i, order := range i.GetOrders().Body {
//			if order.OrderId == orderId {
//				x = i
//			}
//		}
//		myOrder := i.GetOrders().Body[x]
//		var payload Payload
//		payload.DirectExchange = false
//		payload.Legs = make([]struct {
//			InstrumentId   int    `json:"instrumentId"`
//			PositionEffect string `json:"positionEffect"`
//			RatioQuantity  int    `json:"ratioQuantity"`
//			Symbol         string `json:"symbol"`
//		}, 1)
//		fmt.Println(myOrder.Legs[0].PositionCode)
//
//		payload.Legs[0].InstrumentId = myOrder.Legs[0].InstrumentId
//		payload.Legs[0].PositionEffect = "OPENING"
//		payload.Legs[0].RatioQuantity = 1
//		var symbol string
//		for key, value := range symbols {
//			if value == myOrder.Legs[0].InstrumentId {
//				symbol = key
//			}
//		}
//		payload.Legs[0].Symbol = symbol
//		switch newLimit {
//		case -1:
//			payload.LimitPrice = 0
//			payload.OrderType = "MARKET"
//		default:
//			payload.LimitPrice = newLimit
//			payload.OrderType = "LIMIT"
//		}
//		var orderSide string
//		if myOrder.Quantity > 0 {
//			orderSide = "BUY"
//		} else {
//			orderSide = "SELL"
//		}
//		payload.OrderSide = orderSide
//		payload.Quantity = myOrder.Quantity
//		payload.RefOrderChainId = myOrder.OrderChainId
//		fmt.Println("ORDER IDS", myOrder.OrderChainId, myOrder.OrderId, myOrder.ParentOrderId, myOrder.RequestId, myOrder.ThenOrdersIds)
//		if newTp != 0 {
//			payload.TakeProfit.FixedOffset = 0
//			payload.TakeProfit.FixedPrice = newTp
//			payload.TakeProfit.OrderType = "LIMIT"
//			payload.TakeProfit.PriceFixed = true
//			payload.TakeProfit.QuantityForProtection = myOrder.Quantity
//			payload.TakeProfit.RefOrderChainId = "3870852:20601"
//			payload.TakeProfit.Removed = false
//		}
//		if newSl != 0 {
//			payload.StopLoss.FixedOffset = 0
//			payload.StopLoss.FixedPrice = newSl
//			payload.StopLoss.OrderType = "STOP"
//			payload.StopLoss.PriceFixed = true
//			payload.StopLoss.QuantityForProtection = myOrder.Quantity
//			payload.StopLoss.RefOrderChainId = "3870852:20602"
//			payload.StopLoss.Removed = false
//		}
//		payload.TimeInForce = "GTC"
//		payload.RequestId = "gwt-uid-2553-" + uuid.New().String()
//		url := "https://dxtrade." + i.Server + ".com/api/orders/single"
//		var payloadJson []byte
//		payloadJson, err := json.Marshal(payload)
//		fmt.Println(string(payloadJson))
//		if err != nil {
//			fmt.Println(err)
//			return false
//		}
//		client := &http.Client{}
//		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJson))
//
//		if err != nil {
//			fmt.Println(err)
//			return false
//		}
//		req.Header.Add("content-type", "application/json; charset=UTF-8")
//		req.Header.Add("cookie", "DXTFID="+i.Cookies["DXTFID"]+"; JSESSIONID="+i.Cookies["JSESSIONID"])
//		req.Header.Add("x-csrf-token", i.FetchCSRF())
//		req.Header.Add("x-requested-with", "XMLHttpRequest")
//
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Println(err)
//			return false
//		}
//		defer res.Body.Close()
//
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			fmt.Println(err)
//			return false
//		}
//		fmt.Println(string(body))
//		return true
//	}

type PositionMetrix struct {
	AccountId string `json:"accountId"`
	Body      []struct {
		Uid              string  `json:"uid"`
		AccountId        string  `json:"accountId"`
		Margin           float64 `json:"margin"`
		PlOpen           float64 `json:"plOpen"`
		PlClosed         int     `json:"plClosed"`
		TotalCommissions float64 `json:"totalCommissions"`
		TotalFinancing   float64 `json:"totalFinancing"`
		PlRate           float64 `json:"plRate"`
	} `json:"body"`
	Type string `json:"type"`
}
type AccountData struct {
	AccountId string `json:"accountId"`
	Body      struct {
		AccountId  string `json:"accountId"`
		AllMetrics struct {
			AvailableFunds   float64     `json:"availableFunds"`
			MarginCallLevel  interface{} `json:"marginCallLevel"`
			RiskLevel        float64     `json:"riskLevel"`
			OpenPl           float64     `json:"openPl"`
			CashBalance      float64     `json:"cashBalance"`
			Equity           float64     `json:"equity"`
			ConversionRate   int         `json:"conversionRate"`
			ReverseRiskLevel interface{} `json:"reverseRiskLevel"`
			InitialMargin    float64     `json:"initialMargin"`
		} `json:"allMetrics"`
	} `json:"body"`
	Type string `json:"type"`
}
type Order struct {
	AccountId string `json:"accountId"`
	Body      []struct {
		OrderId                int         `json:"orderId"`
		AccountId              string      `json:"accountId"`
		OrderChainId           string      `json:"orderChainId"`
		OcoGroupCode           interface{} `json:"ocoGroupCode"`
		BracketGroupCode       interface{} `json:"bracketGroupCode"`
		Status                 string      `json:"status"`
		StatusDescription      string      `json:"statusDescription"`
		CreatedTime            int64       `json:"createdTime"`
		ModifiedTime           int64       `json:"modifiedTime"`
		Quantity               float64     `json:"quantity"`
		RemainingQuantity      float64     `json:"remainingQuantity"`
		FilledQuantity         int         `json:"filledQuantity"`
		Type                   string      `json:"type"`
		LimitPrice             int         `json:"limitPrice"`
		AveragePrice           string      `json:"averagePrice"`
		StopPrice              string      `json:"stopPrice"`
		StopLimitPrice         int         `json:"stopLimitPrice"`
		TriggerPrice           string      `json:"triggerPrice"`
		Liquidation            bool        `json:"liquidation"`
		TrailPrice             string      `json:"trailPrice"`
		Attached               bool        `json:"attached"`
		StopPriceOffset        string      `json:"stopPriceOffset"`
		StopPriceOffsetPercent int         `json:"stopPriceOffsetPercent"`
		RequestId              interface{} `json:"requestId"`
		TimeInForce            string      `json:"timeInForce"`
		FillPrice              string      `json:"fillPrice"`
		ExpireAt               int         `json:"expireAt"`
		ClosedPL               string      `json:"closedPL"`
		ParentOrderId          interface{} `json:"parentOrderId"`
		ThenOrdersIds          interface{} `json:"thenOrdersIds"`
		OrderRole              string      `json:"orderRole"`
		OrderGroupCode         string      `json:"orderGroupCode"`
		TakeProfit             interface{} `json:"takeProfit"`
		StopLoss               interface{} `json:"stopLoss"`
		Exchange               interface{} `json:"exchange"`
		LastFillTime           interface{} `json:"lastFillTime"`
		StopPriceTriggerTime   interface{} `json:"stopPriceTriggerTime"`
		HasTriggeredStop       bool        `json:"hasTriggeredStop"`
		CommissionFee          interface{} `json:"commissionFee"`
		MarginRate             float64     `json:"marginRate"`
		AlertExpression        interface{} `json:"alertExpression"`
		Reason                 interface{} `json:"reason"`
		AdditionalParameters   struct {
		} `json:"additionalParameters"`
		Legs []struct {
			PositionEffect string      `json:"positionEffect"`
			RatioQuantity  int         `json:"ratioQuantity"`
			PositionCode   string      `json:"positionCode"`
			InstrumentId   int         `json:"instrumentId"`
			Symbol         interface{} `json:"symbol"`
		} `json:"legs"`
		ActualWithinTradingDay   bool        `json:"actualWithinTradingDay"`
		Route                    interface{} `json:"route"`
		OpeningPositionCostBasis int         `json:"openingPositionCostBasis"`
	} `json:"body"`
	Type string `json:"type"`
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
type GetInstruments struct {
	AccountID any `json:"accountId"`
	Body      []struct {
		ID                int     `json:"id"`
		Symbol            string  `json:"symbol"`
		Description       string  `json:"description"`
		Type              string  `json:"type"`
		Subtype           string  `json:"subtype"`
		Currency          string  `json:"currency"`
		CurrencyPrecision int     `json:"currencyPrecision"`
		Precision         int     `json:"precision"`
		PipsSize          float64 `json:"pipsSize"`
		QuantityIncrement float64 `json:"quantityIncrement"`
		QuantityPrecision int     `json:"quantityPrecision"`
		PriceIncrement    float64 `json:"priceIncrement"`
		Version           int     `json:"version"`
		PriceIncrementsTO struct {
			PriceIncrements []float64 `json:"priceIncrements"`
			PricePrecisions []int     `json:"pricePrecisions"`
			BondFraction    bool      `json:"bondFraction"`
		} `json:"priceIncrementsTO"`
		LotSize              int    `json:"lotSize"`
		BaseCurrency         string `json:"baseCurrency"`
		LotName              any    `json:"lotName"`
		Multiplier           int    `json:"multiplier"`
		Open                 bool   `json:"open"`
		Expiration           any    `json:"expiration"`
		FirstNoticeDate      any    `json:"firstNoticeDate"`
		InitialMargin        string `json:"initialMargin"`
		MaintenanceMargin    string `json:"maintenanceMargin"`
		LastTradeDate        any    `json:"lastTradeDate"`
		Underlying           any    `json:"underlying"`
		Mmy                  any    `json:"mmy"`
		OptionParametersTO   any    `json:"optionParametersTO"`
		UnitName             any    `json:"unitName"`
		AdditionalFields     any    `json:"additionalFields"`
		AdditionalObject     any    `json:"additionalObject"`
		CurrencyParametersTO any    `json:"currencyParametersTO"`
		TradingHours         string `json:"tradingHours"`
	} `json:"body"`
	Type string `json:"type"`
}

type Body []struct {
	OrderId                int         `json:"orderId"`
	AccountId              string      `json:"accountId"`
	OrderChainId           string      `json:"orderChainId"`
	OcoGroupCode           interface{} `json:"ocoGroupCode"`
	BracketGroupCode       interface{} `json:"bracketGroupCode"`
	Status                 string      `json:"status"`
	StatusDescription      string      `json:"statusDescription"`
	CreatedTime            int64       `json:"createdTime"`
	ModifiedTime           int64       `json:"modifiedTime"`
	Quantity               float64     `json:"quantity"`
	RemainingQuantity      float64     `json:"remainingQuantity"`
	FilledQuantity         int         `json:"filledQuantity"`
	Type                   string      `json:"type"`
	LimitPrice             int         `json:"limitPrice"`
	AveragePrice           string      `json:"averagePrice"`
	StopPrice              string      `json:"stopPrice"`
	StopLimitPrice         int         `json:"stopLimitPrice"`
	TriggerPrice           string      `json:"triggerPrice"`
	Liquidation            bool        `json:"liquidation"`
	TrailPrice             string      `json:"trailPrice"`
	Attached               bool        `json:"attached"`
	StopPriceOffset        string      `json:"stopPriceOffset"`
	StopPriceOffsetPercent int         `json:"stopPriceOffsetPercent"`
	RequestId              interface{} `json:"requestId"`
	TimeInForce            string      `json:"timeInForce"`
	FillPrice              string      `json:"fillPrice"`
	ExpireAt               int         `json:"expireAt"`
	ClosedPL               string      `json:"closedPL"`
	ParentOrderId          interface{} `json:"parentOrderId"`
	ThenOrdersIds          interface{} `json:"thenOrdersIds"`
	OrderRole              string      `json:"orderRole"`
	OrderGroupCode         string      `json:"orderGroupCode"`
	TakeProfit             interface{} `json:"takeProfit"`
	StopLoss               interface{} `json:"stopLoss"`
	Exchange               interface{} `json:"exchange"`
	LastFillTime           interface{} `json:"lastFillTime"`
	StopPriceTriggerTime   interface{} `json:"stopPriceTriggerTime"`
	HasTriggeredStop       bool        `json:"hasTriggeredStop"`
	CommissionFee          interface{} `json:"commissionFee"`
	MarginRate             float64     `json:"marginRate"`
	AlertExpression        interface{} `json:"alertExpression"`
	Reason                 interface{} `json:"reason"`
	AdditionalParameters   struct{}    `json:"additionalParameters"`
	Legs                   []struct {
		PositionEffect string      `json:"positionEffect"`
		RatioQuantity  int         `json:"ratioQuantity"`
		PositionCode   string      `json:"positionCode"`
		InstrumentId   int         `json:"instrumentId"`
		Symbol         interface{} `json:"symbol"`
	} `json:"legs"`
	ActualWithinTradingDay   bool        `json:"actualWithinTradingDay"`
	Route                    interface{} `json:"route"`
	OpeningPositionCostBasis int         `json:"openingPositionCostBasis"`
}
type CandleStickData struct {
	AccountID any `json:"accountId"`
	Body      struct {
		Subtopic               string `json:"subtopic"`
		HistoryData            bool   `json:"historyData"`
		HistoryDataSnapshotEnd bool   `json:"historyDataSnapshotEnd"`
		Data                   []struct {
			Timestamp     int64   `json:"timestamp"`
			Open          float64 `json:"open"`
			High          float64 `json:"high"`
			Low           float64 `json:"low"`
			Close         float64 `json:"close"`
			ImpVolatility string  `json:"impVolatility"`
			Vwap          float64 `json:"vwap"`
			Volume        int     `json:"volume"`
			Expansion     bool    `json:"expansion"`
			Visible       bool    `json:"visible"`
			StudyValues   any     `json:"studyValues"`
			Time          int64   `json:"time"`
		} `json:"data"`
		RequestID int `json:"requestId"`
	} `json:"body"`
	Type string `json:"type"`
}
type TradeHistory struct {
	Id                string      `json:"id"`
	SequenceNumber    int         `json:"sequenceNumber"`
	OrderChainId      string      `json:"orderChainId"`
	AccountId         string      `json:"accountId"`
	LastStatusChange  int64       `json:"lastStatusChange"`
	Symbol            string      `json:"symbol"`
	Status            string      `json:"status"`
	Quantity          float64     `json:"quantity"`
	RemainingQuantity float64     `json:"remainingQuantity"`
	FilledQuantity    interface{} `json:"filledQuantity"`
	Type              string      `json:"type"`
	TriggerPrice      string      `json:"triggerPrice"`
	FillPrice         interface{} `json:"fillPrice"`
	Price             string      `json:"price"`
	AccountCode       string      `json:"accountCode"`
	TakeProfitPrice   string      `json:"takeProfitPrice"`
	ExpireAt          interface{} `json:"expireAt"`
	TimeInForce       string      `json:"timeInForce"`
	StopLossPrice     string      `json:"stopLossPrice"`
	Commission        struct {
	} `json:"commission"`
	Side         string      `json:"side"`
	RejectReason interface{} `json:"rejectReason"`
	Instrument   struct {
		Id                int     `json:"id"`
		Symbol            string  `json:"symbol"`
		Description       string  `json:"description"`
		Type              string  `json:"type"`
		Subtype           string  `json:"subtype"`
		Currency          string  `json:"currency"`
		CurrencyPrecision int     `json:"currencyPrecision"`
		Precision         int     `json:"precision"`
		PipsSize          int     `json:"pipsSize"`
		QuantityIncrement float64 `json:"quantityIncrement"`
		QuantityPrecision int     `json:"quantityPrecision"`
		PriceIncrement    float64 `json:"priceIncrement"`
		Version           int     `json:"version"`
		PriceIncrementsTO struct {
			PriceIncrements []float64 `json:"priceIncrements"`
			PricePrecisions []int     `json:"pricePrecisions"`
			BondFraction    bool      `json:"bondFraction"`
		} `json:"priceIncrementsTO"`
		LotSize              int         `json:"lotSize"`
		BaseCurrency         interface{} `json:"baseCurrency"`
		LotName              interface{} `json:"lotName"`
		Multiplier           int         `json:"multiplier"`
		Open                 bool        `json:"open"`
		Expiration           interface{} `json:"expiration"`
		FirstNoticeDate      interface{} `json:"firstNoticeDate"`
		InitialMargin        string      `json:"initialMargin"`
		MaintenanceMargin    string      `json:"maintenanceMargin"`
		LastTradeDate        interface{} `json:"lastTradeDate"`
		Underlying           interface{} `json:"underlying"`
		Mmy                  interface{} `json:"mmy"`
		OptionParametersTO   interface{} `json:"optionParametersTO"`
		UnitName             interface{} `json:"unitName"`
		AdditionalFields     interface{} `json:"additionalFields"`
		AdditionalObject     interface{} `json:"additionalObject"`
		CurrencyParametersTO interface{} `json:"currencyParametersTO"`
		TradingHours         string      `json:"tradingHours"`
	} `json:"instrument"`
	ExecutionKeyId       int         `json:"executionKeyId"`
	NetPl                interface{} `json:"netPl"`
	GrossPl              interface{} `json:"grossPl"`
	AdditionalParameters interface{} `json:"additionalParameters"`
}
