package main

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to compare two 2D slices regardless of order
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort both slices for consistent comparison
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] != a[j][0] {
			return a[i][0] < a[j][0]
		}
		if a[i][1] != a[j][1] {
			return a[i][1] < a[j][1]
		}
		return a[i][2] < a[j][2]
	})

	sort.Slice(b, func(i, j int) bool {
		if b[i][0] != b[j][0] {
			return b[i][0] < b[j][0]
		}
		if b[i][1] != b[j][1] {
			return b[i][1] < b[j][1]
		}
		return b[i][2] < b[j][2]
	})

	return reflect.DeepEqual(a, b)
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		// Happy Cases
		{"classic example", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"all zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"multiple solutions", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"duplicates in input", []int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},

		// Edge Cases
		{"too few elements", []int{0, 1}, [][]int{}},
		{"no solution", []int{0, 1, 1}, [][]int{}},
		{"all positive", []int{1, 2, 3}, [][]int{}},
		{"all negative", []int{-3, -2, -1}, [][]int{}},
		{"many duplicates", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !equalTriplets(got, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
