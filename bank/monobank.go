package bank

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	API_URL     = "https://api.monobank.ua/bank/currency"
	MONO_FORMAT = "20060102"
	READ_FORMAT = "2006-01-02 15:04:05"
)

type ApiResponse struct {
	CurrencyCodeA    int32   `json:"currencyCodeA"`
	CurrencyCodeB    int32   `json:"currencyCodeB"`
	Date             int64   `json:"date"`
	RateBuy          float64 `json:"rateBuy"`
	RateSell         float64 `json:"rateSell"`
	RateCross        float64 `json:"rateCross"`
}
type R []ApiResponse

type Monobank struct {
	client *http.Client
}

func (mb *Monobank) CurrentExchangeRate() ExchangeRate {
	// today := time.Now()
	// dateStr := today.Format(DATE_FORMAT)

	response, err := mb.client.Get(API_URL)
	if err != nil {
		fmt.Println(err)
	}
	responseJson, _ := io.ReadAll(response.Body)

	data := []ApiResponse{}
	parseErr := json.Unmarshal(responseJson, &data)
	
	if parseErr != nil {
		fmt.Println(parseErr)
	} else {
		for _, item := range data {
			t := time.Unix(item.Date, 0)
			now := time.Now()
			if t.Format(MONO_FORMAT) != now.Format(MONO_FORMAT) {
				continue
			}
			
			if ParseISO_4217(item.CurrencyCodeA) == "" || ParseISO_4217(item.CurrencyCodeB) == "" {
				continue
			}
			
			fmt.Println("===================================")
			fmt.Println("cur A: ", ParseISO_4217(item.CurrencyCodeA))
			fmt.Println("cur B:", ParseISO_4217(item.CurrencyCodeB))
			fmt.Println("date: ", t.Format(READ_FORMAT))
			fmt.Println("buy rate: ", item.RateBuy)
			fmt.Println("sell rate: ", item.RateSell)
		}
	}
		
	rate := ExchangeRate{
		Buy:  1,
		Sell: 1,
	}
	return rate
}

func NewMonobank(client *http.Client) IBank {
	return &Monobank{
		client: client,
	}
}
