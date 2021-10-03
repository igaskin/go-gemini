package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GetAccountDetailsInput struct {
	ShortName string
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

func (c *Client) GetAccountDetails(ctx context.Context, i *GetAccountDetailsInput) (*GetAccountDetailsResponse, error) {
	var response *GetAccountDetailsResponse
	// unix milliseconds as nonce
	nonce := time.Now().UnixNano() / 1000000

	values := map[string]string{
		"request": "/v1/account",
		"account": i.ShortName,
		"nonce":   fmt.Sprintf("%d", nonce),
	}
	json_data, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.BaseURL+"account",
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
