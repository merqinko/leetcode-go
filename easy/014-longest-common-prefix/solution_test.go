package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		// Happy Cases
		{"common prefix", []string{"flower", "flow", "flight"}, "fl"},
		{"no common prefix", []string{"dog", "racecar", "car"}, ""},
		{"long common prefix", []string{"interspecies", "interstellar", "interstate"}, "inters"},

		// Edge Cases
		{"single string", []string{"single"}, "single"},
		{"empty strings", []string{"", ""}, ""},
		{"one character", []string{"a"}, "a"},
		{"empty slice", []string{}, ""},
		{"all same", []string{"abc", "abc", "abc"}, "abc"},
		{"prefix longer than some strings", []string{"ab", "abc"}, "ab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.expected)
			}
		})
	}
}
