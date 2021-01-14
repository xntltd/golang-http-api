package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// DataTypes for aggregations
	quotesDataType string = "quotes"
	tradesDataType string = "trades"

	// Order actions
	cancelOrder  string = "cancel"
	replaceOrder string = "replace"

	// URL actions
	accountsAction         string = "accounts"
	crossratesAction       string = "crossrates"
	exchangesAction        string = "exchanges"
	changeAction           string = "change"
	typesAction            string = "types"
	transactionsAction     string = "transactions"
	ordersAction           string = "orders"
	symbolsAction          string = "symbols"
	groupsAction           string = "groups"
	activeOrdersAction     string = "orders/active"
	scheduleAction         string = "schedule"
	specificationAction    string = "specification"
	nearestAction          string = "nearest"
	feedAction             string = "feed"
	lastAction             string = "last"
	ohlcAction             string = "ohlc"
	ticksAction            string = "ticks"
	summaryAction          string = "summary"
	ordersStreamAction     string = "stream/orders"
	execOrdersStreamAction string = "stream/trades"
	tradeStreamAction      string = "feed/trades"

	// Error messages
	errInvalidAggregationMessage   string = "INVALID aggType.Value MUST BE ONE OF trades/quotes"
	errUndefinedAPIVersionMessage  string = "UNDEFINED API VERSION"
	errUndefinedCategoryAPIMessage string = "UNDEFINED CATEGORY API"

	// Max Response size
	responseLimit       int = 1000
	defaultResponseSize int = 10
)

// NewAPI constructor for lib
func NewAPI(
	baseAPIURL, version,
	ApplicationID, ClientID, SharedKey string,
	JwtTTL int,
	BasicAuthUsername, BasicAuthPassword string) HTTPApi {

	h := HTTPApi{
		Auth: Auth{
			JWT: JWTAuthMethod{
				ApplicationID: ApplicationID,
				ClientID:      ClientID,
				SharedKey:     SharedKey,
				JwtTTL:        JwtTTL,
			},
			Basic: BasicAuthMethod{
				Username: BasicAuthUsername,
				Password: BasicAuthPassword,
			},
		},
		httpClient: &http.Client{
			Transport: &libTransport{
				underlyingTransport: http.DefaultTransport,
			},
		},
		baseAPIURL: baseAPIURL,
		version:    version,
	}
	return h
}

var emptyPostPayload = bytes.NewReader([]byte{})

func intToString(value int) string                 { return strconv.Itoa(value) }
func stringToUpperCase(value string) string        { return strings.ToUpper(value) }
func boolToString(value bool) string               { return strconv.FormatBool(value) }
func strungSliceToString(value []string) string    { return strings.Join(value, commaSeparator) }
func typeToJSON(value interface{}) ([]byte, error) { return json.Marshal(&value) }
func joinWithCommaSeparator(s []string) string     { return strings.Join(uncodeString(s), commaSeparator) }
func joinWithSlashSeparator(s []string) string     { return strings.Join(uncodeString(s), slashSeparator) }
func uncodeString(s []string) (o []string) {
	for _, i := range s {
		o = append(o, url.PathEscape(i))
	}
	return o
}

// QueryStringParams like ?foo=bar/
type QueryStringParams map[string]string

// StartStopOptionalPayload optional params
type StartStopOptionalPayload struct{ Start, Stop int }

// OffsetLimitPayload offset and limit optional params
type OffsetLimitPayload struct{ Offset, Limit int }

// DatetimePayload datetime optional field for requests
type DatetimePayload struct {
	Datetime string `json:"datetime"`
}

// DatetimeRangePayload optional query params for requsts
type DatetimeRangePayload struct{ From, To int }

func (d DatetimeRangePayload) getFromParam() string {
	if d.From > 0 {
		return intToString(d.From)
	}
	return emptyString
}

func (d DatetimeRangePayload) getToParam() string {
	if d.To > 0 {
		return intToString(d.To)
	}
	return emptyString
}

// ResponseLimit max size of response
type ResponseLimit struct{ Limit int }

