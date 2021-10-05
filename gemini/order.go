package gemini

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type NewOrderInput struct {
	Request       string   `json:"request"`
	Nonce         int64    `json:"nonce"`
	ClientOrderID string   `json:"client_order_id"`
	Symbol        string   `json:"symbol"`
	Amount        string   `json:"amount"`
	MinAmount     string   `json:"min_amount"`
	Price         string   `json:"price"`
	Side          string   `json:"side"`
	OrderType     string   `json:"type"`
	Options       []string `json:"options"`
	StopPrice     string   `json:"stop_price"`
	Account       string   `json:"account"`
}

type NewOrderResponse struct {
	OrderID           string        `json:"order_id"`
	ID                string        `json:"id"`
	Symbol            string        `json:"symbol"`
	Exchange          string        `json:"exchange"`
	AvgExecutionPrice string        `json:"avg_execution_price"`
	Side              string        `json:"side"`
	Type              string        `json:"type"`
	Timestamp         string        `json:"timestamp"`
	Timestampms       int64         `json:"timestampms"`
	IsLive            bool          `json:"is_live"`
	IsCancelled       bool          `json:"is_cancelled"`
	IsHidden          bool          `json:"is_hidden"`
	WasForced         bool          `json:"was_forced"`
	ExecutedAmount    string        `json:"executed_amount"`
	Options           []interface{} `json:"options"`
	Price             string        `json:"price"`
	OriginalAmount    string        `json:"original_amount"`
	RemainingAmount   string        `json:"remaining_amount"`
}

// NewOrder submits a order to the exchange
func (c *Client) NewOrder(ctx context.Context, i *NewOrderInput) (*NewOrderResponse, error) {
	var response *NewOrderResponse
	i.Request = "/v1/order/new"
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
