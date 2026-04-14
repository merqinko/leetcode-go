package main

import (
	"fmt"
	"sort"
)

// ============================================================
// SOLUTION: Two Pointers After Sorting (Extended 3Sum)
// ============================================================
// Approach:
// 1. Sort the array
// 2. Fix TWO numbers with nested loops, use two pointers for rest
// 3. Skip duplicates at every level
// 4. Early exit optimizations
// Time: O(n^3) - two outer loops + two pointer scan
// Space: O(1) - excluding result storage

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums) // Sort for two-pointer technique

	// Fix first number
	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicate first numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Fix second number
		for j := i + 1; j < len(nums)-2; j++ {
			// Skip duplicate second numbers
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Two pointers for remaining two numbers
			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum < target {
					left++  // Need larger sum
				} else if sum > target {
					right-- // Need smaller sum
				} else {
					// Found valid quadruplet!
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// Skip duplicates for left pointer
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// Skip duplicates for right pointer
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					// Move both pointers
					left++
					right--
				}
			}
		}
	}

	return result
}

func main() {
	tests := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0, [][]int{{-3, 0, 1, 2}, {-3, -1, 1, 3}, {-3, -2, 2, 3}, {-2, -1, 0, 3}, {-2, 0, 1, 1}, {-1, 0, 0, 1}}},
		{[]int{0, 0, 0}, 0, [][]int{}},
	}

	fmt.Println("=== 4Sum ===")
	for i, tt := range tests {
		result := fourSum(tt.nums, tt.target)
		fmt.Printf("Test %d: nums=%v, target=%d\n", i+1, tt.nums, tt.target)
		fmt.Printf("  Result: %v\n\n", result)
	}
}
