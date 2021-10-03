package gemini

import (
	"net/http"
	"reflect"
	"testing"
	"time"
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
					Timeout: 3 * time.Second,
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
