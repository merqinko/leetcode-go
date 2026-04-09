# 014 - Longest Common Prefix

**Difficulty:** Easy  
**URL:** https://leetcode.com/problems/longest-common-prefix/

---

## 📝 Problem Statement

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

### Examples

**Example 1:**
```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

**Example 2:**
```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

---

## 💡 Approach

**Horizontal Scanning:**
1. Start with `prefix = first string`
2. Compare prefix with second string
3. If doesn't match, shorten prefix by removing last character
4. Repeat until prefix matches or becomes empty
5. Move to next string and repeat

---

## ⏱️ Complexity Analysis

| Metric  | Value              | Explanation                              |
|---------|--------------------|------------------------------------------|
| **Time**  | O(S)               | S = total characters across all strings  |
| **Space** | O(1)               | Only store prefix string                 |

---

## 🧪 Testing

Run tests:
```bash
go test -v ./easy/014-longest-common-prefix/
```

Run manually:
```bash
go run easy/014-longest-common-prefix/solution.go
```

---

## ✅ Test Cases Covered

- ✅ Common prefix: `["flower", "flow", "flight"]`
- ✅ No common prefix: `["dog", "racecar", "car"]`
- ✅ Long prefix: `["interspecies", "interstellar", "interstate"]`
- ✅ Single string: `["single"]`
- ✅ Empty strings: `["", ""]`
- ✅ Empty slice: `[]`
- ✅ All same: `["abc", "abc", "abc"]`
- ✅ Prefix longer than some strings: `["ab", "abc"]`

---

## 🎯 Key Learnings

1. **`strings.HasPrefix()`** is perfect for prefix matching
2. **Shortening strategy**: Remove one char at a time until match
3. **Edge cases matter**: Empty slice, empty strings, single string
4. **Early exit**: Return `""` as soon as prefix becomes empty
