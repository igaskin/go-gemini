# go-gemini
unoffcial go client for Gemini

[![GoDoc](https://godoc.org/github.com/igaskin/go-gemini?status.svg)](https://pkg.go.dev/github.com/igaskin/go-gemini)


## Features
* Authentication
* Submit Buy and Sell Orders

## Requirements
* Go >= 1.15

## Installation

```bash
go get github.com/igaskin/go-gemini
```

## Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/igaskin/go-gemini/gemini"
)

func main() {
    // set environment variables
    // GEMINI_API_KEY
    // GEMINI_API_SECRET
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
```