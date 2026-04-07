package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{"positive palindrome", 121, true},
		{"negative number", -121, false},
		{"ends with zero", 10, false},
		{"zero", 0, true},
		{"single digit", 5, true},
		{"multi digit palindrome", 12321, true},
		{"non-palindrome", 12345, false},
		{"palindrome with even digits", 1221, true},
		{"large palindrome", 123454321, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, expected %v", tt.x, got, tt.expected)
			}
		})
	}
}