func (r ResponseLimit) getLimit() string {
	limit := r.Limit
	if r.Limit == 0 {
		limit = defaultResponseSize
	}
	if r.Limit > responseLimit {
		limit = responseLimit
	}
	return intToString(limit)
}

// GetAccountSummaryPayload input params for GetAccountSummary method
type GetAccountSummaryPayload struct {
	DatetimePayload
	Currency string
}

// GetOrdersPayload input params for GetOrders method
type GetOrdersPayload struct {
	Account string
	Limit   ResponseLimit
	DatetimeRangePayload
}

// GetActiveOrderPayload optional input params for GetActiveOrder method
type GetActiveOrderPayload struct {
	Account, SymbolID string
	Limit             ResponseLimit
}

// GetTransactionsOptionalPayload optional params
// for GetTransactions endpoint function
type GetTransactionsOptionalPayload struct {
	OrderPos                             int
	Account, UUID, Asset, OrderID, Order string
	OffsetLimit                          OffsetLimitPayload
	DatetimeRangePayload
}

// GetOHLCOptionalPayload optional params for ohcl methods
type GetOHLCOptionalPayload struct {
	ResponseLimit
	StartStopOptionalPayload
}

// GetTicksByQuotesPayload method optional payload
type GetTicksByQuotesPayload GetOHLCOptionalPayload

// GetTicksByTradesPayload method optional payload
type GetTicksByTradesPayload GetTicksByQuotesPayload

// CancelOrderPayload method optional payload
type CancelOrderPayload struct {
	Action string `json:"action"`
}

// ReplaceOrderPayload method optional payload
type ReplaceOrderPayload struct {
	CancelOrderPayload
	Quantity      int    `json:"quantity"`
	LimitPrice    string `json:"limitPrice"`
	StopPrice     string `json:"stopPrice"`
	PriceDistance string `json:"priceDistance"`
}

// GetUserAccounts return the list of user accounts and their statuses
func (h HTTPApi) GetUserAccounts() (*UserAccounts, error) {
	m := NewUserAccounts()
	err := h.get(m, requestData{
		action:  accountsAction,
		version: h.getVersion(),
	})
	return m, err
}

// GetCurrencies return the list of available currencies
func (h HTTPApi) GetCurrencies() (*Сurrencys, error) {
	m := NewСurrencys()
	err := h.get(m, requestData{
		action:  crossratesAction,
		version: h.getVersion(),
	})
	return m, err
}

// GetExchanges return list of exchanges
func (h HTTPApi) GetExchanges() (*Exchanges, error) {
	m := NewExchanges()
	err := h.get(m, requestData{
		action:  exchangesAction,
		version: h.getVersion(),
	})
	return m, err
}

// GetChangesV1 return the list of daily changes for all or requested instruments
func (h HTTPApi) GetChangesV1(symbolIDs ...string) (*ChangesV1, error) {
	m := NewChangesV1()
	err := h.get(m, requestData{
		action:     changeAction,
		pathParams: joinWithCommaSeparator(symbolIDs),
		version:    APIv1,
	})
	return m, err
}

// GetChangesV2 return the list of daily changes for all or requested instruments
func (h HTTPApi) GetChangesV2(symbolIDs ...string) (*ChangesV2, error) {
	m := NewChangesV2()
	err := h.get(m, requestData{
		action:     changeAction,
		pathParams: joinWithCommaSeparator(symbolIDs),
		version:    APIv2,
	})
	return m, err
}

// GetChangesV3 return the list of daily changes for all or requested instruments
func (h HTTPApi) GetChangesV3(symbolIDs ...string) (*ChangesV3, error) {
	m := NewChangesV3()
	err := h.get(m, requestData{
		action:     changeAction,
		pathParams: joinWithCommaSeparator(symbolIDs),
		version:    APIv3,
	})
	return m, err
}

// GetSymbolsByGroupV1 return financial instruments which belong to specified group
func (h HTTPApi) GetSymbolsByGroupV1(group string) (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action:     groupsAction,
		pathParams: group,
		version:    APIv1,
	})
	return m, err
}

