package httplib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func newMockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/md/1.0/accounts", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"status":"Full","accountId":"WWB1220.001"}, 
            {"status":"Full","accountId": "WWB1220.002"}]`)
	})
	mux.HandleFunc("/md/2.0/accounts", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"status":"Full","accountId":"WWB1220.001"},
            {"status":"Full","accountId": "WWB1220.002"}]`)
	})
	mux.HandleFunc("/md/3.0/accounts", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"status":"Full","accountId":"WWB1220.001"},
            {"status":"Full","accountId": "WWB1220.002"}]`)
	})
	mux.HandleFunc("/md/1.0/crossrates", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"currencies":["PLN","UST","ZEC"]}`)
	})
	mux.HandleFunc("/md/2.0/crossrates", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"currencies":["PLN","UST","ZEC"]}`)
	})
	mux.HandleFunc("/md/3.0/crossrates", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"currencies":["PLN","UST","ZEC"]}`)
	})
	mux.HandleFunc("/md/1.0/exchanges", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"id":"USD","name":"USD: Other non-US Bonds","country":"LK"},
        {"id":"DSE","name":"DSE Dusseldorf Stock Exchange","country":"AU"},
        {"id":"GBP","name":"GBP: Eurobonds","country":"GG"}]`)
	})
	mux.HandleFunc("/md/2.0/exchanges", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"id":"USD","name":"USD: Other non-US Bonds","country":"LK"},
        {"id":"DSE","name":"DSE Dusseldorf Stock Exchange","country":"AU"},
        {"id":"GBP","name":"GBP: Eurobonds","country":"GG"}]`)
	})
	mux.HandleFunc("/md/3.0/exchanges", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"id":"USD","name":"USD: Other non-US Bonds","country":"LK"},
        {"id":"DSE","name":"DSE Dusseldorf Stock Exchange","country":"AU"},
        {"id":"GBP","name":"GBP: Eurobonds","country":"GG"}]`)
	})
	mux.HandleFunc("/md/2.0/groups", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"group":"VNQ","name":"Option on Vanguard REIT ETF - DNQ",
            "types":["OPTION"],"exchange":"CBOE"},
            {"group":"PTEN","name":"Patterson-UTI Energy",
            "types":["OPTION"],"exchange":"CBOE"},
            {"group":"MA","name":"Mastercard","types":["OPTION"],
            "exchange":"CBOE"}]`)
	})
	mux.HandleFunc("/md/1.0/groups/CBOE/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, `[{"optionData":{"optionGroupId":"CBOE.CBOE.20X2020.P*",
            "strikePrice":65,"right":"PUT"},"name":"Cboe Global Markets",
            "description":"Cboe Global Markets 20 Nov 2020 PUT 65","country":"US",
            "exchange":"CBOE","id":"CBOE.CBOE.20X2020.P65","currency":"USD",
            "mpi":0.05,"type":"OPTION","ticker":"CBOE","expiration":1605906000000,
            "group":"CBOE"},{"optionData":{"optionGroupId":"CBOE.CBOE.20X2020.P*",
            "strikePrice":80,"right":"PUT"},"name":"Cboe Global Markets",
            "description":"Cboe Global Markets 20 Nov 2020 PUT 80","country":"US",
            "exchange":"CBOE","id":"CBOE.CBOE.20X2020.P80","currency":"USD",
            "mpi":0.05,"type":"OPTION","ticker":"CBOE","expiration":1605906000000,
            "group":"CBOE"}]`)
	})
	mux.HandleFunc("/md/1.0/symbols/CBOE.CBOE.20X2020.P65/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"optionData":{"optionGroupId":"CBOE.CBOE.20X2020.P*",
		"right":"PUT","strikePrice":65},"i18n":"","name":"Cboe Global Markets",
		"description":"Cboe Global Markets 20 Nov 2020 PUT 65","country":"US",
		"exchange":"CBOE","id":"CBOE.CBOE.20X2020.P65","currency":"USD",
		"mpi":0.05,"type":"OPTION","ticker":"CBOE","expiration":1605906000000,
		"group":"CBOE"}`)
	})
	mux.HandleFunc("/md/2.0/symbols/CBOE.CBOE.20X2020.P65/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"optionData":{"optionGroupId":"CBOE.CBOE.20X2020.P*",
		"right":"PUT","strikePrice":"65"},"i18n":"","name":"Cboe Global Markets",
		"description":"Cboe Global Markets 20 Nov 2020 PUT 65","country":"US",
		"exchange":"CBOE","id":"CBOE.CBOE.20X2020.P65","currency":"USD",
		"mpi":"0.05","type":"OPTION","ticker":"CBOE","expiration":1605906000000,
		"group":"CBOE"}`)
	})
	mux.HandleFunc("/md/3.0/symbols/CBOE.CBOE.20X2020.P65/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"optionData":{"optionGroupId":"CBOE.CBOE.20X2020.P*",
		"right":"PUT","strikePrice":"65"},"i18n":"","name":"Cboe Global Markets",
		"description":"Cboe Global Markets 20 Nov 2020 PUT 65","country":"US",
		"exchange":"CBOE","id":"CBOE.CBOE.20X2020.P65","currency":"USD",
		"mpi":0.05,"type":"OPTION","ticker":"CBOE","expiration":1605906000000,
		"group":"CBOE"}`)
	})
	mux.HandleFunc("/md/2.0/types", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"id":"FX_SPOT"},{"id":"FUND"},{"id":"CURRENCY"},
		{"id":"CALENDAR_SPREAD"},{"id":"CFD"},{"id":"STOCK"},
		{"id":"FUTURE"},{"id":"OPTION"},{"id":"BOND"}]`)
	})
	mux.HandleFunc("/md/2.0/summary/WWB1220.002/USD/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `{"currencies":[{"code":"EUR","convertedValue":"1181995.82",
		"value":"996770.88"},{"code":"USD","convertedValue":"-6197460.15","value":"-6197460.15"}],
		"timestamp":1606081146000,"freeMoney":"0.0","netAssetValue":"-2672564.32",
		"moneyUsedForMargin":"778473.01","marginUtilization":"2.0",
		"positions":[{"convertedPnl":"-3836534.01","quantity":"20000","pnl":"-3836534.01",
		"convertedValue":"2342900.0","price":"117.145","id":"AAPL.NASDAQ","symbolType":"STOCK",
		"currency":"USD","averagePrice":"308.9717","value":"2342900.0"}],"sessionDate":null,
		"currency":"USD","account":"WWB1220.002"}`)
	})

	mux.HandleFunc("/trade/2.0/orders", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"placeTime":"2020-11-19T17:29:38.926Z","username":"ab@exante.eu",
		"orderState":{"status":"working","lastUpdate":"2020-11-19T17:29:38.938Z","fills":[]},
		"accountId":"WWB1220.001","id":"67856292-b62b-4f53-8481-55d33185cbe7",
		"clientTag":"d2e0746bab4c41c6afc54ade5082c3da","orderParameters":{"side":"buy",
		"duration":"good_till_cancel","quantity":"10","ocoGroup":null,"ifDoneParentId":null,
		"orderType":"limit","limitPrice":"110","instrument":"AAPL.NASDAQ"},
		"currentModificationId":"67856292-b62b-4f53-8481-55d33185cbe7"}]`)
	})
	mux.HandleFunc("/trade/2.0/orders/active", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, `[{"placeTime":"2020-11-19T17:29:38.926Z","username":"ab@exante.eu",
		"orderState":{"status":"working","lastUpdate":"2020-11-19T17:29:38.938Z",
		"fills":[]},"accountId":"WWB1220.001","id":"67856292-b62b-4f53-8481-55d33185cbe7",
		"clientTag":"d2e0746bab4c41c6afc54ade5082c3da","orderParameters":{"side":"buy",
		"duration":"good_till_cancel","quantity":"10","ocoGroup":null,"ifDoneParentId":null,
		"orderType":"limit","limitPrice":"110","instrument":"AAPL.NASDAQ"},
		"currentModificationId":"67856292-b62b-4f53-8481-55d33185cbe7"}]`)
	})

	ts := httptest.NewServer(mux)
	return ts
}

