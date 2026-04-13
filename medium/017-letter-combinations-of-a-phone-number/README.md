# 017 - Letter Combinations of a Phone Number

**Difficulty:** Medium  
**URL:** https://leetcode.com/problems/letter-combinations-of-a-phone-number/

---

## 📝 Problem Statement

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters is just like on the telephone buttons. Note that `1` does not map to any letters.

### Examples

**Example 1:**
```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**
```
Input: digits = ""
Output: []
```

---

## 💡 Approach

**Backtracking (Depth First Search):**

1. **Map digits to letters** (standard phone keypad mapping)
2. **DFS Recursion**: Build strings character by character
3. **Base case**: String length equals input length → add to results
4. **Recursive step**: Try all possible letters for current digit

**Visual Tree (Input "23"):**
```
Start
├── a (from 2)
│   ├── ad (from 3) ✅
│   ├── ae (from 3) ✅
│   └── af (from 3) ✅
├── b (from 2)
│   ├── bd (from 3) ✅
│   ├── be (from 3) ✅
│   └── bf (from 3) ✅
└── c (from 2)
    ├── cd (from 3) ✅
    ├── ce (from 3) ✅
    └── cf (from 3) ✅
```

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                              |
|---------|--------------------|------------------------------------------|
| **Time**  | O(4^n × n)         | n = digits, 4 = max chars per digit (7,9) |
| **Space** | O(n)               | Recursion stack depth = n                |

---

## 🧪 Testing

Run tests:
```bash
go test -v ./medium/017-letter-combinations-of-a-phone-number/
```

Run manually:
```bash
go run medium/017-letter-combinations-of-a-phone-number/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Two digits: `"23"`
- ✅ Single digit: `"2"`
- ✅ Digit with 4 chars: `"7"` (pqrs)
- ✅ Four digits: `"234"`
- ✅ Empty string: `""`
- ✅ No mapping digits: `"1"`, `"0"`

---

## 🎯 Key Learnings

1. **Backtracking Pattern**: Try option → Recurse → Undo option
2. **DFS vs BFS**: DFS naturally builds paths depth-first, simpler code
3. **String Concatenation**: In Go, `path + char` creates new string (no need to "undo" slice)
4. **Edge Cases**: Empty input returns empty slice, not `[""]`
5. **Phone Mapping**: 7 and 9 have 4 letters, others have 3
