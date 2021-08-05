package cu

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const RatesUrl = "https://open.er-api.com/v6/latest/USD"

type Response struct {
	BaseCode    string  `json:"base_code"`
	Time        string  `json:"time_next_update_utc"`
	Rates       Rates
}

type Rates struct {
	USD  string  `json:"USD"`
	CNY  string  `json:"CNY"`
}

func Fetch() (*Rates, error) {
	resp, err := http.Get(RatesUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	// FIXME: discuss with pylon
	if err := json.Unmarshal(respData, &res); err != nil {
		log.Fatal("ooo! an error occurred!")
	}

	rates := Rates{
		USD: res.Rates.USD,
		CNY: res.Rates.CNY,
	}
	return &rates, nil
}