var s = newMockServer()

func makeFakeApp(version string) HTTPApi {
	return HTTPApi{
		httpClient: s.Client(),
		baseAPIURL: s.URL,
		version:    version,
	}
}

var fakeAPIv1 = makeFakeApp(APIv1)
var fakeAPIv2 = makeFakeApp(APIv2)
var fakeAPIv3 = makeFakeApp(APIv3)
var fakeAPIs = []HTTPApi{fakeAPIv1, fakeAPIv2, fakeAPIv3}

type apiInstance struct{ HTTPApi }

func TestNewAPI(t *testing.T) {
	for _, tt := range fakeAPIs {
		t.Run("TestNewAPI", func(t *testing.T) {
			got := NewAPI(
				tt.baseAPIURL, tt.version, tt.Auth.JWT.ApplicationID,
				tt.Auth.JWT.ClientID, tt.Auth.JWT.SharedKey, tt.Auth.JWT.JwtTTL,
				tt.Auth.Basic.Username, tt.Auth.Basic.Password)
			if got.getVersion() == emptyString {
				t.Error("Version getter error")
			}
			if reflect.TypeOf(got) != reflect.TypeOf(HTTPApi{}) {
				t.Error("Wrong api type")
			}
		})
	}
}

func TestHTTPApi_GetUserAccounts(t *testing.T) {
	expectedVal := &UserAccounts{
		UserAccount{Status: "Full", AccountID: "WWB1220.001"},
		UserAccount{Status: "Full", AccountID: "WWB1220.002"},
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		want        *UserAccounts
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, expectedVal, false},
		{APIv2, apiInstance{fakeAPIv2}, expectedVal, false},
		{APIv3, apiInstance{fakeAPIv3}, expectedVal, false},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetUserAccounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetUserAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetUserAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetCurrencies(t *testing.T) {
	expectedVal := &Сurrencys{[]Currency{"PLN", "UST", "ZEC"}}
	tests := []struct {
		name        string
		apiInstance apiInstance
		want        *Сurrencys
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, expectedVal, false},
		{APIv2, apiInstance{fakeAPIv2}, expectedVal, false},
		{APIv3, apiInstance{fakeAPIv3}, expectedVal, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetCurrencies()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetCurrencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetCurrencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetExchanges(t *testing.T) {
	u := Exchanges{
		{ID: "USD", Name: "USD: Other non-US Bonds", Country: "LK"},
		{ID: "DSE", Name: "DSE Dusseldorf Stock Exchange", Country: "AU"},
		{ID: "GBP", Name: "GBP: Eurobonds", Country: "GG"},
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		want        *Exchanges
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, &u, false},
		{APIv2, apiInstance{fakeAPIv2}, &u, false},
		{APIv3, apiInstance{fakeAPIv3}, &u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetExchanges()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetExchanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetExchanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetChangesV1(t *testing.T) {
	u := NewChangesV1()
	type apiInstance struct{ HTTPApi }
	type args struct{ symbolIDs []string }
	test := struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *ChangesV1
		wantErr     bool
	}{APIv1, apiInstance{fakeAPIv1}, args{[]string{"test1", "test2"}}, u, true}
	t.Run(test.name, func(t *testing.T) {
		got, err := test.apiInstance.GetChangesV1(test.args.symbolIDs...)
		if (err != nil) != test.wantErr {
			t.Errorf("HTTPApi.GetChangesV1() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("HTTPApi.GetChangesV1() = %v, want %v", got, test.want)
		}
	})
}

func TestHTTPApi_GetChangesV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ChangesV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetChangesV2(tt.args.symbolIDs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetChangesV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetChangesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetChangesV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ChangesV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetChangesV3(tt.args.symbolIDs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetChangesV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetChangesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByGroupV1(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	type args struct{ group string }
	u := SymbolsV1{
		SymbolV1{
			SymbolOptionData: SymbolOptionData{
				OptionIDGroup: "CBOE.CBOE.20X2020.P*",
				StrikePrice:   65,
				Right:         "PUT",
			},
			Name:        "Cboe Global Markets",
			Description: "Cboe Global Markets 20 Nov 2020 PUT 65",
			Country:     "US",
			Exchange:    "CBOE",
			ID:          "CBOE.CBOE.20X2020.P65",
			Currency:    "USD",
			Mpi:         0.05,
			Type:        "OPTION",
			Ticker:      "CBOE",
			Expiration:  1605906000000,
			Group:       "CBOE",
		}, SymbolV1{
			SymbolOptionData: SymbolOptionData{
				OptionIDGroup: "CBOE.CBOE.20X2020.P*",
				StrikePrice:   80404,
				Right:         "PUT",
			},
			Name:        "Cboe Global Markets",
			Description: "Cboe Global Markets 20 Nov 2020 PUT 80",
			Country:     "US",
			Exchange:    "CBOE",
			ID:          "CBOE.CBOE.20X2020.P80",
			Currency:    "USD",
			Mpi:         0.05,
			Type:        "OPTION",
			Ticker:      "CBOE",
			Expiration:  1605906000000,
			Group:       "CBOE",
		},
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *SymbolsV1
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, args{"CBOE"}, &u, false},
		{"1.0 404 success test", apiInstance{fakeAPIv1}, args{"SameFakeGroup"}, &u, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.apiInstance.GetSymbolsByGroupV1(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByGroupV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByGroupV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByGroupV2(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByGroupV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByGroupV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByGroupV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByGroupV3(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByGroupV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByGroupV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByTypeV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByTypeV1(tt.args.symbolType)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByTypeV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByTypeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByTypeV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByTypeV2(tt.args.symbolType)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByTypeV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByTypeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByTypeV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByTypeV3(tt.args.symbolType)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByTypeV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByTypeV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *SymbolsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsV1()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *SymbolsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsV2()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *SymbolsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsV3()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolV1(t *testing.T) {
	type fields struct{ HTTPApi }
	type args struct{ symbolID string }
	u := SymbolV1{
		SymbolOptionData: SymbolOptionData{
			OptionIDGroup: "CBOE.CBOE.20X2020.P*",
			StrikePrice:   65,
			Right:         "PUT",
		},
		Name:        "Cboe Global Markets",
		Description: "Cboe Global Markets 20 Nov 2020 PUT 65",
		Country:     "US",
		Exchange:    "CBOE",
		ID:          "CBOE.CBOE.20X2020.P65",
		Currency:    "USD",
		Mpi:         0.05,
		Type:        "OPTION",
		Ticker:      "CBOE",
		Expiration:  1605906000000,
		Group:       "CBOE",
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *SymbolV1
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, args{"CBOE.CBOE.20X2020.P65"}, &u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetSymbolV1(tt.args.symbolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolV2(t *testing.T) {
	type fields struct{ HTTPApi }
	type args struct{ symbolID string }
	u := SymbolV2{
		SymbolOptionDataV2: SymbolOptionDataV2{
			OptionIDGroup: "CBOE.CBOE.20X2020.P*",
			StrikePrice:   "65",
			Right:         "PUT",
		},
		Name:        "Cboe Global Markets",
		Description: "Cboe Global Markets 20 Nov 2020 PUT 65",
		Country:     "US",
		Exchange:    "CBOE",
		ID:          "CBOE.CBOE.20X2020.P65",
		Currency:    "USD",
		Mpi:         "0.05",
		Type:        "OPTION",
		Ticker:      "CBOE",
		Expiration:  1605906000000,
		Group:       "CBOE",
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *SymbolV2
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, args{"CBOE.CBOE.20X2020.P65"}, &u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetSymbolV2(tt.args.symbolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.ValueOf(got) == reflect.ValueOf(tt.want) {
				t.Errorf("HTTPApi.GetSymbolV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolschedule(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbolID string
		useTypes bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolschedule(tt.args.symbolID, tt.args.useTypes)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolschedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolschedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetTypes(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	u := &Types{
		Type{"FX_SPOT"},
		Type{"FUND"},
		Type{"CURRENCY"},
		Type{"CALENDAR_SPREAD"},
		Type{"CFD"},
		Type{"STOCK"},
		Type{"FUTURE"},
		Type{"OPTION"},
		Type{"BOND"},
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		want        *Types
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolSpec(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolSpecification
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolSpec(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetGroups(t *testing.T) {
	u := &Groups{
		{
			Group:    "VNQ",
			Name:     "Option on Vanguard REIT ETF - DNQ",
			Types:    []string{"OPTION"},
			Exchange: "CBOE",
		}, {
			Group:    "PTEN",
			Name:     "Patterson-UTI Energy",
			Types:    []string{"OPTION"},
			Exchange: "CBOE",
		}, {
			Group:    "MA",
			Name:     "Mastercard",
			Types:    []string{"OPTION"},
			Exchange: "CBOE",
		},
	}
	type apiInstance struct{ HTTPApi }
	test := struct {
		name        string
		apiInstance apiInstance
		want        *Groups
		wantErr     bool
	}{APIv2, apiInstance{fakeAPIv2}, u, false}
	t.Run(test.name, func(t *testing.T) {
		got, err := test.apiInstance.GetGroups()
		if (err != nil) != test.wantErr {
			t.Errorf("HTTPApi.GetGroups() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("HTTPApi.GetGroups() = %v, want %v", got, test.want)
		}
	})
}

func TestHTTPApi_GetNearestV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetNearestV1(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetNearestV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetNearestV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetNearestV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetNearestV2(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetNearestV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetNearestV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetNearestV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetNearestV3(tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetNearestV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetNearestV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByExchV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		exchange string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByExchV1(tt.args.exchange)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByExchV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByExchV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByExchV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		exchange string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByExchV2(tt.args.exchange)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByExchV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByExchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetSymbolsByExchV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		exchange string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SymbolsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetSymbolsByExchV3(tt.args.exchange)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetSymbolsByExchV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetSymbolsByExchV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetLastQuote(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Quote
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetLastQuote(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetLastQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetLastQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOHLCTrades(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol   string
		duration int
		g        GetOHLCOptionalPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OHLCTrades
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOHLCTrades(tt.args.symbol, tt.args.duration, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOHLCTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOHLCTrades() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOHLCQuotes(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol   string
		duration int
		g        GetOHLCOptionalPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OHLCQuotes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOHLCQuotes(tt.args.symbol, tt.args.duration, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOHLCQuotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOHLCQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetTicksByQuotes(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
		g      GetTicksByQuotesPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetTicksByQuotes(tt.args.symbol, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTicksByQuotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTicksByQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetTicksByTrades(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
		g      GetTicksByTradesPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetTicksByTrades(tt.args.symbol, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTicksByTrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTicksByTrades() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetAccountSummary(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	u := &Summary{
		Currencies: []CurrencyPos{
			{Code: "EUR", ConvertedValue: "1181995.82", Value: "996770.88"},
			{Code: "USD", ConvertedValue: "-6197460.15", Value: "-6197460.15"},
		},
		Positions: []Position{
			{ConvertedPNL: "-3836534.01", Quantity: "20000", PNL: "-3836534.01",
				ConvertedValue: "2342900.0", Price: "117.145", ID: "AAPL.NASDAQ",
				SymbolType: "STOCK", Currency: "USD", AveragePrice: "308.9717",
				Value: "2342900.0"},
		},
		Currency:           "USD",
		Account:            "WWB1220.002",
		Timestamp:          1606081146000,
		FreeMoney:          "0.0",
		NetAssetValue:      "-2672564.32",
		MoneyUsedForMargin: "778473.01",
		MarginUtilization:  "2.0",
	}
	type args struct {
		account  string
		currency string
		p        GetAccountSummaryPayload
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *Summary
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, args{
			account: "WWB1220.002", currency: "USD"}, u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetAccountSummary(tt.args.account, tt.args.currency, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetAccountSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetAccountSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetTransactionsV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
		opType []string
		p      GetTransactionsOptionalPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransactionsV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetTransactionsV1(tt.args.symbol, tt.args.opType, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTransactionsV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTransactionsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetTransactionsV2(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	u := NewTransactionsV2()
	type args struct {
		symbol string
		opType []string
		p      GetTransactionsOptionalPayload
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *TransactionsV2
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, args{symbol: "CBOE.CBOE.20X2020.P65",
			opType: []string{"INTEREST"}, p: GetTransactionsOptionalPayload{}}, u, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.apiInstance.GetTransactionsV2(tt.args.symbol, tt.args.opType, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTransactionsV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPApi_GetTransactionsV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
		opType []string
		p      GetTransactionsOptionalPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransactionsV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetTransactionsV3(tt.args.symbol, tt.args.opType, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetTransactionsV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTransactionsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOrdersV1(t *testing.T) {
	type fields struct{ HTTPApi }
	type args struct{ g GetOrdersPayload }
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrdersV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOrdersV1(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrdersV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrdersV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOrdersV2(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	u := &OrdersV2{
		{
			PlaceTime: "2020-11-19T17:29:38.926Z",
			Username:  "ab@exante.eu",
			OrderState: OrderState{
				Status:     "working",
				LastUpdate: "2020-11-19T17:29:38.938Z",
			},
			OrderParameters: OrderParameters{
				Side:       "buy",
				Duration:   "good_till_cancel",
				Quantity:   "10",
				OrderType:  "limit",
				LimitPrice: "110",
				Instrument: "AAPL.NASDAQ",
			},
			AccountID:             "WWB1220.001",
			ID:                    "67856292-b62b-4f53-8481-55d33185cbe7",
			ClientTag:             "d2e0746bab4c41c6afc54ade5082c3da",
			CurrentModificationID: "67856292-b62b-4f53-8481-55d33185cbe7",
		},
	}
	type args struct {
		g GetOrdersPayload
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *OrdersV2
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, args{
			GetOrdersPayload{
				Limit: ResponseLimit{Limit: 1}},
		}, u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.apiInstance.GetOrdersV2(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrdersV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPApi_GetOrdersV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		g GetOrdersPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrdersV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOrdersV3(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrdersV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrdersV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOrderV1(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	type args struct{ orderID string }
	u := OrderV1{}
	tests := []struct {
		name        string
		apiInstance apiInstance
		args        args
		want        *OrderV1
		wantErr     bool
	}{
		{APIv1, apiInstance{fakeAPIv1}, args{"test"}, &u, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.apiInstance.GetOrderV1(tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrderV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrderV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOrderV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderV2
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOrderV2(tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrderV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetOrderV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetOrderV3(tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetOrderV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrderV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetActiveOrdersV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *OrdersV1
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetActiveOrdersV1()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetActiveOrdersV1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetActiveOrdersV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_GetActiveOrdersV2(t *testing.T) {
	type apiInstance struct{ HTTPApi }
	u := &OrdersV2{
		{
			PlaceTime: "2020-11-19T17:29:38.926Z",
			Username:  "ab@exante.eu",
			OrderState: OrderState{
				Status:     "working",
				LastUpdate: "2020-11-19T17:29:38.938Z",
			},
			AccountID: "WWB1220.001",
			ID:        "67856292-b62b-4f53-8481-55d33185cbe7",
			ClientTag: "d2e0746bab4c41c6afc54ade5082c3da",
			OrderParameters: OrderParameters{
				Side:       "buy",
				Duration:   "good_till_cancel",
				Quantity:   "10",
				OrderType:  "limit",
				LimitPrice: "110",
				Instrument: "AAPL.NASDAQ",
			},
			CurrentModificationID: "67856292-b62b-4f53-8481-55d33185cbe7",
		},
	}
	tests := []struct {
		name        string
		apiInstance apiInstance
		want        *OrdersV2
		wantErr     bool
	}{
		{APIv2, apiInstance{fakeAPIv2}, u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.apiInstance.GetActiveOrdersV2()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetActiveOrdersV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHTTPApi_GetActiveOrdersV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *OrdersV3
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, err := h.GetActiveOrdersV3()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.GetActiveOrdersV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetActiveOrdersV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPApi_PlaceOrderV1(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		o *OrderSentTypeV1
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			if err := h.PlaceOrderV1(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.PlaceOrderV1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPApi_PlaceOrderV2(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		o *OrderSentTypeV2
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			if err := h.PlaceOrderV2(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.PlaceOrderV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPApi_PlaceOrderV3(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		o *OrderSentTypeV3
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			if err := h.PlaceOrderV3(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.PlaceOrderV3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPApi_CancelOrder(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		orderID string
		c       CancelOrderPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			if err := h.CancelOrder(tt.args.orderID, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPApi_ReplaceOrder(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		orderID string
		r       ReplaceOrderPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			if err := h.ReplaceOrder(tt.args.orderID, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HTTPApi.ReplaceOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTPApi_GetOrdersStream(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name   string
		fields fields
		want   chan []byte
		want1  chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, got1 := h.GetOrdersStream()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetOrdersStream() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HTTPApi.GetOrdersStream() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHTTPApi_GetExecOrdersStream(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name   string
		fields fields
		want   chan []byte
		want1  chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, got1 := h.GetExecOrdersStream()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetExecOrdersStream() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HTTPApi.GetExecOrdersStream() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHTTPApi_GetTradeStream(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chan []byte
		want1  chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, got1 := h.GetTradeStream(tt.args.symbol)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetTradeStream() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HTTPApi.GetTradeStream() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHTTPApi_GetQuoteStream(t *testing.T) {
	type fields struct {
		Auth       Auth
		httpClient *http.Client
		baseAPIURL string
		version    string
	}
	tests := []struct {
		name   string
		fields fields
		want   chan []byte
		want1  chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HTTPApi{
				Auth:       tt.fields.Auth,
				httpClient: tt.fields.httpClient,
				baseAPIURL: tt.fields.baseAPIURL,
				version:    tt.fields.version,
			}
			got, got1 := h.GetQuoteStream()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPApi.GetQuoteStream() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HTTPApi.GetQuoteStream() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
