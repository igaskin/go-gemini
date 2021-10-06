package gemini

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	BaseURLV1          = "https://api.gemini.com"
	sandboxBaseURLV1   = "https://api.sandbox.gemini.com"
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
	// TODO(igaskin) implement heartbeatA
	// https://docs.gemini.com/rest-api/#require-heartbeat
	// Heartbeat bool
}

// NewClient new client wiht defaults
func NewClient() *Client {
	return NewClientFromConfig(Config{
		BaseURL: BaseURLV1,
	})
}

// NewClientFromConfig new client using custom configurations
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
	_, err = hash.Write([]byte(payload))
	if err != nil {
		return nil, err
	}
	geminiSignature := hex.EncodeToString(hash.Sum(nil))

	// Authenticated APIs do not submit their payload as POSTed data,
	// but instead put it in the X-GEMINI-PAYLOAD header
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-GEMINI-APIKEY", c.apiKey)
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

func (c *Client) doPublicRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
