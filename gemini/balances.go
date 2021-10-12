package gemini

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type BalancesResponse []struct {
	Type                   string `json:"type"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	Available              string `json:"available"`
	AvailableForWithdrawal string `json:"availableForWithdrawal"`
}

type Balance struct {
	Type                   string `json:"type"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	Available              string `json:"available"`
	AvailableForWithdrawal string `json:"availableForWithdrawal"`
}

type BalancesInput struct {
	Request string `json:"request"`
	Nonce   int64  `json:"nonce"`
}

func (c *Client) Balances(ctx context.Context) (*BalancesResponse, error) {
	var response *BalancesResponse
	var i BalancesInput
	i.Request = "/v1/balances"
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
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) BalanceSymbol(ctx context.Context, i *BalancesResponse, symbol string) (*Balance, error) {
	for _, v := range *i {
		if v.Currency == symbol {
			return &Balance{
				Type:                   v.Type,
				Currency:               v.Currency,
				Amount:                 v.Amount,
				Available:              v.Available,
				AvailableForWithdrawal: v.AvailableForWithdrawal,
			}, nil
		}
	}
	return nil, errors.New("no balance for symbol")
}
