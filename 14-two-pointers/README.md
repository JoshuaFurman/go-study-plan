# 14. Two Pointers & Sliding Window

This example demonstrates two-pointer techniques and sliding window algorithms, essential for solving array and string problems efficiently. These patterns are commonly used in coding interviews.

## Key Concepts

### Two Pointers Technique
- Use two pointers to traverse array from different positions
- Often used for sorted arrays or when searching for pairs
- Can solve problems in O(n) time instead of O(n²)

### Sliding Window
- Maintains a window of elements as we traverse
- Useful for substring/subarray problems
- Can find optimal subarrays efficiently

## Two Pointers Patterns

### Opposite Direction (Converging)
```go
left, right := 0, len(arr)-1
for left < right {
    // Move pointers based on condition
    if condition {
        left++
    } else {
        right--
    }
}
```

### Same Direction (Fast/Slow)
```go
slow, fast := 0, 0
for fast < len(arr) {
    // Fast moves ahead
    // Slow moves conditionally
    if condition {
        slow++
    }
    fast++
}
```

## Sliding Window Patterns

### Fixed Size Window
```go
windowSize := k
for i := 0; i <= len(arr)-windowSize; i++ {
    // Process window from i to i+windowSize-1
}
```

### Variable Size Window
```go
left := 0
for right := 0; right < len(arr); right++ {
    // Expand window
    // Shrink from left when condition met
    for condition && left <= right {
        left++
    }
}
```

## Common Problems & Solutions

### 1. Two Sum (Sorted Array)
```go
func TwoSum(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return []int{left, right}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return nil
}
```

### 2. Remove Duplicates from Sorted Array
```go
func RemoveDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}
```

### 3. Container With Most Water
```go
func MaxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0
    for left < right {
        width := right - left
        h := min(height[left], height[right])
        area := width * h
        maxArea = max(maxArea, area)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}
```

### 4. Minimum Size Subarray Sum
```go
func MinSubArrayLen(target int, nums []int) int {
    minLength := math.MaxInt32
    sum := 0
    left := 0
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target && left <= right {
            minLength = min(minLength, right-left+1)
            sum -= nums[left]
            left++
        }
    }
    if minLength == math.MaxInt32 {
        return 0
    }
    return minLength
}
```

### 5. Longest Substring Without Repeating Characters
```go
func LengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLength := 0
    left := 0
    for right := 0; right < len(s); right++ {
        if idx, exists := charIndex[s[right]]; exists && idx >= left {
            left = idx + 1
        }
        charIndex[s[right]] = right
        maxLength = max(maxLength, right-left+1)
    }
    return maxLength
}
```

## Time Complexity Analysis

| Problem | Time Complexity | Space Complexity |
|---------|----------------|------------------|
| Two Sum (sorted) | O(n) | O(1) |
| Remove Duplicates | O(n) | O(1) |
| Container With Most Water | O(n) | O(1) |
| 3Sum | O(n²) | O(1) |
| Minimum Subarray Sum | O(n) | O(1) |
| Longest Substring | O(n) | O(min(n, charset)) |

## When to Use Two Pointers

1. **Sorted Arrays**: Finding pairs, triplets with target sum
2. **String Problems**: Palindromes, anagrams, substring searches
3. **Array Modification**: Remove elements, deduplication
4. **Optimization Problems**: Maximum area, minimum length

## When to Use Sliding Window

1. **Substring Problems**: Longest substring with condition
2. **Subarray Problems**: Maximum/minimum subarray with constraints
3. **Contiguous Elements**: Problems requiring contiguous segments
4. **Fixed/Variable Windows**: Based on problem requirements

## Common Mistakes

1. **Off-by-one errors**: Careful with loop bounds
2. **Infinite loops**: Ensure pointers move correctly
3. **Edge cases**: Empty arrays, single elements
4. **Pointer updates**: Update both pointers appropriately
5. **Return values**: Return indices vs values vs lengths

## Practice Problems

### Easy
- Remove Duplicates from Sorted Array
- Move Zeroes
- Reverse String
- Valid Palindrome

### Medium
- Container With Most Water
- 3Sum
- Minimum Size Subarray Sum
- Longest Substring Without Repeating Characters

### Hard
- Trapping Rain Water
- Longest Substring with At Most K Distinct Characters
- Minimum Window Substring
- Sliding Window Maximum

## Running the Example

```bash
cd 14-two-pointers
go run main.go
```

This example demonstrates 12 different two-pointer and sliding window problems, providing a comprehensive foundation for solving similar problems in coding interviews.