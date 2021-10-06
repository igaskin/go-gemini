package gemini

import (
	"context"
	"encoding/json"
	"net/http"
)

type TickerInput struct {
	Request string
	Ticker  string
}

type TickerResponse struct {
	Ask    string `json:"ask"`
	Bid    string `json:"bid"`
	Last   string `json:"last"`
	Volume Volume `json:"volume"`
}
type Volume struct {
	Btc       string `json:"BTC"`
	Usd       string `json:"USD"`
	Timestamp int64  `json:"timestamp"`
}

// Ticker get ticker information about trading symbols
func (c *Client) Ticker(ctx context.Context, i *TickerInput) (*TickerResponse, error) {
	var response *TickerResponse
	i.Request = "/v1/pubticker/"

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		c.BaseURL+i.Request+i.Ticker,
		nil,
	)
	if err != nil {
		return nil, err
	}
	rawResponse, err := c.doPublicRequest(req)
	if rawResponse == nil || err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, err
	}
	return response, nil
}
