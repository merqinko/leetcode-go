package main

import (
	"fmt"
	"sort"
)

// ============================================================
// SOLUTION: Two Pointers After Sorting
// ============================================================
// Approach:
// 1. Sort the array
// 2. Fix one number, use two pointers to find other two
// 3. Skip duplicates to avoid duplicate triplets
// Time: O(n^2) - outer loop + two pointer scan
// Space: O(1) - excluding result storage

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums) // Sort first: enables two-pointer technique

	// Fix the first number
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate first numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early exit: if first number is positive, sum can't be zero
		if nums[i] > 0 {
			break
		}

		// Two pointers: search for two numbers that sum to -nums[i]
		left := i + 1       // Start after fixed number
		right := len(nums) - 1 // End of array

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				// Sum too small, need larger numbers
				left++
			} else if sum > 0 {
				// Sum too large, need smaller numbers
				right--
			} else {
				// Found a triplet that sums to zero!
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers inward
				left++
				right--
			}
		}
	}

	return result
}

func main() {
	// Test cases
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			input:    []int{-1, 0, 1, 0},
			expected: [][]int{{-1, 0, 1}},
		},
	}

	fmt.Println("=== 3Sum ===")
	for i, tt := range tests {
		result := threeSum(tt.input)
		fmt.Printf("Test %d: %v\n", i+1, tt.input)
		fmt.Printf("  Result: %v\n", result)
		fmt.Printf("  Expected: %v\n\n", tt.expected)
	}
}
