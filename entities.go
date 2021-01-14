package main

// NewUserAccounts constructor
func NewUserAccounts() *UserAccounts { return new(UserAccounts) }

// NewСurrencys constructor
func NewСurrencys() *Сurrencys { return new(Сurrencys) }

// NewExchanges constructor
func NewExchanges() *Exchanges { return new(Exchanges) }

// NewChangesV1 constructor
func NewChangesV1() *ChangesV1 { return new(ChangesV1) }

// NewChangesV2 constructor
func NewChangesV2() *ChangesV2 { return new(ChangesV2) }

// NewChangesV3 constructor
func NewChangesV3() *ChangesV3 { return new(ChangesV3) }

// NewSymbolsV1 constructor
func NewSymbolsV1() *SymbolsV1 { return new(SymbolsV1) }

// NewSymbolsV2 constructor
func NewSymbolsV2() *SymbolsV2 { return new(SymbolsV2) }

// NewSymbolsV3 constructor
func NewSymbolsV3() *SymbolsV3 { return new(SymbolsV3) }

// NewSymbolV1 constructor
func NewSymbolV1() *SymbolV1 { return new(SymbolV1) }

// NewSymbolV2 constructor
func NewSymbolV2() *SymbolV2 { return new(SymbolV2) }

// NewSymbolV3 constructor
func NewSymbolV3() *SymbolV3 { return new(SymbolV3) }

// NewSchedule constructor
func NewSchedule() *Schedule { return new(Schedule) }

// NewTypes constructor
func NewTypes() *Types { return new(Types) }

// NewSymbolSpecification constructor
func NewSymbolSpecification() *SymbolSpecification { return new(SymbolSpecification) }

// NewGroups constructor
func NewGroups() *Groups { return new(Groups) }

// NewQuote constructor
func NewQuote() *Quote { return new(Quote) }

// NewOHLCTrades constructor
func NewOHLCTrades() *OHLCTrades { return new(OHLCTrades) }

// NewOHLCQuotes constructor
func NewOHLCQuotes() *OHLCQuotes { return new(OHLCQuotes) }

// NewSummary constructor
func NewSummary() *Summary { return new(Summary) }

// NewTransactionsV1 constructor
func NewTransactionsV1() *TransactionsV1 { return new(TransactionsV1) }

// NewTransactionsV2 constructor
func NewTransactionsV2() *TransactionsV2 { return new(TransactionsV2) }

// NewTransactionsV3 constructor
func NewTransactionsV3() *TransactionsV3 { return new(TransactionsV3) }

// NewOrderV1 constructor
func NewOrderV1() *OrderV1 { return new(OrderV1) }

// NewOrderV2 constructor
func NewOrderV2() *OrderV2 { return new(OrderV2) }

// NewOrderV3 constructor
func NewOrderV3() *OrderV3 { return new(OrderV3) }

// NewOrdersV1 constructor
func NewOrdersV1() *OrdersV1 { return new(OrdersV1) }

// NewOrdersV2 constructor
func NewOrdersV2() *OrdersV2 { return new(OrdersV2) }

// NewOrdersV3 constructor
func NewOrdersV3() *OrdersV3 { return new(OrdersV3) }

// UserAccount model
type UserAccount struct {
	Status    string `json:"status"`
	AccountID string `json:"accountId"`
}

// UserAccounts model
type UserAccounts []UserAccount

// Currency ..
type Currency string

// Сurrencys model
type Сurrencys struct {
	Currencies []Currency `json:"currencies"`
}

// Exchange model
type Exchange struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// Exchanges model
type Exchanges []Exchange

// ChangeV1 model
type ChangeV1 struct {
	BasePrice   string `json:"basePrice"`
	DailyChange string `json:"dailyChange"`
	SymbolID    string `json:"symbolId"`
}

// ChangesV1 model
type ChangesV1 []ChangeV1

// ChangeV2 model
type ChangeV2 ChangeV1

// ChangesV2 model
type ChangesV2 []ChangeV2

// ChangeV3 model
type ChangeV3 struct {
	LastSessionClosePrice string `json:"lastSessionClosePrice"`
	DailyChange           string `json:"dailyChange"`
	SymbolID              string `json:"symbolId"`
}

// ChangesV3 model
type ChangesV3 []ChangeV3

// SymbolOptionData optionals params for symbol
type SymbolOptionData struct {
	OptionIDGroup string  `json:"optionGroupId"`
	Right         string  `json:"right"`
	StrikePrice   float64 `json:"strikePrice"`
}

// SymbolOptionDataV2 ...
type SymbolOptionDataV2 struct {
	OptionIDGroup string `json:"optionGroupId"`
	Right         string `json:"right"`
	StrikePrice   string `json:"strikePrice"`
}

// SymbolOptionDataV3 optionals params for symbol (api v3)
type SymbolOptionDataV3 struct {
	OptionIDGroup string `json:"optionGroupId"`
	OptionRight   string `json:"optionRight"`
	StrikePrice   string `json:"strikePrice"`
}

