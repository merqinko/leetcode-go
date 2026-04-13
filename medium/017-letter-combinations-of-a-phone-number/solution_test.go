package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		// Happy Cases
		{"two digits", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"single digit", "2", []string{"a", "b", "c"}},
		{"digit with 4 letters", "7", []string{"p", "q", "r", "s"}},
		{"four digits", "234", []string{
			"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
			"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
			"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
		}},

		// Edge Cases
		{"empty string", "", []string{}},
		{"digit 1 (no mapping)", "1", []string{}},
		{"digit 0 (no mapping)", "0", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			
			// Sort both slices for consistent comparison
			sort.Strings(got)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.expected)
			}
		})
	}
}
