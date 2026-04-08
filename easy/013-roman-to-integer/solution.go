package main

import "fmt"

// ============================================================
// SOLUTION 1: Switch Statement (RECOMMENDED - Faster)
// ============================================================
// Why better? No map allocation, direct comparison, faster lookup
// Time: O(n) | Space: O(1)

// value1 returns the integer value of a Roman numeral character
// Using switch: Go compiles this to efficient jump table
func value1(r byte) int {
	switch r { // Check the character
	case 'I':
		return 1 // I = 1
	case 'V':
		return 5 // V = 5
	case 'X':
		return 10 // X = 10
	case 'L':
		return 50 // L = 50
	case 'C':
		return 100 // C = 100
	case 'D':
		return 500 // D = 500
	case 'M':
		return 1000 // M = 1000
	default:
		return 0 // Invalid character (shouldn't happen)
	}
}

// romanToInt1 converts Roman numeral to integer using switch
// Logic: Loop RIGHT to LEFT
// - If current < previous → subtract (e.g., IV = 5-1 = 4)
// - If current >= previous → add (e.g., VI = 5+1 = 6)
func romanToInt1(s string) int {
	result := 0 // Accumulate the total here
	prev := 0   // Store the value of the previous (right-side) character

	// Loop backwards: from last character to first
	// Example: "MCMXCIV" → process V, then I, then C, etc.
	for i := len(s) - 1; i >= 0; i-- {
		curr := value1(s[i]) // Get integer value of current Roman character

		// If current is less than previous, it means subtraction case
		// Examples: IV (4), IX (9), XL (40), XC (90), CD (400), CM (900)
		if curr < prev {
			result -= curr // Subtract: e.g., IV → result = 5 - 1 = 4
		} else {
			result += curr // Add: e.g., VI → result = 5 + 1 = 6
		}

		prev = curr // Save current value for next iteration comparison
	}

	return result // Final converted integer
}

// ============================================================
// SOLUTION 2: Map (More Readable, Slightly Slower)
// ============================================================
// Why use it? Easier to understand, easier to add new symbols
// Time: O(n) | Space: O(1) (map is fixed size: 7 entries)

// romanMap stores all Roman numeral → integer mappings
// Created once as package-level variable (not recreated every call)
var romanMap = map[byte]int{
	'I': 1,    // I = 1
	'V': 5,    // V = 5
	'X': 10,   // X = 10
	'L': 50,   // L = 50
	'C': 100,  // C = 100
	'D': 500,  // D = 500
	'M': 1000, // M = 1000
}

// romanToInt2 converts Roman numeral to integer using map lookup
// Same logic as Solution 1, but uses map instead of switch
func romanToInt2(s string) int {
	result := 0 // Accumulate the total here
	prev := 0   // Store the value of the previous (right-side) character

	// Loop backwards: from last character to first
	for i := len(s) - 1; i >= 0; i-- {
		curr := romanMap[s[i]] // Lookup value in map

		// Subtraction case: e.g., IV, IX, XL, XC, CD, CM
		if curr < prev {
			result -= curr // Subtract from total
		} else {
			result += curr // Add to total
		}

		prev = curr // Update previous for next iteration
	}

	return result // Final converted integer
}

func main() {
	// Test cases for both solutions
	tests := []struct {
		roman    string
		expected int
	}{
		{"III", 3},        // 1+1+1 = 3
		{"IV", 4},         // 5-1 = 4 (subtraction case)
		{"IX", 9},         // 10-1 = 9 (subtraction case)
		{"LVIII", 58},     // 50+5+1+1+1 = 58
		{"MCMXCIV", 1994}, // 1000+900+90+4 = 1994
		{"XXVII", 27},     // 10+10+5+1+1 = 27
	}

	fmt.Println("=== SOLUTION 1 (Switch) ===")
	for _, tt := range tests {
		result := romanToInt1(tt.roman)
		status := "✅"
		if result != tt.expected {
			status = "❌"
		}
		fmt.Printf("%s %s → %d (expected %d)\n", status, tt.roman, result, tt.expected)
	}

	fmt.Println("\n=== SOLUTION 2 (Map) ===")
	for _, tt := range tests {
		result := romanToInt2(tt.roman)
		status := "✅"
		if result != tt.expected {
			status = "❌"
		}
		fmt.Printf("%s %s → %d (expected %d)\n", status, tt.roman, result, tt.expected)
	}
}
