package main

import (
	"context"
	"fmt"

	"github.com/igaskin/go-gemini/gemini"
)

func main() {
	client := gemini.NewClient()
	ctx := context.Background()

	resp, err := client.NewOrder(ctx, &gemini.NewOrderInput{
		Symbol:    "ethusd",
		Side:      "BUY",
		Amount:    ".001",
		OrderType: "exchange limit",
		Account:   "primary",
		Price:     "9999999",
		Options:   []string{"immediate-or-cancel"},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", resp)
}
