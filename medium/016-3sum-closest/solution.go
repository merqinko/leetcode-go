package main

import (
	"fmt"
	"math"
	"sort"
)

// ============================================================
// SOLUTION: Two Pointers After Sorting
// ============================================================
// Approach:
// 1. Sort the array
// 2. Fix one number, use two pointers to find closest sum to target
// 3. Track the minimum difference and closest sum found
// Time: O(n^2) - outer loop + two pointer scan
// Space: O(1) - excluding result storage

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)                   // Sort for two-pointer technique
	closestSum := nums[0] + nums[1] + nums[2] // Initialize with first triplet
	minDiff := math.Abs(float64(target - closestSum)) // Track minimum difference

	// Fix the first number
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1              // Start after fixed number
		right := len(nums) - 1     // End of array

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			currentDiff := math.Abs(float64(target - currentSum))

			// Update closest sum if we found a smaller difference
			if currentDiff < minDiff {
				minDiff = currentDiff
				closestSum = currentSum
			}

			// Move pointers based on comparison with target
			if currentSum < target {
				left++  // Need larger sum
			} else if currentSum > target {
				right-- // Need smaller sum
			} else {
				// Exact match found - can't get closer than this!
				return target
			}
		}
	}

	return closestSum
}

func main() {
	// Test cases
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},        // Closest: 2 (-1+2+1=2, diff=1)
		{[]int{0, 0, 0}, 1, 0},             // Closest: 0 (0+0+0=0, diff=1)
		{[]int{1, 1, 1, 0}, -100, 2},       // Closest: 2 (0+1+1=2, diff=102)
		{[]int{1, 1, 1, 1}, 0, 3},          // Closest: 3 (1+1+1=3, diff=3)
		{[]int{-1, 2, 1, -4}, 2, 2},        // Closest: 2 (-1+2+1=2, exact match)
		{[]int{0, 1, 2}, 3, 3},             // Closest: 3 (0+1+2=3, exact match)
	}

	fmt.Println("=== 3Sum Closest ===")
	for i, tt := range tests {
		result := threeSumClosest(tt.nums, tt.target)
		status := "✅"
		if result != tt.expected {
			status = "❌"
		}
		fmt.Printf("%s Test %d: nums=%v, target=%d → %d (expected %d)\n",
			status, i+1, tt.nums, tt.target, result, tt.expected)
	}
}