// GetSymbolsByGroupV2 return financial instruments which belong to specified group
func (h HTTPApi) GetSymbolsByGroupV2(group string) (*SymbolsV2, error) {
	m := NewSymbolsV2()
	err := h.get(m, requestData{
		action:     groupsAction,
		pathParams: group,
		version:    APIv2,
	})
	return m, err
}

// GetSymbolsByGroupV3 return financial instruments which belong to specified group
func (h HTTPApi) GetSymbolsByGroupV3(group string) (*SymbolsV3, error) {
	m := NewSymbolsV3()
	err := h.get(m, requestData{
		action:     groupsAction,
		pathParams: group,
		version:    APIv3,
	})
	return m, err
}

// GetSymbolsByTypeV1 return financial instruments of the requested type
// symbolType: instument type like STOCK, FUTURE, BOND, CURRENCY etc
func (h HTTPApi) GetSymbolsByTypeV1(symbolType string) (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action:     typesAction,
		pathParams: symbolType,
		version:    APIv1,
	})
	return m, err
}

// GetSymbolsByTypeV2 return financial instruments of the requested type
// symbolType: instument types like STOCK, FUTURE, BOND, CURRENCY etc
func (h HTTPApi) GetSymbolsByTypeV2(symbolType string) (*SymbolsV2, error) {
	m := NewSymbolsV2()
	err := h.get(m, requestData{
		action:     typesAction,
		pathParams: symbolType,
		version:    APIv2,
	})
	return m, err
}

// GetSymbolsByTypeV3 return financial instruments of the requested type
// symbolType: instument types like STOCK, FUTURE, BOND, CURRENCY etc
func (h HTTPApi) GetSymbolsByTypeV3(symbolType string) (*SymbolsV3, error) {
	m := NewSymbolsV3()
	err := h.get(m, requestData{
		action:     typesAction,
		pathParams: symbolType,
		version:    APIv3,
	})
	return m, err
}

// GetSymbolsV1 return list of instruments available for authorized user
func (h HTTPApi) GetSymbolsV1() (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action:  symbolsAction,
		version: APIv1,
	})
	return m, err
}

// GetSymbolsV2 return list of instruments available for authorized user
func (h HTTPApi) GetSymbolsV2() (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action:  symbolsAction,
		version: APIv2,
	})
	return m, err
}

// GetSymbolsV3 return list of instruments available for authorized user
func (h HTTPApi) GetSymbolsV3() (*SymbolsV3, error) {
	m := NewSymbolsV3()
	err := h.get(m, requestData{
		action:  symbolsAction,
		version: APIv3,
	})
	return m, err
}

// GetSymbolV1 return instrument available for authorized user
func (h HTTPApi) GetSymbolV1(symbolID string) (*SymbolV1, error) {
	m := NewSymbolV1()
	err := h.get(m, requestData{
		action:     symbolsAction,
		version:    APIv1,
		pathParams: symbolID,
	})
	return m, err
}

// GetSymbolV2 return instrument available for authorized user
func (h HTTPApi) GetSymbolV2(symbolID string) (*SymbolV2, error) {
	m := NewSymbolV2()
	err := h.get(m, requestData{
		action:     symbolsAction,
		version:    APIv2,
		pathParams: symbolID,
	})
	return m, err
}

// GetSymbolV3 return instrument available for authorized user
func (h HTTPApi) GetSymbolV3(symbolID string) (*SymbolV3, error) {
	m := NewSymbolV3()
	err := h.get(m, requestData{
		action:     symbolsAction,
		version:    APIv3,
		pathParams: symbolID,
	})
	return m, err
}

// GetSymbolschedule return financial schedule for requested instrument
func (h HTTPApi) GetSymbolschedule(
	symbolID string, useTypes bool) (*Schedule, error) {

	m := NewSchedule()
	err := h.get(m, requestData{
		action: symbolsAction,
		pathParams: joinWithSlashSeparator(
			[]string{symbolID, scheduleAction},
		),
		version: h.getVersion(),
		queryStringParams: QueryStringParams{
			typesAction: boolToString(useTypes),
		},
	})
	return m, err
}

