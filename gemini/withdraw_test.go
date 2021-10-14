package gemini

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Withdraw(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		apiSecret  string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
		i   *WithdrawInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WithdrawResponse
		wantErr bool
	}{
		{
			name: "withdraw eth",
			args: args{
				ctx: context.Background(),
				i: &WithdrawInput{
					Address:  "",
					Amount:   "0.001",
					Currency: "eth",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.Withdraw(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
