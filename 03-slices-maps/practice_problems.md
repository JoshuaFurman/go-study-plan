# Practice Problems - Slices and Maps

These practice problems reinforce slice and map operations. Focus on efficient algorithms and proper Go idioms.

## Problem 1: Array Deduplication

**Description:** Write a function that removes duplicate elements from a slice of integers while preserving the original order. Return a new slice with unique elements.

**Examples:**
- Input: [1, 2, 2, 3, 4, 4, 5] → Output: [1, 2, 3, 4, 5]
- Input: [1, 1, 1, 1] → Output: [1]
- Input: [] → Output: []

**Requirements:**
- Preserve original order
- Handle empty slices
- Use a map for efficient duplicate detection

### Solution

```go
package main

import "fmt"

func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	seen := make(map[int]bool)
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 2, 3, 4, 4, 5},
		{1, 1, 1, 1},
		{},
		{5},
		{3, 1, 4, 1, 5, 9, 2, 6, 5},
	}

	for i, test := range testCases {
		result := removeDuplicates(test)
		fmt.Printf("Test %d: %v → %v\n", i+1, test, result)
	}
}
```

**Explanation:**
1. Use a map to track seen elements (O(1) lookup)
2. Iterate through original slice, only append unseen elements
3. Pre-allocate result slice capacity for efficiency
4. Handle edge cases (empty slice, single element)

## Problem 2: Word Frequency Counter

**Description:** Implement a word frequency counter that:
- Takes a slice of strings (words)
- Returns a map[string]int with word frequencies
- Ignores case sensitivity
- Handles punctuation (basic cleaning)

**Examples:**
- Input: ["Hello", "world", "Hello"] → Output: {"hello": 2, "world": 1}
- Input: ["Go", "is", "great", "go", "GO"] → Output: {"go": 3, "is": 1, "great": 1}

**Requirements:**
- Case insensitive counting
- Basic punctuation removal (commas, periods)
- Return sorted keys for consistent output

### Solution

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

func wordFrequency(words []string) map[string]int {
	frequency := make(map[string]int)

	for _, word := range words {
		// Clean the word: convert to lowercase and remove basic punctuation
		cleanWord := strings.ToLower(word)
		cleanWord = strings.Trim(cleanWord, ".,!?;:\"'")

		if cleanWord != "" {
			frequency[cleanWord]++
		}
	}

	return frequency
}

func printSortedFrequency(freq map[string]int) {
	// Get sorted keys for consistent output
	keys := make([]string, 0, len(freq))
	for word := range freq {
		keys = append(keys, word)
	}
	sort.Strings(keys)

	fmt.Println("Word frequencies:")
	for _, word := range keys {
		fmt.Printf("  %s: %d\n", word, freq[word])
	}
}

func main() {
	// Test cases
	testCases := [][]string{
		{"Hello", "world", "Hello"},
		{"Go", "is", "great", "go", "GO"},
		{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog."},
		{"One", "two", "one", "three", "two", "one"},
	}

	for i, test := range testCases {
		fmt.Printf("\nTest %d: %v\n", i+1, test)
		freq := wordFrequency(test)
		printSortedFrequency(freq)
	}
}
```

**Explanation:**
1. Use map[string]int to store word frequencies
2. Clean each word: lowercase and remove punctuation
3. Skip empty strings after cleaning
4. Sort keys for consistent output display
5. Handle various punctuation and case variations