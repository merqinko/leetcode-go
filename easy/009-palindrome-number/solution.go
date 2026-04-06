package main

import "fmt"

// ============================================================
// algorithm Palindrome Number (LeetCode #9)
// ============================================================
// Input:  An integer x
// Output: true if x is a palindrome, false otherwise
//
// Apa itu palindrome?
// Angka yang dibaca dari kiri ke kanan SAMA dengan dari kanan ke kiri.
// Contoh: 121 -> true, -121 -> false, 10 -> false
//
// Approach: Reverse the integer MATHEMATICALLY (no string conversion)
//
// Step-by-step logic:
// 1. If x is negative → return false (tanda minus bikin ga mungkin palindrome)
// 2. Simpan nilai asli x ke variabel original (karena x akan berubah saat di-loop)
// 3. Reverse x secara matematis:
//    a. x % 10     → ambil digit paling kanan (contoh: 121 % 10 = 1)
//    b. reversed = reversed * 10 + digit → geser reversed ke kiri, tambah digit baru
//    c. x /= 10    → buang digit paling kanan (contoh: 121 / 10 = 12)
// 4. Ulangi sampai x habis (x <= 0)
// 5. Bandingkan original == reversed → kalau sama, berarti palindrome
//
// Contoh: x = 121
//   Iterasi 1: digit=1, reversed=1,  x=12
//   Iterasi 2: digit=2, reversed=12, x=1
//   Iterasi 3: digit=1, reversed=121, x=0 → stop
//   original(121) == reversed(121) → true ✅

func isPalindrome(x int) bool {
	// Bilangan negatif bukan palindrome
	// Negative numbers can't be palindrome because of the "-" sign
	// Example: -121 → reads "121-" from right to left → not same
	if x < 0 {
		return false
	}

	// Simpan nilai asli x sebelum diubah
	// Store original value because x will be modified in the loop
	original := x

	// Variabel untuk menampung angka yang sudah dibalik
	// Variable to store the reversed number
	reversed := 0

	// Loop sampai x habis (semua digit sudah diproses)
	// Loop until all digits are processed (x becomes 0)
	for x > 0 {
		// Ambil digit paling kanan menggunakan modulo 10
		// Get the last digit using modulo 10
		// Example: 121 % 10 = 1
		digit := x % 10

		// Geser reversed ke kiri (kalikan 10), lalu tambah digit baru
		// Shift reversed left by multiplying 10, then add the new digit
		// Example: reversed=0 → 0*10+1=1 → 1*10+2=12 → 12*10+1=121
		reversed = reversed*10 + digit

		// Buang digit paling kanan dari x dengan integer division 10
		// Remove the last digit from x by integer division by 10
		// Example: 121 / 10 = 12 (integer division drops decimal)
		x /= 10
	}

	// Bandingkan angka asli dengan angka yang sudah dibalik
	// Compare original number with reversed number
	// Kalau sama → palindrome (true), kalau beda → bukan palindrome (false)
	return original == reversed
}

func main() {
	// Test case 1: Angka palindrome positif
	// 121 dibaca dari kiri atau kanan tetap 121
	fmt.Println(isPalindrome(121)) // Output: true

	// Test case 2: Angka negatif
	// -121 dibaca dari kanan jadi "121-" → bukan palindrome
	fmt.Println(isPalindrome(-121)) // Output: false

	// Test case 3: Angka diakhiri 0 (kecuali 0 sendiri)
	// 10 dibaca dari kanan jadi "01" → bukan palindrome
	fmt.Println(isPalindrome(10)) // Output: false
}