// GetTypes return list of known instrument types
func (h HTTPApi) GetTypes() (*Types, error) {
	m := NewTypes()
	err := h.get(m, requestData{
		action:  typesAction,
		version: h.getVersion(),
	})
	return m, err
}

// GetSymbolSpec return additional parameters for requested instrument
func (h HTTPApi) GetSymbolSpec(symbol string) (*SymbolSpecification, error) {
	m := NewSymbolSpecification()
	err := h.get(m, requestData{
		action: symbolsAction,
		pathParams: joinWithSlashSeparator(
			[]string{symbol, specificationAction},
		),
		version: h.getVersion(),
	})
	return m, err
}

// GetGroups return list of available instrument groups
func (h HTTPApi) GetGroups() (*Groups, error) {
	m := NewGroups()
	err := h.get(m, requestData{
		action:  groupsAction,
		version: h.getVersion(),
	})
	return m, err
}

// GetNearestV1 return financial instrument which
// has the nearest expiration in the group
func (h HTTPApi) GetNearestV1(group string) (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action: groupsAction,
		pathParams: joinWithSlashSeparator(
			[]string{group, nearestAction},
		),
		version: APIv1,
	})
	return m, err
}

// GetNearestV2 return financial instrument which
// has the nearest expiration in the group
func (h HTTPApi) GetNearestV2(group string) (*SymbolsV2, error) {
	m := NewSymbolsV2()
	err := h.get(m, requestData{
		action: groupsAction,
		pathParams: joinWithSlashSeparator(
			[]string{group, nearestAction},
		),
		version: APIv2,
	})
	return m, err
}

// GetNearestV3 return financial instrument which
// has the nearest expiration in the group
func (h HTTPApi) GetNearestV3(group string) (*SymbolsV3, error) {
	m := NewSymbolsV3()
	err := h.get(m, requestData{
		action: groupsAction,
		pathParams: joinWithSlashSeparator(
			[]string{group, nearestAction},
		),
		version: APIv3,
	})
	return m, err
}

// GetSymbolsByExchV1 return the requested exchange financial instruments
func (h HTTPApi) GetSymbolsByExchV1(exchange string) (*SymbolsV1, error) {
	m := NewSymbolsV1()
	err := h.get(m, requestData{
		action:     exchangesAction,
		pathParams: exchange,
		version:    APIv1,
	})
	return m, err
}

// GetSymbolsByExchV2 return the requested exchange financial instruments
func (h HTTPApi) GetSymbolsByExchV2(exchange string) (*SymbolsV2, error) {
	m := NewSymbolsV2()
	err := h.get(m, requestData{
		action:     exchangesAction,
		pathParams: exchange,
		version:    APIv2,
	})
	return m, err
}

// GetSymbolsByExchV3 return the requested exchange financial instruments
func (h HTTPApi) GetSymbolsByExchV3(exchange string) (*SymbolsV3, error) {
	m := NewSymbolsV3()
	err := h.get(m, requestData{
		action:     exchangesAction,
		pathParams: exchange,
		version:    APIv3,
	})
	return m, err
}

// GetLastQuote return the last quote for the specified financial instrument
func (h HTTPApi) GetLastQuote(symbol string) (*Quote, error) {
	m := NewQuote()
	err := h.get(m, requestData{
		action: feedAction,
		pathParams: joinWithSlashSeparator(
			[]string{symbol, lastAction},
		),
	})
	return m, err
}

// GetOHLCTrades return the list of OHLC candles
// for the specified financial instrument and duration
func (h HTTPApi) GetOHLCTrades(symbol string, duration int,
	g GetOHLCOptionalPayload) (*OHLCTrades, error) {

	m := NewOHLCTrades()
	u := requestData{
		queryStringParams: QueryStringParams{
			"size": g.getLimit(),
			"type": tradesDataType,
		},
		action: ohlcAction,
		pathParams: joinWithSlashSeparator(
			[]string{symbol, intToString(duration)},
		),
	}
	if g.Start != 0 {
		u.queryStringParams["start"] = intToString(g.Start)
	}
	if g.Stop != 0 {
		u.queryStringParams["end"] = intToString(g.Stop)
	}
	err := h.get(m, u)
	return m, err
}

