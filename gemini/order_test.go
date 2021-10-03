package gemini

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_NewOrder(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		apiSecret  string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
		i   *NewOrderInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sell btc",
			args: args{
				ctx: context.Background(),
				i: &NewOrderInput{
					Symbol:    "btcusd",
					Side:      "SELL",
					Amount:    ".1",
					OrderType: "exchange limit",
					Account:   "primary",
					Price:     "48126.75",
				},
			},
			wantErr: false,
		},
		{
			name: "buy btc limit",
			args: args{
				ctx: context.Background(),
				i: &NewOrderInput{
					Symbol:    "btcusd",
					Side:      "BUY",
					Amount:    ".1",
					OrderType: "exchange limit",
					Account:   "primary",
					Price:     "47964.5",
				},
			},
			wantErr: false,
		},
		{
			name: "buy btc market order fewest satoshis",
			args: args{
				ctx: context.Background(),
				i: &NewOrderInput{
					Symbol:    "btcusd",
					Side:      "BUY",
					Amount:    ".00001",
					OrderType: "exchange limit",
					Account:   "primary",
					Price:     "9999999",
					Options:   []string{"immediate-or-cancel"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO(igaskin): mock the response instead of hitting actual sandbox
			c := NewClient()
			c.BaseURL = sandboxBaseURLV1

			if _, err := c.NewOrder(tt.args.ctx, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Client.NewOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
