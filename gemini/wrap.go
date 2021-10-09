package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type WrapOrderInput struct {
	Request       string `json:"request"`
	Nonce         int64  `json:"nonce"`
	Account       string `json:"account"`
	Amount        string `json:"amount"`
	Side          string `json:"side"`
	ClientOrderID string `json:"client_order_id"`
	symbol        string
}

type WrapOrderResponse struct {
	OrderID            int    `json:"orderId"`
	Pair               string `json:"pair"`
	Price              string `json:"price"`
	PriceCurrency      string `json:"priceCurrency"`
	Side               string `json:"side"`
	Quantity           string `json:"quantity"`
	QuantityCurrency   string `json:"quantityCurrency"`
	TotalSpend         string `json:"totalSpend"`
	TotalSpendCurrency string `json:"totalSpendCurrency"`
	Fee                string `json:"fee"`
	FeeCurrency        string `json:"feeCurrency"`
	DepositFee         string `json:"depositFee"`
	DepositFeeCurrency string `json:"depositFeeCurrency"`
}

func (c *Client) WrapOrder(ctx context.Context, i *WrapOrderInput) (*WrapOrderResponse, error) {
	var response *WrapOrderResponse
	i.Request = "/v1/wrap/" + i.symbol
	if i.Nonce == 0 {
		i.Nonce = time.Now().UnixNano() / 1000000
	}

	json_data, err := json.Marshal(&i)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.BaseURL+i.Request,
		strings.NewReader(string(json_data)),
	)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.doRequest(req)
	if rawResponse == nil || err != nil {
		return nil, err
	}
	fmt.Printf("%+s\n", rawResponse)
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", response)
	return response, nil
}