// GetOHLCQuotes return the list of OHLC candles
// for the specified financial instrument and duration
func (h HTTPApi) GetOHLCQuotes(symbol string, duration int,
	g GetOHLCOptionalPayload) (*OHLCQuotes, error) {

	m := NewOHLCQuotes()
	u := requestData{
		queryStringParams: QueryStringParams{
			"size": g.getLimit(),
			"type": quotesDataType,
		},
		action: ohlcAction,
		pathParams: joinWithSlashSeparator(
			[]string{symbol, intToString(duration)},
		),
	}
	if g.Start != 0 {
		u.queryStringParams["start"] = intToString(g.Start)
	}
	if g.Stop != 0 {
		u.queryStringParams["end"] = intToString(g.Stop)
	}
	err := h.get(m, u)
	return m, err
}

// GetTicksByQuotes return the list of ticks for the specified financial instrument
// aggregated by quotes data type
func (h HTTPApi) GetTicksByQuotes(
	symbol string, g GetTicksByQuotesPayload) (interface{}, error) {

	m := NewOHLCQuotes()
	u := requestData{
		action:     ticksAction,
		pathParams: symbol,
		queryStringParams: QueryStringParams{
			"size": g.getLimit(),
			"type": quotesDataType,
		},
	}
	if g.Start != 0 {
		u.queryStringParams["start"] = intToString(g.Start)
	}
	if g.Stop != 0 {
		u.queryStringParams["end"] = intToString(g.Stop)
	}
	err := h.get(m, u)
	return m, err
}

// GetTicksByTrades return the list of ticks for the specified financial instrument
// aggregated by trades data type
func (h HTTPApi) GetTicksByTrades(
	symbol string, g GetTicksByTradesPayload) (interface{}, error) {

	m := NewOHLCTrades()
	u := requestData{
		action:     ticksAction,
		pathParams: symbol,
		queryStringParams: QueryStringParams{
			"size": g.getLimit(),
			"type": quotesDataType,
		},
	}
	if g.Start != 0 {
		u.queryStringParams["start"] = intToString(g.Start)
	}
	if g.Stop != 0 {
		u.queryStringParams["end"] = intToString(g.Stop)
	}
	err := h.get(m, u)
	return m, err
}

// GetAccountSummary return the summary for the specified account
func (h HTTPApi) GetAccountSummary(account, currency string,
	p GetAccountSummaryPayload) (interface{}, error) {

	m := NewSummary()
	c := stringToUpperCase(currency)

	err := h.get(m, requestData{
		action:  summaryAction,
		version: h.getVersion(),
		pathParams: joinWithSlashSeparator(
			[]string{account, c},
		),
	})
	return m, err
}

func (h HTTPApi) getTransactionsQueryString(symbol string, opType []string,
	p GetTransactionsOptionalPayload) QueryStringParams {

	payload := QueryStringParams{
		"symbolId":      symbol,
		"uuid":          p.UUID,
		"accountId":     p.Account,
		"asset":         p.Asset,
		"order":         p.Order,
		"orderId":       p.OrderID,
		"operationType": strungSliceToString(opType),
		"offset":        intToString(p.OffsetLimit.Offset),
		"limit":         intToString(p.OffsetLimit.Limit),
		"orderPos":      intToString(p.OrderPos),
		"fromDate":      intToString(p.From),
		"toDate":        intToString(p.To),
	}
	return payload
}

// GetTransactionsV1 return the list of transactions with the specified filter
func (h HTTPApi) GetTransactionsV1(symbol string, opType []string,
	p GetTransactionsOptionalPayload) (*TransactionsV1, error) {

	queryStringData := h.getTransactionsQueryString(symbol, opType, p)
	m := NewTransactionsV1()
	err := h.get(m, requestData{
		version:           APIv1,
		queryStringParams: queryStringData,
		action:            transactionsAction,
	})
	return m, err
}

