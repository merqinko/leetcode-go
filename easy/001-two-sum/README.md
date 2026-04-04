# Problem #1: Two Sum

**Level:** Easy  
**Difficulty:** Start  
**URL:** https://leetcode.com/problems/two-sum/

## Problem Description
Given an array of integers `nums` and an integer `target`, return the indices of the two numbers that add up to `target`.

You may assume that each input has exactly one solution, and you may not use the same element twice.

### Example
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: nums[0] + nums[1] == 9, so we return [0, 1].
```

## Approach
**Hash Map (Optimal)**
- Use a map to store numbers and their indices
- For each number, check if its complement (target - num) exists
- If found, return the indices immediately
- If not found, store the current number in the map

### Why This Works
- Single pass through array
- Constant lookup time in map
- No need to check every pair

## Complexity Analysis
- **Time:** O(n) - single pass through array
- **Space:** O(n) - hash map can store up to n elements

## Alternative Approaches
1. **Brute Force** - O(n²) time, O(1) space (nested loops)
2. **Sorting + Two Pointers** - O(n log n) time, O(1) space

## Solution
See `solution.go`
