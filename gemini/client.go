package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	BaseURLV1          = "https://api.gemini.com/v1/"
	sandboxBaseURLV1   = "https://api.sandbox.gemini.com/v1/"
	defaultHTTPTimeout = 3 * time.Second
)

type Client struct {
	BaseURL    string
	apiKey     string
	apiSecret  string
	HTTPClient *http.Client
}

type Config struct {
	APIKey    string
	APISecret string
	BaseURL   string
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

func NewClient() *Client {
	return NewClientFromConfig(Config{
		BaseURL: BaseURLV1,
	})
}

func NewClientFromConfig(c Config) *Client {
	if c.BaseURL == "" {
		c.BaseURL = BaseURLV1
	}
	if c.APIKey == "" {
		c.APIKey = os.Getenv("GEMINI_API_KEY")
	}
	if c.APISecret == "" {
		c.APISecret = os.Getenv("GEMINI_API_SECRET")
	}

	httpClient := &http.Client{Timeout: defaultHTTPTimeout}

	return &Client{
		BaseURL:    c.BaseURL,
		apiKey:     c.APIKey,
		apiSecret:  c.APISecret,
		HTTPClient: httpClient,
	}
}

// TODO(igaskin): move this to a separate file
func (c *Client) GetAccountDetails(ctx context.Context) (*GetAccountDetailsResponse, error) {
	var response *GetAccountDetailsResponse
	// unix milliseconds as nonce
	nonce := time.Now().UnixNano() / 1000000

	values := map[string]string{
		"request": "/v1/account",
		"account": "primary",
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
	fmt.Println(string(rawResponse))
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	// nil out the body
	// https://docs.sandbox.gemini.com/rest-api/#private-api-invocation#payload
	req.Body = nil
	payload := base64.StdEncoding.EncodeToString(body)

	// generate hmac signature from secret
	hash := hmac.New(sha512.New384, []byte(c.apiSecret))

	// Write Data to it
	_, err = hash.Write([]byte(payload))
	if err != nil {
		return nil, err
	}
	geminiSignature := hex.EncodeToString(hash.Sum(nil))

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-GEMINI-APIKEY", c.apiKey)
	// Authenticated APIs do not submit their payload as POSTed data,
	// but instead put it in the X-GEMINI-PAYLOAD header
	req.Header.Set("X-GEMINI-PAYLOAD", payload)
	req.Header.Set("X-GEMINI-SIGNATURE", geminiSignature)
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
