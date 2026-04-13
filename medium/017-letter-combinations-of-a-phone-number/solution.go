package main

import "fmt"

// ============================================================
// SOLUTION: Backtracking (DFS)
// ============================================================
// Approach:
// 1. Map digits to characters (like a phone keypad)
// 2. Use DFS to generate all combinations
// 3. Backtrack: add char, recurse, remove char, try next
// Time: O(4^n) - n = number of digits, 4 = max chars per digit
// Space: O(n) - recursion depth = length of input digits

var phoneMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result // Return empty slice for empty input
	}

	// Start DFS with empty path
	var backtrack func(path string, index int)
	backtrack = func(path string, index int) {
		// Base case: path length equals digits length
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}

		// Get letters for current digit
		digit := digits[index]
		letters := phoneMap[digit]

		// Try each letter for this digit
		for i := 0; i < len(letters); i++ {
			backtrack(path+string(letters[i]), index+1) // Recurse
		}
	}

	backtrack("", 0) // Start with empty path at index 0
	return result
}

func main() {
	tests := []struct {
		digits   string
		expected []string
	}{
		{digits: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", expected: []string{}},
		{digits: "2", expected: []string{"a", "b", "c"}},
		{digits: "7", expected: []string{"p", "q", "r", "s"}},
	}

	fmt.Println("=== Letter Combinations of a Phone Number ===")
	for i, tt := range tests {
		result := letterCombinations(tt.digits)
		fmt.Printf("Test %d: Input: %q\n", i+1, tt.digits)
		fmt.Printf("  Result: %v\n", result)
		fmt.Printf("  Expected: %v\n\n", tt.expected)
	}
}
