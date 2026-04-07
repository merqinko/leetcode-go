# LeetCode Solutions - Go

Daily practice solving LeetCode problems in Go.

## 📊 Progress Tracker
- **Easy:** 3/100 ✅
- **Medium:** 0/100
- **Hard:** 0/100
- **Total:** 3/300

## 📁 Structure
```
easy/     - Easy difficulty problems (0-100)
medium/   - Medium difficulty problems (100-300)
hard/     - Hard difficulty problems (300+)
```

## 🧪 How to Run

**Run all tests:**
```bash
go test ./...
```

**Run specific category:**
```bash
go test ./easy/...
```

**Verbose output:**
```bash
go test -v ./...
```

**Run & see output:**
```bash
go run easy/001-two-sum/solution.go
```

## ✅ Completed Problems

### Easy
1. **[001 - Two Sum](./easy/001-two-sum/)** ✅
   - Approach: Hash Map
   - Time: O(n) | Space: O(n)

2. **[002 - Add Two Numbers](./easy/002-add-two-numbers/)** ✅
   - Approach: Iterative with Carry
   - Time: O(max(m,n)) | Space: O(max(m,n))

3. **[009 - Palindrome Number](./easy/009-palindrome-number/)** ✅
   - Approach: Mathematical Reverse
   - Time: O(log₁₀ n) | Space: O(1)

---

## 🚀 Daily Workflow

1. **Pick a problem** from LeetCode
2. **Create folder:** `mkdir easy/00X-problem-name`
3. **Write solution:** `solution.go`
4. **Write tests:** `solution_test.go`
5. **Create README:** `README.md`
6. **Run tests:** `go test ./easy/00X-problem-name/`
7. **Commit:** `git add . && git commit -m "Add problem #X"`

## 💡 Tips for Success
- ✅ Write tests FIRST
- ✅ Optimize for readability first
- ✅ Add comments explaining approach
- ✅ Track time/space complexity
- ✅ Test edge cases
- ✅ Commit daily
- ✅ Explain in README
