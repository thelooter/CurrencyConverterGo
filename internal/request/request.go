package request

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"thelooter.de/currencyconverter/datastructures"
	"thelooter.de/currencyconverter/globals"
)

func createRequestURL(startingCurrency string, targetCurrency string) string {
	requestURL := globals.Config.BaseURL + globals.Config.APIKey
	requestURL += "/pair/"
	requestURL += startingCurrency
	requestURL += "/"
	requestURL += targetCurrency

	return requestURL
}

func GetRate(startingCurrency string, targetCurrency string) float64 {
	url := createRequestURL(startingCurrency, targetCurrency)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var data datastructures.Request

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data.ConversionRate
}
