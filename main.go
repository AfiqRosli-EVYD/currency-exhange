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
	USD float64 `json:"USD"`
}

func exchangingCurrency(c *gin.Context) {
	var tc TargetCurrency
	var rc ResultCurrency
	rate := getExchangeRate()

	if err := c.BindJSON(&tc); err != nil {
		return
	}

	exchangedTotal := tc.Amount * rate

	rc.Currency = tc.ToCurrency
	rc.Amount = fmt.Sprintf("%.2f", exchangedTotal)
	rc.ExchangeRate = rate

	c.IndentedJSON(http.StatusCreated, rc)
}

func getExchangeRate() float64 {
	// TODO: Get API key from the Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	currencyScoopAPIKey := os.Getenv("CURRENCY_SCOOP_API")
	apiURL := "https://api.currencyscoop.com/v1/latest?base=BND&symbols=USD&api_key=" + currencyScoopAPIKey

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

	return resObj.Response.Rates.USD
}

func main() {
	router := gin.Default()
	router.POST("/currency_exchange", exchangingCurrency)

	router.Run("localhost:8080")
}
