package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TargetCurrency struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Amount       float64 `json:"amount"`
}

type ResultCurrency struct {
	Currency     string  `json:"currency"`
	Amount       string  `json:"amount"`
	ExchangeRate float64 `json:"exchange_rate"`
}

type ExchangeRate struct {
	Response Response `json:"response"`
}

type Response struct {
	Rates Rates `json:"rates"`
}

type Rates struct {
	BND float64 `json:"BND"`
	SGD float64 `json:"SGD"`
	USD float64 `json:"USD"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func exchangingCurrency(c *gin.Context) {
	var targetCurrency TargetCurrency
	var resultCurrency ResultCurrency

	if err := c.BindJSON(&targetCurrency); err != nil {
		return
	}

	if !isSupportedExchangeCurrency(targetCurrency.ToCurrency) {
		var em ErrorMessage
		em.Message = "Supported currencies are BND, SGD, USD"

		c.IndentedJSON(http.StatusForbidden, em)

		return
	}

	rate := getExchangeRate(targetCurrency.FromCurrency, targetCurrency.ToCurrency)
	exchangedTotal := targetCurrency.Amount * rate

	resultCurrency.Currency = targetCurrency.ToCurrency
	resultCurrency.Amount = fmt.Sprintf("%.2f", exchangedTotal)
	resultCurrency.ExchangeRate = rate

	c.IndentedJSON(http.StatusCreated, resultCurrency)
}

func getExchangeRate(from string, to string) float64 {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	currencyScoopAPIKey := os.Getenv("CURRENCY_SCOOP_API")
	apiURL := fmt.Sprintf("https://api.currencyscoop.com/v1/latest?base=%s&symbols=%s&api_key=%s", from, to, currencyScoopAPIKey)

	res, err := http.Get(apiURL)
	if err != nil {
		fmt.Print(err.Error())
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resObj ExchangeRate
	json.Unmarshal(resData, &resObj)

	if to == "BND" {
		return resObj.Response.Rates.BND
	} else if to == "SGD" {
		return resObj.Response.Rates.SGD
	} else {
		return resObj.Response.Rates.USD
	}
}

func isSupportedExchangeCurrency(to string) bool {
	currencies := [3]string{"BND", "SGD", "USD"}

	for _, currency := range currencies {
		if currency == to {
			return true
		}
	}

	return false
}

func main() {
	router := gin.Default()
	router.POST("/currency_exchange", exchangingCurrency)

	router.Run("localhost:8080")
}
