package gemini

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_GetAccountDetails(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		apiSecret  string
		HTTPClient *http.Client
	}
	type args struct {
		ctx context.Context
		i   *GetAccountDetailsInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
				i: &GetAccountDetailsInput{
					ShortName: "primary",
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

			_, err := c.GetAccountDetails(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAccountDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// TODO(igaskin): uncomment this when response is mocked
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Client.GetAccountDetails() = %v, want %v", got, tt.want)
			// }
		})
	}
}
