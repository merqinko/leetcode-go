# #9 - Palindrome Number

**Difficulty:** Easy  
**Approach:** Mathematical Reverse (No String Conversion)  
**Time:** O(log₁₀ n) - number of digits  
**Space:** O(1)

## Problem
Given an integer `x`, return `true` if `x` is a palindrome integer, and `false` otherwise.

An integer is a **palindrome** when it reads the same backward as forward.

## Examples
```
Input: x = 121
Output: true

Input: x = -121
Output: false

Input: x = 10
Output: false
```

## Approach
1. **Negative numbers** → always `false` (minus sign breaks symmetry)
2. Store `original = x`
3. Reverse `x` mathematically:
   - Extract last digit: `digit = x % 10`
   - Build reversed: `reversed = reversed * 10 + digit`
   - Remove last digit: `x /= 10`
4. Return `original == reversed`

## Why Not String?
Converting to string uses O(n) extra space. Mathematical approach is **O(1) space**.

## Edge Cases
- ✅ Negative numbers → `false`
- ✅ Zero → `true`
- ✅ Single digit → `true`
- ✅ Numbers ending in 0 (except 0) → `false`

## Tests
```bash
go test ./easy/009-palindrome-number/ -v
```
