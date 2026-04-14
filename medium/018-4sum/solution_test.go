package main

import (
	"reflect"
	"sort"
	"testing"
)

// Helper to compare quadruplets regardless of order
func equalQuadruplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return false
	})

	sort.Slice(b, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if b[i][k] != b[j][k] {
				return b[i][k] < b[j][k]
			}
		}
		return false
	})

	return reflect.DeepEqual(a, b)
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		// Happy Cases
		{"classic example", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"all same", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{"multiple solutions", []int{-3, -2, -1, 0, 0, 1, 2, 3}, 0, [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},

		// Edge Cases
		{"too few elements", []int{0, 0, 0}, 0, [][]int{}},
		{"no solution", []int{1, 2, 3, 4}, 100, [][]int{}},
		{"negative target", []int{-5, -4, -3, -2, -1}, -10, [][]int{{-4, -3, -2, -1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			if !equalQuadruplets(got, tt.expected) {
				t.Errorf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
