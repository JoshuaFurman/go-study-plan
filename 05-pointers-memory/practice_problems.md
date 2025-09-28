# Practice Problems - Pointers and Memory

These practice problems reinforce pointer concepts and memory management. Focus on understanding value vs reference semantics.

## Problem 1: Linked List Reversal (In-Place)

**Description:** Reverse a singly linked list in-place using only constant extra space.

**Requirements:**
- Define a ListNode struct with Val and Next fields
- Implement a function that reverses the list
- Do not create new nodes, modify the existing structure
- Handle edge cases (empty list, single node)

**Examples:**
- Input: 1 → 2 → 3 → 4 → nil
- Output: 4 → 3 → 2 → 1 → nil

**Constraints:**
- Use iterative approach (no recursion for this problem)
- Time complexity: O(n), Space complexity: O(1)

### Solution

```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next  // Save next node
		current.Next = prev   // Reverse the link
		prev = current        // Move prev forward
		current = next        // Move current forward
	}

	return prev // New head of reversed list
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" → ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2},
		{},
		{42},
	}

	for i, nums := range testCases {
		fmt.Printf("Test %d:\n", i+1)
		list := createList(nums)
		fmt.Print("Original: ")
		printList(list)

		reversed := reverseList(list)
		fmt.Print("Reversed: ")
		printList(reversed)
		fmt.Println()
	}
}
```

**Explanation:**
1. Use three pointers: prev, current, and next
2. Iterate through list, reversing each node's Next pointer
3. prev becomes the new head when current reaches nil
4. Only constant extra space used (no new nodes created)

## Problem 2: In-Place Array Modification

**Description:** Modify an array in-place to move all zeros to the end while maintaining the relative order of non-zero elements.

**Requirements:**
- Do not create a new array
- Modify the existing array in-place
- Preserve order of non-zero elements
- Use two-pointer technique

**Examples:**
- Input: [0, 1, 0, 3, 12] → Output: [1, 3, 12, 0, 0]
- Input: [0, 0, 1] → Output: [1, 0, 0]

**Constraints:**
- Time complexity: O(n), Space complexity: O(1)

### Solution

```go
package main

import "fmt"

func moveZeros(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// Two-pointer approach
	// i tracks position for next non-zero element
	// j iterates through array
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			// Swap non-zero element to front
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

func main() {
	// Test cases
	testCases := [][]int{
		{0, 1, 0, 3, 12},
		{0, 0, 1},
		{1, 2, 3, 0, 0},
		{0, 0, 0},
		{1, 2, 3},
		{},
	}

	for i, nums := range testCases {
		fmt.Printf("Test %d:\n", i+1)
		fmt.Printf("Before: %v\n", nums)

		// Create a copy to show original
		original := make([]int, len(nums))
		copy(original, nums)

		moveZeros(nums)
		fmt.Printf("After:  %v\n", nums)
		fmt.Println()
	}
}
```

**Explanation:**
1. Use two pointers: i tracks where next non-zero should go
2. j iterates through entire array
3. When non-zero found, swap with position i and increment i
4. All zeros naturally move to end as non-zeros are moved forward
5. In-place modification with O(1) extra space