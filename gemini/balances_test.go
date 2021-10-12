package gemini

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Balances(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		apiSecret  string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BalancesResponse
		wantErr bool
	}{
		{
			name:    "get balance",
			wantErr: false,
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			c.BaseURL = sandboxBaseURLV1
			got, err := c.Balances(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Balances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			// TODO move this to a seperate test
			b, err := c.BalanceSymbol(tt.args.ctx, got, "BTC")
			assert.Nil(t, err)
			assert.NotNil(t, b)
			fmt.Printf("%+v", b)
		})
	}
}