// SymbolV1 model
type SymbolV1 struct {
	SymbolOptionData `json:"optionData"`
	I18n             string  `json:"i18n"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Country          string  `json:"country"`
	Exchange         string  `json:"exchange"`
	ID               string  `json:"id"`
	Currency         string  `json:"currency"`
	Mpi              float64 `json:"mpi"`
	Type             string  `json:"type"`
	Ticker           string  `json:"ticker"`
	Expiration       float64 `json:"expiration"`
	Group            string  `json:"group"`
}

// SymbolsV1 model
type SymbolsV1 []SymbolV1

// SymbolV2 model
type SymbolV2 struct {
	SymbolOptionDataV2 `json:"optionData"`
	I18n               string  `json:"i18n"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Country            string  `json:"country"`
	Exchange           string  `json:"exchange"`
	ID                 string  `json:"id"`
	Currency           string  `json:"currency"`
	Mpi                string  `json:"mpi"`
	Type               string  `json:"type"`
	Ticker             string  `json:"ticker"`
	Expiration         float64 `json:"expiration"`
	Group              string  `json:"group"`
}

// SymbolsV2 model
type SymbolsV2 []SymbolV2

// SymbolV3 model
type SymbolV3 struct {
	SymbolOptionDataV3 `json:"optionData"`
	I18n               string  `json:"i18n"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Country            string  `json:"country"`
	Exchange           string  `json:"exchange"`
	MinPriceIncrement  string  `json:"minPriceIncrement"`
	Currency           string  `json:"currency"`
	SymbolType         string  `json:"symbolType"`
	Ticker             string  `json:"ticker"`
	Expiration         float64 `json:"expiration"`
	Group              string  `json:"group"`
}

// SymbolsV3 model
type SymbolsV3 []SymbolV3

// Type concrete knows type
type Type struct {
	ID string `json:"id"`
}

// Types known type
type Types []Type

// Period - part of schedule
type Period struct {
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

// Schedule model
type Schedule struct {
	Period []Period

	Name       string `json:"name"`
	OrderTypes string `json:"orderTypes"`
}

// SymbolSpecification model
type SymbolSpecification struct {
	Leverage           string `json:"leverage"`
	ContractMultiplier string `json:"contractMultiplier"`
	PriceUnit          string `json:"priceUnit"`
	Units              string `json:"unitd"`
	LotSize            string `json:"lotSize"`
}

// Group model
type Group struct {
	Group    string   `json:"group"`
	Name     string   `json:"name"`
	Exchange string   `json:"exchange"`
	Types    []string `json:"types"`
}

// Groups model
type Groups []Group

// Quote model
type Quote struct {
	Timestamp int    `json:"timestamp"`
	SymbolID  int    `json:"symbolId"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
}

// OHLCQuote model
type OHLCQuote struct {
	Open      int `json:"open_"`
	Low       int `json:"low"`
	High      int `json:"high"`
	Close     int `json:"close"`
	Timestamp int `json:"timestamp"`
}

// OHLCQuotes models
type OHLCQuotes []OHLCQuote

// OHLCTrade model
type OHLCTrade struct {
	Open      int    `json:"open_"`
	Low       int    `json:"low"`
	High      int    `json:"high"`
	Close     int    `json:"close"`
	Timestamp int    `json:"timestamp"`
	Volume    string `json:"volume"`
}

// OHLCTrades model
type OHLCTrades []OHLCTrade

// OrderParameters model
type OrderParameters struct {
	Side           string `json:"side"`
	Duration       string `json:"duration"`
	Quantity       string `json:"quantity"`
	Instrument     string `json:"instrument"`
	OrderType      string `json:"orderType"`
	OcoGroup       string `json:"ocoGroup"`
	IfDoneParentID string `json:"ifDoneParentId"`
	LimitPrice     string `json:"limitPrice"`
	StopPrice      string `json:"stopPrice"`
	PriceDistance  string `json:"priceDistance"`
	PartQuantity   string `json:"partQuantity"`
	PlaceInterval  string `json:"placeInterval"`
}

// OrderFill model
type OrderFill struct {
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Time     string `json:"timestamp"`
	Position int    `json:"position"`
}

// OrderState model
type OrderState struct {
	Fills []OrderFill `json:"fills"`

	Status     string `json:"status"`
	LastUpdate string `json:"lastUpdate"`
}

// OrderV1 model
type OrderV1 struct {
	OrderState      OrderState      `json:"orderState"`
	OrderParameters OrderParameters `json:"orderParameters"`

	PlaceTime             string `json:"placeTime"`
	OrderID               string `json:"orderId"`
	ClientTag             string `json:"clientTag"`
	CurrentModificationID string `json:"currentModificationId"`
	ExanteAccount         string `json:"exanteAccount"`
	Username              string `json:"username"`
}

// OrdersV1 model
type OrdersV1 []OrderV1

