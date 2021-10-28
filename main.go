package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func exchangingCurrency(c *gin.Context) {
	var tc TargetCurrency
	var rc ResultCurrency
	rate := 0.74188

	if err := c.BindJSON(&tc); err != nil {
		return
	}

	exchangedTotal := tc.Amount * rate

	rc.Currency = tc.ToCurrency
	rc.Amount = fmt.Sprintf("%.2f", exchangedTotal)
	rc.ExchangeRate = rate

	c.IndentedJSON(http.StatusCreated, rc)
}

func main() {
	router := gin.Default()
	router.POST("/currency_exchange", exchangingCurrency)

	router.Run("localhost:8080")
}
