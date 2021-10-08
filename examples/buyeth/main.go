package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/igaskin/go-gemini/gemini"
)

const (
	sandboxBaseURLV1 = "https://api.sandbox.gemini.com"
)

func main() {
	client := gemini.NewClient()
	client.BaseURL = sandboxBaseURLV1
	ctx := context.Background()
	amountUSD := 200.0
	maxRetry := 5
	retry := 0

	for {
		ticker, err := client.Ticker(ctx, &gemini.TickerInput{
			Ticker: "ethusd",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", ticker)

		// The ask price refers to the lowest price a seller will accept
		ask, _ := strconv.ParseFloat(ticker.Ask, 64)
		amount := amountUSD / ask

		order, err := client.NewOrder(ctx, &gemini.NewOrderInput{
			Symbol:    "ethusd",
			Side:      "BUY",
			Amount:    fmt.Sprintf("%f", amount),
			OrderType: "exchange limit",
			Account:   "primary",
			Price:     fmt.Sprintf("%.2f", ask),
			Options:   []string{"immediate-or-cancel"},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", order)

		// break if order is filled or hit retry limit
		if !order.IsCancelled || retry > maxRetry {
			break
		}
		retry += 1
	}

}
