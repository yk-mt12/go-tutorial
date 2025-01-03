package main

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Even number 10", 10, true},
		{"Odd number 7", 7, false},
		{"Odd number 1", 1, false},
		{"Even number 0", 0, true},
		{"Even number -2", -2, true},
		{"Odd number -5", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.n); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
