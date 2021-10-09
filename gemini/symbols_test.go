package gemini

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Symbols(t *testing.T) {
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
		want    *TickerResponse
		wantErr bool
	}{
		{
			name: "symbols",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.Symbols(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Symbols() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}
