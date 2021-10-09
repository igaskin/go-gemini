package gemini

import (
	"context"
	"encoding/json"
	"net/http"
)

type SymbolsResponse []string

type SymbolDetailsInput struct {
	Ticker string
}

type SymbolDetailsResponse struct {
	Symbol         string  `json:"symbol"`
	BaseCurrency   string  `json:"base_currency"`
	QuoteCurrency  string  `json:"quote_currency"`
	TickSize       float64 `json:"tick_size"`
	QuoteIncrement float64 `json:"quote_increment"`
	MinOrderSize   string  `json:"min_order_size"`
	Status         string  `json:"status"`
	WrapEnabled    bool    `json:"wrap_enabled"`
}

// Symbols return a list of supported symbols
func (c *Client) Symbols(ctx context.Context) (*SymbolsResponse, error) {
	var response *SymbolsResponse
	request := "/v1/symbols"

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		c.BaseURL+request,
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

func (c *Client) SymbolDetails(ctx context.Context, i *SymbolDetailsInput) (*SymbolDetailsResponse, error) {
	var response *SymbolDetailsResponse
	request := "/v1/symbols/details"

	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		c.BaseURL+request+i.Ticker,
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
