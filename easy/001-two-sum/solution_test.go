package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "different indices",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "duplicate numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, 5, 8},
			target: 5,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != 2 || got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
