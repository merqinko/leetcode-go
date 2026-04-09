package main

import (
	"fmt"
	"strings"
)

// ============================================================
// SOLUTION: Horizontal Scanning
// ============================================================
// Approach: Compare prefix of first two strings, then compare
// that result with third, and so on.
// Time: O(S) where S = total characters in all strings
// Space: O(1) excluding result string

func longestCommonPrefix(strs []string) string {
	// Edge case: empty slice
	if len(strs) == 0 {
		return ""
	}

	// Start with first string as the prefix
	prefix := strs[0]

	// Compare prefix with each subsequent string
	for i := 1; i < len(strs); i++ {
		// Keep shortening prefix until it matches start of strs[i]
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1] // Remove last character

			// If prefix becomes empty, no common prefix exists
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	// Test cases
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"single"}, "single"},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
	}

	fmt.Println("=== Longest Common Prefix ===")
	for i, tt := range tests {
		result := longestCommonPrefix(tt.input)
		status := "✅"
		if result != tt.expected {
			status = "❌"
		}
		fmt.Printf("%s Test %d: %v → %q (expected %q)\n", status, i+1, tt.input, result, tt.expected)
	}
}
