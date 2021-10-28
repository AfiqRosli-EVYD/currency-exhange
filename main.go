package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeCurrency struct {
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Amount       float64 `json:"amount"`
}

func exchangingCurrency(c *gin.Context) {
	var ec ExchangeCurrency
	rate := 0.74

	if err := c.BindJSON(&ec); err != nil {
		return
	}

	exchangedTotal := ec.Amount * rate

	c.IndentedJSON(http.StatusCreated, exchangedTotal)
}

func main() {
	router := gin.Default()
	router.POST("/currency_exchange", exchangingCurrency)

	router.Run("localhost:8080")
}
