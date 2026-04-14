# 018 - 4Sum

**Difficulty:** Medium  
**URL:** https://leetcode.com/problems/4sum/

---

## 📝 Problem Statement

Given an array `nums` of `n` integers, return all the quadruplets `[nums[a], nums[b], nums[c], nums[d]]` such that:
- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are distinct
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

**Notice:** The solution set must not contain duplicate quadruplets.

### Examples

**Example 1:**
```
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**Example 2:**
```
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
```

---

## 💡 Approach

**Two Pointers After Sorting (Extended 3Sum):**

1. **Sort** the array
2. **Fix TWO numbers** with nested loops
3. **Use two pointers** to find the remaining two numbers
4. **Skip duplicates** at every level to avoid duplicate quadruplets

**Pattern:**
- 2Sum → Two Pointers: O(n)
- 3Sum → Fix 1 + Two Pointers: O(n²)
- 4Sum → Fix 2 + Two Pointers: O(n³)
- k-Sum → Fix (k-2) + Two Pointers: O(n^(k-1))

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                              |
|---------|--------------------|------------------------------------------|
| **Time**  | O(n³)              | Two nested loops O(n²) × Two pointers O(n) |
| **Space** | O(1) or O(n)       | O(1) excluding result, O(n) if sorting counts |

---

## 🧪 Testing

Run tests:
```bash
go test -v ./medium/018-4sum/
```

Run manually:
```bash
go run medium/018-4sum/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Classic example: `[1, 0, -1, 0, -2, 2], target=0`
- ✅ All same numbers: `[2, 2, 2, 2, 2], target=8`
- ✅ Multiple solutions: `[-3, -2, -1, 0, 0, 1, 2, 3], target=0`
- ✅ Too few elements: `[0, 0, 0]`
- ✅ No solution: `[1, 2, 3, 4], target=100`

---

## 🎯 Key Learnings

1. **k-Sum pattern**: Fix (k-2) numbers, then use two pointers for the rest
2. **Skip duplicates at every level** — essential to avoid duplicate results
3. **Sorting enables two-pointer technique** — reduces brute force O(n^k) to O(n^(k-1))
4. **General formula**: k-Sum → O(n^(k-1)) with sorting + two pointers
