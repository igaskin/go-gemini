package gemini

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WrapOrder(t *testing.T) {
	type fields struct {
		BaseURL    string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
		i   *WrapOrderInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WrapOrderResponse
		wantErr bool
	}{
		{
			name: "sell GUSD",
			args: args{
				ctx: context.Background(),
				i: &WrapOrderInput{
					Side:    "sell",
					Amount:  "1",
					Symbol:  "gusdusd",
					Account: "primary",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			// for some reason wrap orders
			// do not work on the sandbox environment
			c.BaseURL = sandboxBaseURLV1
			got, err := c.WrapOrder(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.WrapOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WrapOrder() = %v, want %v", got, tt.want)
			}
			assert.NotNil(t, got)
		})
	}
}
