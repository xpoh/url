package butty

import (
	"testing"
)

func TestNewButtyService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test New Butty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewButtyService()
			if GetService() == nil {
				t.Error("NewButtyService() Service == nil")
			}
		})
	}
}
