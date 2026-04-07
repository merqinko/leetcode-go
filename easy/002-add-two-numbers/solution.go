package main

import "fmt"

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Problem #2: Add Two Numbers
// Input: Two linked lists representing numbers in reverse order
// Output: Linked list representing the sum
// Example: 342 + 465 = 807
//          2→4→3 + 5→6→4 = 7→0→8

// addTwoNumbers adds two numbers represented as linked lists
// Time: O(max(len(l1), len(l2)))
// Space: O(max(len(l1), len(l2))) for result
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// dummy: fake starting node (we ignore it at the end)
	// temp: pointer to track where we're building the result
	dummy := &ListNode{}
	temp := dummy
	carry := 0 // Track overflow from addition (0 or 1)

	// Loop while both lists have nodes OR we have a carry
	for l1 != nil || l2 != nil || carry != 0 {
		// Get value from l1, or 0 if l1 is empty
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		// Get value from l2, or 0 if l2 is empty
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Add both digits + carry from previous iteration
		sum := val1 + val2 + carry

		// Extract carry for next iteration (0 or 1)
		carry = sum / 10

		// Create new node with digit (0-9)
		temp.Next = &ListNode{Val: sum % 10}

		// Move pointer to new node
		temp = temp.Next
	}

	// Return dummy.Next to skip the fake dummy node
	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     sum := 0
//     carry := 0
//     l3 := &ListNode{}
//     head := l3

//     for l1 != nil || l2 != nil {
//         sum = 0

//         if l1 != nil {
//             sum += l1.Val
//             l1 = l1.Next
//         }

//         if l2 != nil {
//             sum += l2.Val
//             l2 = l2.Next
//         }
//         sum += carry
//         carry = sum / 10
//         sum %= 10
//         l3.Next = &ListNode{Val: sum}
//         l3 = l3.Next
//     }

//     if carry != 0 {
//         l3.Next = &ListNode{Val: carry}
//     }

//     return head.Next
// }
// Helper: Create linked list from slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper: Print linked list
func printList(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func main() {
	// Test case 1: 342 + 465 = 807
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	fmt.Println("Test 1 (342 + 465):", printList(result)) // Expected: [7, 0, 8]

	// Test case 2: 0 + 0 = 0
	l1 = createList([]int{0})
	l2 = createList([]int{0})
	result = addTwoNumbers(l1, l2)
	fmt.Println("Test 2 (0 + 0):", printList(result)) // Expected: [0]

	// Test case 3: 9999999 + 9999 = 10009998
	l1 = createList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = createList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	fmt.Println("Test 3 (9999999 + 9999):", printList(result)) // Expected: [8, 9, 9, 9, 0, 0, 0, 1]
}
