package gemini

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type GetAccountDetailsInput struct {
	Request string `json:"request"`
	Account string `json:"account"`
	Nonce   int64  `json:"nonce"`
}

type GetAccountDetailsResponse struct {
	Account           Account `json:"account"`
	Users             []Users `json:"users"`
	MemoReferenceCode string  `json:"memo_reference_code"`
}

type Account struct {
	AccountName       string `json:"accountName"`
	ShortName         string `json:"shortName"`
	Type              string `json:"type"`
	Created           string `json:"created"`
	VerificationToken string `json:"verificationToken"`
}

type Users struct {
	Name        string    `json:"name"`
	LastSignIn  time.Time `json:"lastSignIn"`
	Status      string    `json:"status"`
	CountryCode string    `json:"countryCode"`
	IsVerified  bool      `json:"isVerified"`
}

// GetAccountDetails returns gemini account information
func (c *Client) GetAccountDetails(ctx context.Context, i *GetAccountDetailsInput) (*GetAccountDetailsResponse, error) {
	var response *GetAccountDetailsResponse
	i.Request = "/v1/account"
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
