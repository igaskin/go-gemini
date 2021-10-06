package gemini

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	type fields struct {
		BaseURL    string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
		i   *TickerInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TickerResponse
		wantErr bool
	}{
		{
			name: "btc ticker",
			args: args{
				ctx: context.Background(),
				i: &TickerInput{
					Ticker: "btcusd",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			_, err := c.Ticker(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Ticker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// TODO(igaskin) uncomment with mocked server response
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Client.Ticker() = %v, want %v", got, tt.want)
			// }
		})
	}
}
