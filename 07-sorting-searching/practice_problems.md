# Practice Problems - Sorting and Searching

These practice problems reinforce sorting and searching algorithms. Focus on implementing efficient algorithms and using Go's sort package effectively.

## Problem 1: Custom Sorting with Multiple Criteria

**Description:** Implement a Person struct that can be sorted by multiple criteria: first by age (ascending), then by name (alphabetical) for same ages.

**Requirements:**
- Define Person struct with Name and Age fields
- Implement sort.Interface methods (Len, Less, Swap)
- Sort primarily by age, secondarily by name
- Handle edge cases (empty slice, same ages)

**Examples:**
- Input: [{Alice, 25}, {Bob, 20}, {Charlie, 25}]
- Output: [{Bob, 20}, {Alice, 25}, {Charlie, 25}]

**Constraints:**
- Use sort.Sort function
- Stable sort behavior for same ages

### Solution

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAgeThenName []Person

func (p ByAgeThenName) Len() int {
	return len(p)
}

func (p ByAgeThenName) Less(i, j int) bool {
	// First compare by age
	if p[i].Age != p[j].Age {
		return p[i].Age < p[j].Age
	}
	// If ages are equal, compare by name
	return p[i].Name < p[j].Name
}

func (p ByAgeThenName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 25},
		{"Diana", 30},
		{"Eve", 20},
		{"Frank", 25},
	}

	fmt.Println("Before sorting:")
	for _, p := range people {
		fmt.Printf("  %s (%d)\n", p.Name, p.Age)
	}

	sort.Sort(ByAgeThenName(people))

	fmt.Println("\nAfter sorting (by age, then name):")
	for _, p := range people {
		fmt.Printf("  %s (%d)\n", p.Name, p.Age)
	}
}
```

**Explanation:**
1. Define Person struct with Name and Age
2. Create ByAgeThenName type that implements sort.Interface
3. Less method compares age first, then name for tie-breaking
4. Use sort.Sort to sort the slice
5. Demonstrates multi-criteria sorting

## Problem 2: Binary Search with Bounds

**Description:** Implement binary search that finds the leftmost and rightmost positions of a target value in a sorted array.

**Requirements:**
- Return the range [left, right] where target appears
- If target not found, return [-1, -1]
- Handle duplicate values correctly
- Use iterative approach

**Examples:**
- Input: [1, 2, 3, 3, 3, 4, 5], target: 3 → [2, 4]
- Input: [1, 2, 3, 4, 5], target: 6 → [-1, -1]

**Constraints:**
- Array is sorted in ascending order
- Time complexity: O(log n)

### Solution

```go
package main

import "fmt"

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}

	last := findLast(nums, target)
	return []int{first, last}
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 3, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{}, 1},
	}

	for i, tc := range testCases {
		result := searchRange(tc.nums, tc.target)
		fmt.Printf("Test %d: nums=%v, target=%d → [%d, %d]\n",
			i+1, tc.nums, tc.target, result[0], result[1])
	}
}
```

**Explanation:**
1. findFirst: Searches for leftmost occurrence by continuing left when target found
2. findLast: Searches for rightmost occurrence by continuing right when target found
3. searchRange: Combines both to return the range
4. Handles edge cases (empty array, target not found, all duplicates)
5. Uses iterative binary search for efficiency