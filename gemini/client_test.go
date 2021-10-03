package client

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "valid",
			want: &Client{
				BaseURL: "https://api.gemini.com/v1/",
				HTTPClient: &http.Client{
					Timeout: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountDetails(t *testing.T) {
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
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO(igaskin): mock the response instead of hitting actual sandbox
			c := NewClient()
			c.BaseURL = sandboxBaseURLV1

			_, err := c.GetAccountDetails(tt.args.ctx)
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
