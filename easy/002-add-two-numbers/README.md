# #2 - Add Two Numbers

**Difficulty:** Medium  
**Approach:** Iterative with Carry  
**Time:** O(max(len(l1), len(l2)))  
**Space:** O(max(len(l1), len(l2))) for result

## Problem
You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each node contains a single digit. Add the two numbers and return the sum as a linked list.

## Example
```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

## Approach
1. Use a **dummy node** to simplify list construction
2. Track **carry** from addition (0 or 1)
3. Iterate while either list has nodes OR carry exists
4. At each step: `sum = val1 + val2 + carry`
5. New node value = `sum % 10`
6. New carry = `sum / 10`

## Edge Cases
- ✅ Different length lists
- ✅ Final carry creates extra node
- ✅ Both lists are single `0`
- ✅ Large numbers with many 9s

## Tests
```bash
go test ./easy/002-add-two-numbers/ -v
```
