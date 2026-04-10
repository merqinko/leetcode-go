# 015 - 3Sum

**Difficulty:** Medium  
**URL:** https://leetcode.com/problems/3sum/

---

## 📝 Problem Statement

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that:
- `i != j`, `i != k`, and `j != k`
- `nums[i] + nums[j] + nums[k] == 0`

**Notice:** The solution set must not contain duplicate triplets.

### Examples

**Example 1:**
```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0
The distinct triplets are [-1,-1,2] and [-1,0,1].
```

**Example 2:**
```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
```

**Example 3:**
```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
```

---

## 💡 Approach

**Two Pointers After Sorting:**

1. **Sort** the array (enables two-pointer technique)
2. **Fix** one number (`nums[i]`), search for other two using two pointers
3. **Skip duplicates** to avoid duplicate triplets
4. **Early exit** if fixed number is positive (sum can't be zero)

**Key Insight:**
- After sorting, if `nums[i] > 0`, we can stop (all remaining numbers are positive)
- For each fixed number, use `left` and `right` pointers to find pairs that sum to `-nums[i]`

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                              |
|---------|--------------------|------------------------------------------|
| **Time**  | O(n²)              | Outer loop O(n) × Two pointers O(n)      |
| **Space** | O(1) or O(n)       | O(1) excluding result, O(n) if sorting counts |

---

## 🔍 Algorithm Walkthrough

```
Input: [-1, 0, 1, 2, -1, -4]
After sort: [-4, -1, -1, 0, 1, 2]

i=0: nums[0] = -4
  left=1, right=5: -4 + (-1) + 2 = -3 (too small, move left)
  left=2, right=5: -4 + (-1) + 2 = -3 (too small, move left)
  left=3, right=5: -4 + 0 + 2 = -2 (too small, move left)
  left=4, right=5: -4 + 1 + 2 = -1 (too small, move left)

i=1: nums[1] = -1
  left=2, right=5: -1 + (-1) + 2 = 0 ✅ Found! [-1, -1, 2]
  left=3, right=4: -1 + 0 + 1 = 0 ✅ Found! [-1, 0, 1]

i=2: nums[2] = -1 (skip, duplicate of i=1)

i=3: nums[3] = 0
  left=4, right=5: 0 + 1 + 2 = 3 (too large, move right)

Result: [[-1, -1, 2], [-1, 0, 1]]
```

---

## 🧪 Testing

Run tests:
```bash
go test -v ./medium/015-3sum/
```

Run manually:
```bash
go run medium/015-3sum/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Classic example: `[-1, 0, 1, 2, -1, -4]`
- ✅ All zeros: `[0, 0, 0]`
- ✅ Multiple solutions: `[-2, 0, 1, 1, 2]`
- ✅ Duplicates in input: `[-1, 0, 1, 0]`
- ✅ Too few elements: `[0, 1]`
- ✅ No solution: `[0, 1, 1]`
- ✅ All positive: `[1, 2, 3]`
- ✅ All negative: `[-3, -2, -1]`

---

## 🎯 Key Learnings

1. **Sorting enables two-pointer technique** - reduces O(n³) to O(n²)
2. **Skip duplicates** at every level to avoid duplicate triplets
3. **Early exit optimization** - if `nums[i] > 0`, break (can't sum to zero)
4. **Two pointer pattern** - works for any k-Sum problem (2Sum, 3Sum, 4Sum)
5. **Handle edge cases** - arrays with < 3 elements return empty result
