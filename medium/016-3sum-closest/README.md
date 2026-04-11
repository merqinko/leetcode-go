# 016 - 3Sum Closest

**Difficulty:** Medium  
**URL:** https://leetcode.com/problems/3sum-closest/

---

## 📝 Problem Statement

Given an integer array `nums` of length `n` and an integer `target`, return the sum of three integers in `nums` such that the sum is closest to `target`.

You may assume that each input would have exactly one solution.

### Examples

**Example 1:**
```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to 1 is 2. (-1 + 2 + 1 = 2)
```

**Example 2:**
```
Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to 1 is 0. (0 + 0 + 0 = 0)
```

---

## 💡 Approach

**Two Pointers After Sorting:**

1. **Sort** the array
2. **Fix** one number, use two pointers to explore all triplets
3. **Track** the closest sum found so far using `minDiff`
4. **Early exit** if exact match found (diff = 0)

**Key Difference from 3Sum:**
- Instead of finding sum == 0, find sum **closest to target**
- Track `minDiff` and `closestSum` instead of collecting all triplets
- No need to skip duplicates (we only need one answer)

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                              |
|---------|--------------------|------------------------------------------|
| **Time**  | O(n²)              | Outer loop O(n) × Two pointers O(n)      |
| **Space** | O(1) or O(n)       | O(1) excluding result, O(n) if sorting counts |

---

## 🔍 Algorithm Walkthrough

```
Input: nums = [-1,2,1,-4], target = 1
After sort: [-4,-1,1,2]

Initialize: closestSum = -4 + (-1) + 1 = -4, minDiff = |1 - (-4)| = 5

i=0: nums[0] = -4
  left=1, right=3: sum = -4 + (-1) + 2 = -3
    diff = |1 - (-3)| = 4 < 5 → Update! closestSum = -3, minDiff = 4
    sum < target → left++
  
  left=2, right=3: sum = -4 + 1 + 2 = -1
    diff = |1 - (-1)| = 2 < 4 → Update! closestSum = -1, minDiff = 2
    sum < target → left++

i=1: nums[1] = -1
  left=2, right=3: sum = -1 + 1 + 2 = 2
    diff = |1 - 2| = 1 < 2 → Update! closestSum = 2, minDiff = 1
    sum > target → right--

Result: 2 ✅
```

---

## 🧪 Testing

Run tests:
```bash
go test -v ./medium/016-3sum-closest/
```

Run manually:
```bash
go run medium/016-3sum-closest/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Classic example: `[-1, 2, 1, -4], target=1`
- ✅ All zeros: `[0, 0, 0], target=1`
- ✅ Exact match: `[-1, 2, 1, -4], target=2`
- ✅ Large negative target: `[1, 1, 1, 0], target=-100`
- ✅ All same numbers: `[1, 1, 1, 1], target=0`
- ✅ Minimum array size: `[-1, 0, 1], target=0`
- ✅ Target smaller/larger than any possible sum

---

## 🎯 Key Learnings

1. **Two pointers works for any 3Sum variant** - just change the condition
2. **Track state** (`minDiff`, `closestSum`) instead of collecting results
3. **Early exit on exact match** - if `sum == target`, return immediately
4. **No duplicate skipping needed** - problem guarantees exactly one solution
5. **`math.Abs()`** for calculating absolute difference
