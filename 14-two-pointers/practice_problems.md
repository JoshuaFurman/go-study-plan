# Practice Problems - Two Pointers

These practice problems reinforce two-pointer techniques and sliding window algorithms. Focus on efficient array/string manipulation.

## Problem 1: Two Sum (Sorted Array)

**Description:** Find two numbers in a sorted array that add up to a target sum. Return their indices (1-based).

**Requirements:**
- Array is sorted in ascending order
- Use two-pointer technique
- Return indices of the two numbers
- Handle edge cases (no solution, duplicates)

**Examples:**
- Input: [2, 7, 11, 15], target: 9 → [1, 2]
- Input: [2, 3, 4], target: 6 → [1, 3]
- Input: [1, 2, 3], target: 10 → []

**Constraints:**
- Time complexity: O(n)
- Space complexity: O(1)
- Exactly one solution exists (except when no solution)

### Solution

```go
package main

import "fmt"

func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]

		if currentSum == target {
			// Return 1-based indices
			return []int{left + 1, right + 1}
		} else if currentSum < target {
			left++ // Need larger sum
		} else {
			right-- // Need smaller sum
		}
	}

	return []int{} // No solution found
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 3, 4}, 6},
		{[]int{-1, 0, 1, 2}, 0},
		{[]int{1, 2, 3}, 10},
		{[]int{1, 2, 3, 4, 5}, 8},
	}

	for i, tc := range testCases {
		result := twoSumSorted(tc.nums, tc.target)
		fmt.Printf("Test %d: nums=%v, target=%d → %v\n",
			i+1, tc.nums, tc.target, result)
	}
}
```

**Explanation:**
1. Use two pointers: left at start, right at end
2. Calculate sum of elements at pointers
3. If sum equals target, return indices
4. If sum too small, move left pointer right
5. If sum too large, move right pointer left

## Problem 2: Remove Duplicates from Sorted Array

**Description:** Remove duplicates in-place from a sorted array and return the new length. Each element should appear only once.

**Requirements:**
- Modify array in-place
- Maintain relative order
- Return new length
- Don't use extra space beyond O(1)

**Examples:**
- Input: [1, 1, 2] → Output: 2, array becomes [1, 2, _]
- Input: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4] → Output: 5, array becomes [0, 1, 2, 3, 4, _, _, _, _, _]

**Constraints:**
- Array is sorted
- Time complexity: O(n)
- Space complexity: O(1)

### Solution

```go
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Slow pointer tracks position for next unique element
	slow := 0

	// Fast pointer iterates through array
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// Return length of unique elements
	return slow + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{},
		{1},
	}

	for i, nums := range testCases {
		// Create a copy to preserve original
		original := make([]int, len(nums))
		copy(original, nums)

		newLength := removeDuplicates(nums)

		fmt.Printf("Test %d: original=%v\n", i+1, original)
		fmt.Printf("         result=%v, length=%d\n", nums[:newLength], newLength)
		fmt.Println()
	}
}
```

**Explanation:**
1. Use two pointers: slow tracks unique elements, fast iterates
2. When fast finds different element, copy to slow+1 position
3. Increment slow only when unique element found
4. Return slow+1 as new length
5. Elements after new length are irrelevant