// GetTransactionsV2 return the list of transactions with the specified filter
func (h HTTPApi) GetTransactionsV2(symbol string, opType []string,
	p GetTransactionsOptionalPayload) (*TransactionsV2, error) {

	queryStringData := h.getTransactionsQueryString(symbol, opType, p)
	m := NewTransactionsV2()
	err := h.get(m, requestData{
		version:           APIv2,
		queryStringParams: queryStringData,
		action:            transactionsAction,
	})
	return m, err
}

// GetTransactionsV3 return the list of transactions with the specified filter
func (h HTTPApi) GetTransactionsV3(symbol string, opType []string,
	p GetTransactionsOptionalPayload) (*TransactionsV3, error) {
	queryStringData := h.getTransactionsQueryString(symbol, opType, p)
	m := NewTransactionsV3()
	err := h.get(m, requestData{
		version:           APIv3,
		queryStringParams: queryStringData,
		action:            transactionsAction,
	})
	return m, err
}

// GetOrdersV1 return the list of historical orders
func (h HTTPApi) GetOrdersV1(g GetOrdersPayload) (*OrdersV1, error) {
	m := NewOrdersV1()
	u := requestData{
		action:   ordersAction,
		version:  APIv1,
		category: TRADEAPICategory,
		queryStringParams: QueryStringParams{
			"limit": g.Limit.getLimit(),
		},
	}
	if g.From > 0 {
		u.queryStringParams["from"] = g.getFromParam()
	}
	if g.To > 0 {
		u.queryStringParams["to"] = g.getToParam()
	}
	err := h.get(m, u)
	return m, err
}

// GetOrdersV2 return the list of historical orders
func (h HTTPApi) GetOrdersV2(g GetOrdersPayload) (*OrdersV2, error) {
	m := NewOrdersV2()
	u := requestData{
		action:   ordersAction,
		version:  APIv2,
		category: TRADEAPICategory,
		queryStringParams: QueryStringParams{
			"limit": g.Limit.getLimit(),
		},
	}
	if g.From > 0 {
		u.queryStringParams["from"] = g.getFromParam()
	}
	if g.To > 0 {
		u.queryStringParams["to"] = g.getToParam()
	}
	err := h.get(m, u)
	return m, err
}

// GetOrdersV3 return the list of historical orders
func (h HTTPApi) GetOrdersV3(g GetOrdersPayload) (*OrdersV3, error) {
	m := NewOrdersV3()
	u := requestData{
		action:   ordersAction,
		version:  APIv3,
		category: TRADEAPICategory,
		queryStringParams: QueryStringParams{
			"limit": g.Limit.getLimit(),
		},
	}
	if g.From > 0 {
		u.queryStringParams["from"] = g.getFromParam()
	}
	if g.To > 0 {
		u.queryStringParams["to"] = g.getToParam()
	}
	err := h.get(m, u)
	return m, err
}

// GetOrderV1 return the order with specified identifier
func (h HTTPApi) GetOrderV1(orderID string) (*OrderV1, error) {
	m := NewOrderV1()
	err := h.get(m, requestData{
		action:     ordersAction,
		pathParams: orderID,
		version:    APIv1,
		category:   TRADEAPICategory,
	})
	return m, err
}

// GetOrderV2 return the order with specified identifier
func (h HTTPApi) GetOrderV2(orderID string) (*OrderV2, error) {
	m := NewOrderV2()
	err := h.get(m, requestData{
		action:     ordersAction,
		pathParams: orderID,
		version:    APIv2,
		category:   TRADEAPICategory,
	})
	return m, err
}

// GetOrderV3 return the order with specified identifier
func (h HTTPApi) GetOrderV3(orderID string) (*OrderV3, error) {
	m := NewOrderV3()
	err := h.get(m, requestData{
		action:     ordersAction,
		pathParams: orderID,
		version:    APIv3,
		category:   TRADEAPICategory,
	})
	return m, err
}

