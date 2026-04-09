package main

import "testing"

// ============================================================
// TEST 1: Testing the Switch Solution (Solution 1)
// ============================================================

func TestRomanToInt1(t *testing.T) {
	// Table-Driven Tests | Tes Berbasis Tabel
	tests := []struct {
		name     string // Test name | Nama tes
		roman    string // Input Roman | Input Romawi
		expected int    // Expected output | Output yang diharapkan
	}{
		// Happy Cases | Kasus Normal
		{"simple three", "III", 3},
		{"subtraction four", "IV", 4},
		{"subtraction nine", "IX", 9},
		{"medium number", "LVIII", 58},
		{"large number", "MCMXCIV", 1994},

		// Edge Cases | Kasus Khusus
		{"empty string", "", 0}, // What if input is empty? | Bagaimana jika input kosong?
		{"single I", "I", 1},    // Smallest number | Angka terkecil
	}

	// Loop through test cases | Loop melalui kasus tes
	for _, tt := range tests {
		// Create a sub-test | Buat sub-tes
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt1(tt.roman) // Call function | Panggil fungsi
			if got != tt.expected {
				// Report error | Laporkan error
				t.Errorf("romanToInt1(%q) = %d, want %d", tt.roman, got, tt.expected)
			}
		})
	}
}

// ============================================================
// TEST 2: Testing the Map Solution (Solution 2)
// ============================================================

func TestRomanToInt2(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{"one", "I", 1},
		{"five", "V", 5},
		{"complex", "MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt2(tt.roman)
			if got != tt.expected {
				t.Errorf("romanToInt2(%q) = %d, want %d", tt.roman, got, tt.expected)
			}
		})
	}
}
