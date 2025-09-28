# 18. Common Interview Problems

This example demonstrates solutions to 12 classic LeetCode-style coding interview problems in Go. These problems cover a wide range of algorithms and data structures commonly tested in technical interviews.

## Problems Covered

### 1. Two Sum (Easy)
- **Problem**: Find two numbers in array that add up to target
- **Solution**: Hash map for O(n) time complexity
- **Key Insight**: Store complement values for quick lookup

### 2. Valid Parentheses (Easy)
- **Problem**: Check if string has valid parentheses
- **Solution**: Stack-based approach
- **Key Insight**: Use stack to track opening brackets

### 3. Merge Two Sorted Lists (Easy)
- **Problem**: Merge two sorted linked lists
- **Solution**: Iterative merging with dummy node
- **Key Insight**: Compare nodes and build result list

### 4. Maximum Subarray (Medium)
- **Problem**: Find contiguous subarray with largest sum
- **Solution**: Kadane's algorithm
- **Key Insight**: Track current and global maximum sums

### 5. Climbing Stairs (Easy)
- **Problem**: Number of ways to climb n stairs (1 or 2 steps)
- **Solution**: Dynamic programming
- **Key Insight**: Fibonacci sequence pattern

### 6. Binary Tree Maximum Path Sum (Hard)
- **Problem**: Find maximum path sum in binary tree
- **Solution**: DFS with path tracking
- **Key Insight**: Consider paths through each node

### 7. Longest Palindromic Substring (Medium)
- **Problem**: Find longest palindromic substring
- **Solution**: Expand around centers
- **Key Insight**: Check both odd and even length palindromes

### 8. Median of Two Sorted Arrays (Hard)
- **Problem**: Find median of two sorted arrays
- **Solution**: Binary search approach
- **Key Insight**: Partition arrays to find median

### 9. Regular Expression Matching (Hard)
- **Problem**: Implement regex matching with . and *
- **Solution**: Dynamic programming
- **Key Insight**: 2D DP table for pattern matching

### 10. Word Break (Medium)
- **Problem**: Check if string can be segmented into dictionary words
- **Solution**: Dynamic programming
- **Key Insight**: DP array tracks segmentable prefixes

### 11. Number of Islands (Medium)
- **Problem**: Count islands in 2D grid
- **Solution**: DFS flood fill
- **Key Insight**: Mark visited cells to avoid recounting

### 12. Longest Increasing Subsequence (Medium)
- **Problem**: Find length of longest increasing subsequence
- **Solution**: Dynamic programming
- **Key Insight**: DP tracks LIS ending at each position

## Time & Space Complexity

| Problem | Time Complexity | Space Complexity | Difficulty |
|---------|----------------|------------------|------------|
| Two Sum | O(n) | O(n) | Easy |
| Valid Parentheses | O(n) | O(n) | Easy |
| Merge Two Sorted Lists | O(n+m) | O(1) | Easy |
| Maximum Subarray | O(n) | O(1) | Medium |
| Climbing Stairs | O(n) | O(n) | Easy |
| Binary Tree Max Path Sum | O(n) | O(h) | Hard |
| Longest Palindromic Substring | O(n²) | O(1) | Medium |
| Median of Two Sorted Arrays | O(log(min(m,n))) | O(1) | Hard |
| Regular Expression Matching | O(m*n) | O(m*n) | Hard |
| Word Break | O(n²) | O(n) | Medium |
| Number of Islands | O(m*n) | O(m*n) | Medium |
| Longest Increasing Subsequence | O(n²) | O(n) | Medium |

## Common Patterns

### Hash Map for Quick Lookup
```go
// Two Sum pattern
numMap := make(map[int]int)
for i, num := range nums {
    complement := target - num
    if j, exists := numMap[complement]; exists {
        return []int{j, i}
    }
    numMap[num] = i
}
```

### Stack for Parentheses/Backtracking
```go
// Valid parentheses pattern
stack := []rune{}
for _, char := range s {
    if char == '(' || char == '{' || char == '[' {
        stack = append(stack, char)
    } else {
        if len(stack) == 0 {
            return false
        }
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // Check matching
    }
}
```

