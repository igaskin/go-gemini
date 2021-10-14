package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type WithdrawInput struct {
	Request  string `json:"request"`
	Nonce    int64  `json:"nonce"`
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Currency string `json:"-"`
}

type WithdrawResponse struct {
	Address      string `json:"address"`
	Amount       string `json:"amount"`
	WithdrawalID string `json:"withdrawalId"`
	Message      string `json:"message"`
}

// Withdraw to external crypto wallet, requires "Fund management role" and crypto address whitelist
func (c *Client) Withdraw(ctx context.Context, i *WithdrawInput) (*WithdrawResponse, error) {
	var response *WithdrawResponse
	i.Request = "/v1/withdraw/" + i.Currency
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
