package frankfurter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BASE_URL = "https://api.frankfurter.app"

func Get(url string) ([]byte, error) {
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		return nil, httpGetError
	}
	defer response.Body.Close()

	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		return nil, readBodyError
	}

	return body, nil
}

type LatestResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

func GetLatest() (LatestResponse, error) {
	var url string = fmt.Sprintf("%s/latest", BASE_URL)

	body, getError := Get(url)
	if getError != nil {
		return LatestResponse{}, getError
	}

	var latestResponse LatestResponse
	jsonUnmarshalError := json.Unmarshal(body, &latestResponse)
	if jsonUnmarshalError != nil {
		return LatestResponse{}, jsonUnmarshalError
	}

	return latestResponse, nil
}

func GetCurrencies() (map[string]string, error) {
	var url string = fmt.Sprintf("%s/currencies", BASE_URL)

	body, getError := Get(url)
	if getError != nil {
		return map[string]string{}, getError
	}

	var currenciesResponse map[string]string
	jsonUnmarshalError := json.Unmarshal(body, &currenciesResponse)
	if jsonUnmarshalError != nil {
		return map[string]string{}, jsonUnmarshalError
	}

	return currenciesResponse, nil
}
