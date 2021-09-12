package extstr

import (
	"testing"
)

func TestIsGoKeywords(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"var",
			"var",
			true,
		},
		{
			"aa",
			"aa",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGoKeywords(tt.args); got != tt.want {
				t.Errorf("IsGoKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGoInternalType(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"string",
			"string",
			true,
		},
		{
			"aa",
			"aa",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGoInternalType(tt.args); got != tt.want {
				t.Errorf("IsGoInternalType() = %v, want %v", got, tt.want)
			}
		})
	}
}
