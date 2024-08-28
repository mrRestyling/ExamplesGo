package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"positive palindrome", 1001, true},
		{"positive not palindrome", 1234, false},
		{"negative", -1001, false},
		{"zero", 0, true},
		{"single digit", 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