// OrderV2 model
type OrderV2 struct {
	OrderState      OrderState      `json:"orderState"`
	OrderParameters OrderParameters `json:"orderParameters"`

	OrderID               string `json:"orderId"`
	PlaceTime             string `json:"placeTime"`
	AccountID             string `json:"accountId"`
	ClientTag             string `json:"clientTag"`
	CurrentModificationID string `json:"currentModificationId"`
	ExanteAccount         string `json:"exanteAccount"`
	Username              string `json:"username"`
	ID                    string `json:"id"`
}

// OrdersV2 model
type OrdersV2 []OrderV2

// OrderV3 model
type OrderV3 struct {
	OrderState      OrderState      `json:"orderState"`
	OrderParameters OrderParameters `json:"orderParameters"`

	OrderID               string `json:"orderId"`
	PlaceTime             string `json:"placeTime"`
	AccountID             string `json:"accountId"`
	ClientTag             string `json:"clientTag"`
	CurrentModificationID string `json:"currentModificationId"`
	ExanteAccount         string `json:"exanteAccount"`
	Username              string `json:"username"`
}

// OrdersV3 model
type OrdersV3 []OrderV3

// CurrencyPos model
type CurrencyPos struct {
	Code           string `json:"code"`
	Value          string `json:"value"`
	ConvertedValue string `json:"convertedValue"`
}

// Position model
type Position struct {
	ID             string `json:"id"`
	SymbolType     string `json:"symbolType"`
	Currency       string `json:"currency"`
	Price          string `json:"price"`
	AveragePrice   string `json:"averagePrice"`
	Quantity       string `json:"quantity"`
	Value          string `json:"value"`
	ConvertedValue string `json:"convertedValue"`
	PNL            string `json:"pnl"`
	ConvertedPNL   string `json:"convertedPnl"`
}

// Summary model
type Summary struct {
	Currencies []CurrencyPos `json:"currencies"`
	Positions  []Position    `json:"positions"`

	Account            string `json:"account"`
	SessionDate        string `json:"sessionDate"`
	FreeMoney          string `json:"freeMoney"`
	Timestamp          int    `json:"timestamp"`
	NetAssetValue      string `json:"netAssetValue"`
	MarginUtilization  string `json:"marginUtilization"`
	MoneyUsedForMargin string `json:"moneyUsedForMargin"`
	Currency           string `json:"currency"`
}

// TransactionV1 model
type TransactionV1 struct {
	OperationType string  `json:"operationType"`
	ID            string  `json:"id"`
	Asset         string  `json:"asset"`
	When          int     `json:"when"`
	Sum           float64 `json:"sum"`
	SymbolID      string  `json:"symbolId"`
	AccountID     string  `json:"accountId"`
}

// TransactionsV1 model
type TransactionsV1 []TransactionV1

//TransactionV2 model
type TransactionV2 TransactionV1

// TransactionsV2 model
type TransactionsV2 []TransactionV2

// TransactionV3 model
type TransactionV3 struct {
	OperationType string  `json:"operationType"`
	ID            string  `json:"id"`
	Asset         string  `json:"asset"`
	SymbolID      string  `json:"symbolId"`
	AccountID     string  `json:"accountId"`
	Timestamp     int     `json:"timestamp"`
	Sum           float64 `json:"sum"`
}

// TransactionsV3 model
type TransactionsV3 []TransactionV3

// OrderSentTypeV1 model
type OrderSentTypeV1 struct {
	Account        string `json:"account"`
	Instrument     string `json:"instrument"`
	Side           string `json:"side"`
	Quantity       string `json:"quantity"`
	Duration       string `json:"duration"`
	ClientTag      string `json:"clientTag"`
	OcoGroup       string `json:"ocoGroup"`
	IfDoneParentID string `json:"ifDoneParentId"`
	OrderType      string `json:"orderType"`
}

// OrderSentTypeV2 model
type OrderSentTypeV2 struct {
	AccountID      string `json:"accountId"`
	Instrument     string `json:"instrument"`
	Side           string `json:"side"`
	Quantity       string `json:"quantity"`
	Duration       string `json:"duration"`
	ClientTag      string `json:"clientTag"`
	OcoGroup       string `json:"ocoGroup"`
	IfDoneParentID string `json:"ifDoneParentId"`
	OrderType      string `json:"orderType"`
	TakeProfit     string `json:"takeProfit"`
	StopLoss       string `json:"stopLoss"`
}

// OrderSentTypeV3 model
type OrderSentTypeV3 struct {
	AccountID      string `json:"accountId"`
	Instrument     string `json:"instrument"`
	Side           string `json:"side"`
	Quantity       string `json:"quantity"`
	Duration       string `json:"duration"`
	ClientTag      string `json:"clientTag"`
	OcoGroup       string `json:"ocoGroup"`
	IfDoneParentID string `json:"ifDoneParentId"`
	OrderType      string `json:"orderType"`
	TakeProfit     string `json:"takeProfit"`
	StopLoss       string `json:"stopLoss"`
	SymbolID       string `json:"symbolId"`
}
