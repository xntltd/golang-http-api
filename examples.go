package httplib

import "fmt"

const baseAPIURL = "https://api-demo.exante.eu"

func test() {
	h := NewAPI(baseAPIURL,
		"2.0",
		"56ef81eb-c530-4f9d-936c-a6cefb72d772",
		"85411166-368a-434b-8b4c-4d642fcd4493",
		"xpeNrMLpgG3JEoWaj84+CBxvm0UpSwyd",
		30, "", "",
	)
	data, err := h.GetExchanges()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(data)
	}

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
}
