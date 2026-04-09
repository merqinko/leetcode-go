# 013 - Roman to Integer

**Difficulty:** Easy  
**URL:** https://leetcode.com/problems/roman-to-integer/

---

## 📝 Problem Statement

Given a Roman numeral, convert it to an integer.

Roman numerals are represented by seven different symbols: `I`, `V`, `X`, `L`, `C`, `D` and `M`.

| Symbol | Value |
|--------|-------|
| I      | 1     |
| V      | 5     |
| X      | 10    |
| L      | 50    |
| C      | 100   |
| D      | 500   |
| M      | 1000  |

### Examples

**Example 1:**
```
Input: s = "III"
Output: 3
Explanation: III = 3
```

**Example 2:**
```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V = 5, III = 3
```

**Example 3:**
```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90, IV = 4
```

---

## 💡 Approach

**Key Insight:** Loop from right to left
- If current value < previous value → **subtract** (e.g., IV = 5 - 1 = 4)
- If current value >= previous value → **add** (e.g., VI = 5 + 1 = 6)

**Two implementations provided:**
1. **Switch Statement** (Recommended - faster, no memory allocation)
2. **Map Lookup** (More readable, slightly slower)

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                          |
|---------|--------------------|--------------------------------------|
| **Time**  | O(n)               | Single pass through the string       |
| **Space** | O(1)               | Fixed number of variables            |

---

## 🧪 Testing

Run tests:
```bash
go test ./easy/013-roman-to-integer/
```

Run with verbose output:
```bash
go test -v ./easy/013-roman-to-integer/
```

Run manually:
```bash
go run easy/013-roman-to-integer/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Simple numbers: III, V, X
- ✅ Subtraction cases: IV, IX, XL, XC, CD, CM
- ✅ Medium numbers: LVIII (58), XXVII (27)
- ✅ Large numbers: MCMXCIV (1994)
- ✅ Edge case: Empty string ("")

---

## 🎯 Key Learnings

1. **Right-to-left iteration** simplifies subtraction logic
2. **Switch vs Map tradeoff**: Switch is faster, map is more readable
3. **Subtraction rule**: Only applies when smaller value precedes larger value
4. **Valid input**: Problem guarantees valid Roman numerals (1-3999)
