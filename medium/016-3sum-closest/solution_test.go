package main

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		// Happy Cases
		{"classic example", []int{-1, 2, 1, -4}, 1, 2},
		{"all zeros", []int{0, 0, 0}, 1, 0},
		{"exact match", []int{-1, 2, 1, -4}, 2, 2},
		{"exact match simple", []int{0, 1, 2}, 3, 3},

		// Edge Cases
		{"large negative target", []int{1, 1, 1, 0}, -100, 2},
		{"all same numbers", []int{1, 1, 1, 1}, 0, 3},
		{"minimum array size", []int{-1, 0, 1}, 0, 0},
		{"target smaller than any sum", []int{5, 5, 5, 5}, -100, 15},
		{"target larger than any sum", []int{-5, -5, -5, -5}, 100, -15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("threeSumClosest(%v, %d) = %d, want %d",
					tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
