# Problem #2: Add Two Numbers

**Level:** Medium  
**URL:** https://leetcode.com/problems/add-two-numbers/

## Problem Description
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each node contains a single digit. Add the two numbers and return the sum as a linked list.

### Example
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807

## Approach
**Linked List Traversal with Carry**
- Traverse both lists simultaneously
- At each node, add digits + carry from previous
- Handle carry (result > 9)
- Continue until both lists end and no carry remains

### Why This Works
- Single pass - O(n) time
- Process digits left to right (reverse order = natural order for addition)
- Carry handled elegantly

## Complexity Analysis
- **Time:** O(max(len(l1), len(l2)))
- **Space:** O(max(len(l1), len(l2))) for result list

## Key Concepts
1. **Linked List** - Chain of nodes with Next pointers
2. **Carry** - Overflow when sum > 9 (use `/` and `%`)
3. **Dummy Node** - Fake starting node for easier code

## Solution
See `solution.go`

