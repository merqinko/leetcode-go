package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "342 + 465 = 807",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "0 + 0 = 0",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "9999999 + 9999 = 10009998",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "different lengths",
			l1:       []int{9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.l1)
			l2 := createList(tt.l2)
			result := addTwoNumbers(l1, l2)
			got := printList(result)

			if len(got) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(got))
				return
			}

			for i := range tt.expected {
				if got[i] != tt.expected[i] {
					t.Errorf("index %d: expected %v, got %v", i, tt.expected, got)
					return
				}
			}
		})
	}
}

func TestCreateList(t *testing.T) {
	// Empty slice
	if createList([]int{}) != nil {
		t.Error("expected nil for empty slice")
	}

	// Single element
	head := createList([]int{5})
	if head.Val != 5 || head.Next != nil {
		t.Error("single element list failed")
	}

	// Multiple elements
	head = createList([]int{1, 2, 3})
	if head.Val != 1 || head.Next.Val != 2 || head.Next.Next.Val != 3 {
		t.Error("multiple element list failed")
	}
}

func TestPrintList(t *testing.T) {
	// Nil list
	got := printList(nil)
	if len(got) != 0 {
		t.Errorf("expected empty slice for nil, got %v", got)
	}

	// Single element
	got = printList(&ListNode{Val: 7})
	if len(got) != 1 || got[0] != 7 {
		t.Errorf("expected [7], got %v", got)
	}
}