### Dynamic Programming
```go
// Climbing stairs pattern
dp := make([]int, n+1)
dp[1] = 1
dp[2] = 2
for i := 3; i <= n; i++ {
    dp[i] = dp[i-1] + dp[i-2]
}
```

### Two Pointers
```go
// Merge sorted lists pattern
for list1 != nil && list2 != nil {
    if list1.Val <= list2.Val {
        current.Next = list1
        list1 = list1.Next
    } else {
        current.Next = list2
        list2 = list2.Next
    }
    current = current.Next
}
```

### DFS/BFS for Graph Problems
```go
// Island counting pattern
var dfs func(int, int)
dfs = func(r, c int) {
    if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
        return
    }
    grid[r][c] = '0' // Mark visited
    dfs(r-1, c) // Up
    dfs(r+1, c) // Down
    dfs(r, c-1) // Left
    dfs(r, c+1) // Right
}
```

## Interview Tips

### Problem-Solving Approach
1. **Understand the problem** - Read carefully, ask clarifying questions
2. **Identify patterns** - Recognize similar problems you've solved
3. **Choose data structures** - Pick appropriate structures for the problem
4. **Consider edge cases** - Empty inputs, single elements, boundaries
5. **Optimize** - Look for better time/space complexity
6. **Test thoroughly** - Check with various inputs

### Common Mistakes
- **Off-by-one errors** - Careful with array indices
- **Null pointer dereference** - Check for nil values
- **Infinite loops** - Ensure termination conditions
- **Wrong time complexity** - Understand algorithm complexity
- **Missing edge cases** - Test with extreme inputs

### Communication
- **Think aloud** - Explain your thought process
- **Ask questions** - Clarify requirements and constraints
- **Explain trade-offs** - Discuss different approaches
- **Write clean code** - Use meaningful variable names
- **Test your code** - Verify with examples

## Practice Strategy

### By Difficulty
- **Easy**: Focus on understanding basic patterns
- **Medium**: Combine multiple concepts
- **Hard**: Require deep algorithmic knowledge

### By Topic
- **Arrays/Strings**: Two pointers, sliding window
- **Linked Lists**: Pointer manipulation
- **Trees**: Recursion, traversal
- **Graphs**: DFS/BFS, connectivity
- **Dynamic Programming**: Subproblems, optimization

### By Company
- **FAANG**: Focus on system design + algorithms
- **Startups**: Practical coding + problem-solving
- **General**: Well-rounded algorithmic knowledge

## Running the Example

```bash
cd 18-common-interview
go run main.go
```

This comprehensive example provides solutions to fundamental coding interview problems, demonstrating key algorithms and data structures essential for technical interviews.

## Additional Practice Problems

### 13. Container With Most Water (Medium)

**Problem**: Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

**Solution Approach**:
```go
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {
        // Calculate area with current left and right pointers
        width := right - left
        h := min(height[left], height[right])
        area := width * h
        maxArea = max(maxArea, area)

        // Move the pointer with smaller height
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}
```

**Key Insight**: Use two pointers starting from ends. Move the pointer with smaller height inward, as increasing the smaller height gives better chance for larger area.

### 14. 3Sum (Medium)

**Problem**: Given an array nums of n integers, find all unique triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

**Solution Approach**:
```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}

    for i := 0; i < len(nums)-2; i++ {
        // Skip duplicates for i
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})

                // Skip duplicates for left
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // Skip duplicates for right
                for left < right && nums[right] == nums[right-1] {
                    right--
                }

                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }

    return result
}
```

**Key Insight**: Sort array first, then use two pointers for each element. Skip duplicates to avoid duplicate triplets.

### 15. Longest Substring Without Repeating Characters (Medium)

**Problem**: Given a string s, find the length of the longest substring without repeating characters.

**Solution Approach**:
```go
func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLength := 0
    start := 0

    for end := 0; end < len(s); end++ {
        if lastSeen, exists := charIndex[s[end]]; exists && lastSeen >= start {
            start = lastSeen + 1
        }

        charIndex[s[end]] = end
        currentLength := end - start + 1

        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}
```

**Key Insight**: Use sliding window with a map to track last seen index of each character. Move start pointer when duplicate found.