// GetActiveOrdersV1 return the list of active trading orders
func (h HTTPApi) GetActiveOrdersV1() (*OrdersV1, error) {
	m := NewOrdersV1()
	err := h.get(m, requestData{
		action:   activeOrdersAction,
		version:  APIv1,
		category: TRADEAPICategory,
	})
	return m, err
}

// GetActiveOrdersV2 return the list of active trading orders
func (h HTTPApi) GetActiveOrdersV2() (*OrdersV2, error) {
	m := NewOrdersV2()
	err := h.get(m, requestData{
		action:   activeOrdersAction,
		version:  APIv2,
		category: TRADEAPICategory,
	})
	return m, err
}

// GetActiveOrdersV3 return the list of active trading orders
func (h HTTPApi) GetActiveOrdersV3() (*OrdersV3, error) {
	m := NewOrdersV3()
	err := h.get(m, requestData{
		action:   activeOrdersAction,
		version:  APIv3,
		category: TRADEAPICategory,
	})
	return m, err
}

// PlaceOrderV1 place new trading OrderV1
func (h HTTPApi) PlaceOrderV1(o *OrderSentTypeV1) error {
	err := h.post(o, requestData{
		action:   ordersAction,
		version:  APIv1,
		category: TRADEAPICategory,
	})
	return err
}

// PlaceOrderV2 place new trading OrderV2
func (h HTTPApi) PlaceOrderV2(o *OrderSentTypeV2) error {
	err := h.post(o, requestData{
		action:   ordersAction,
		version:  APIv2,
		category: TRADEAPICategory,
	})
	return err
}

// PlaceOrderV3 place new trading OrderV3
func (h HTTPApi) PlaceOrderV3(o *OrderSentTypeV3) error {
	err := h.post(o, requestData{
		action:   ordersAction,
		version:  APIv3,
		category: TRADEAPICategory,
	})
	return err
}

// CancelOrder cancel trading order
func (h HTTPApi) CancelOrder(orderID string, c CancelOrderPayload) error {
	err := h.post(c, requestData{
		action:     cancelOrder,
		category:   TRADEAPICategory,
		pathParams: orderID,
		version:    h.getVersion(),
	})
	return err
}

// ReplaceOrder replace trading order
func (h HTTPApi) ReplaceOrder(orderID string, r ReplaceOrderPayload) error {
	err := h.post(r, requestData{
		action:     replaceOrder,
		category:   TRADEAPICategory,
		pathParams: orderID,
		version:    h.getVersion(),
	})
	return err
}

// GetOrdersStream return the life quote stream
// for the specified financial instruments
func (h HTTPApi) GetOrdersStream() (chan []byte, chan bool) {
	u := requestData{
		category: TRADEAPICategory,
		action:   ordersStreamAction,
		version:  h.getVersion(),
	}
	outChan, stopChan := h.runStream(u)
	return outChan, stopChan
}

// GetExecOrdersStream return the life quote stream
//for the specified financial instruments
func (h HTTPApi) GetExecOrdersStream() (chan []byte, chan bool) {
	u := requestData{
		category: TRADEAPICategory,
		action:   execOrdersStreamAction,
		version:  h.getVersion(),
	}
	outChan, stopChan := h.runStream(u)
	return outChan, stopChan
}

// GetTradeStream return the trades stream
// for the specified financial instruments
func (h HTTPApi) GetTradeStream(symbol string) (chan []byte, chan bool) {
	u := requestData{
		category:   TRADEAPICategory,
		action:     tradeStreamAction,
		version:    h.getVersion(),
		pathParams: symbol,
	}
	outChan, stopChan := h.runStream(u)
	return outChan, stopChan
}

// GetQuoteStream return the life quote stream
// for the specified financial instruments
func (h HTTPApi) GetQuoteStream() (chan []byte, chan bool) {
	u := requestData{
		category: TRADEAPICategory,
		action:   execOrdersStreamAction,
		version:  h.getVersion(),
	}
	outChan, stopChan := h.runStream(u)
	return outChan, stopChan
}
