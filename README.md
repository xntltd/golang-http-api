# golang-http-api

Golang library for XNT Ltd. HTTP API service provides access to actual endpoints (https://api-live.exante.eu/api-docs/)

### Building and installation

To install golang-http-api library, use ```go get```:
```
$ go get github.com/xntltd/golang-http-api
```

### Staying up to date

To update to the latest version, use
```
$ go get -u github.com/xntltd/golang-http-api
```

## Developing golang-http-api

If you wish to work on golang-http-api itself, you will first need Go installed and configured on your machine (version 1.14+ is preferred, but the minimum required version is 1.8).

Next, using Git, clone the repository via git clone 
```
git clone https://github.com/xntltd/golang-http-api.git
```

## Basic usage

First of all you need to get ApplicationID, ClientID and SharedKey from sandbox or production enviroment. And choose one of the authorization methods. It could be BasicAuth or JWT token.

Then instantiate API with NewAPI constructor
```
h := NewAPI(baseAPIURL,
		"2.0",
		"56ef81eb-c534-499d-936c-a6cefb72d772",
		"8541316993680-124b-8b4c-4d642fcd4493",
		"xpeNrMLpgG3JEoWaj84",
		30, "", "",
	)
```
Then methods can be called. Try with GetExchanges for example:
```
data, err := h.GetExchanges()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(data)
	}
```
Or order initialization:
```
	o := OrderSentTypeV1{
		Side:           "buy",
		Duration:       "fill_or_kill",
		Quantity:       "10",
		OrderType:      "market",
		OcoGroup:       "5",
		Account:        "5",
		Instrument:     "Test",
		IfDoneParentID: "1",
	}

	err = h.PlaceOrderV1(&o)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("OK")
	}
```
Try to get account summary:
```
	accSum, err := h.GetAccountSummary("DEW8032.001", "USD",
		GetAccountSummaryPayload{
			DatetimePayload: DatetimePayload{
				Datetime: "2020-10-10",
			},
		})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%+v\n", accSum)
	}
```
See more code in ```examples.go``` file
All HTTP API methods declared in ```api.go``` file