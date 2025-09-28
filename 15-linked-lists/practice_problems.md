# Practice Problems - Linked Lists

These practice problems reinforce linked list operations and pointer manipulation. Focus on common interview problems with linked lists.

## Problem 1: Reverse Linked List

**Description:** Reverse a singly linked list and return the new head. Do this in-place with O(1) space complexity.

**Requirements:**
- Reverse the list iteratively (no recursion)
- Modify the existing nodes, don't create new ones
- Handle edge cases (empty list, single node)
- Return the new head of reversed list

**Examples:**
- Input: 1 → 2 → 3 → 4 → nil
- Output: 4 → 3 → 2 → 1 → nil

**Constraints:**
- Time complexity: O(n)
- Space complexity: O(1)

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

	return prev // New head
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
1. Use three pointers: prev, current, next
2. Iterate through list, reversing each Next pointer
3. prev becomes the new head when current reaches nil
4. Only constant extra space used

## Problem 2: Detect Cycle in Linked List

**Description:** Determine if a linked list has a cycle. Return true if there is a cycle, false otherwise.

**Requirements:**
- Use Floyd's cycle detection algorithm (tortoise and hare)
- Handle edge cases (empty list, single node, no cycle)
- Don't modify the list
- Use O(1) space

**Examples:**
- List with cycle: 1 → 2 → 3 → 2 (cycle back to 2) → true
- List without cycle: 1 → 2 → 3 → nil → false

**Constraints:**
- Time complexity: O(n)
- Space complexity: O(1)

### Solution

```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head      // Tortoise: moves 1 step
	fast := head.Next // Hare: moves 2 steps

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true // Cycle detected
		}

		slow = slow.Next      // Move tortoise 1 step
		fast = fast.Next.Next // Move hare 2 steps
	}

	return false // No cycle
}

func createListWithCycle(nums []int, cyclePos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	var cycleNode *ListNode

	// Create the list
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next

		// Mark the cycle position
		if i == cyclePos {
			cycleNode = current
		}
	}

	// Create cycle if cyclePos is valid
	if cycleNode != nil && cyclePos >= 0 {
		current.Next = cycleNode
	}

	return head
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		cyclePos int // -1 means no cycle
		expected bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},  // Cycle at position 1
		{[]int{1, 2}, 0, true},         // Cycle at position 0
		{[]int{1}, -1, false},          // No cycle
		{[]int{}, -1, false},           // Empty list
		{[]int{1, 2, 3, 4}, -1, false}, // No cycle
	}

	for i, tc := range testCases {
		list := createListWithCycle(tc.nums, tc.cyclePos)
		result := hasCycle(list)

		fmt.Printf("Test %d: nums=%v, cyclePos=%d → hasCycle=%t (expected %t)\n",
			i+1, tc.nums, tc.cyclePos, result, tc.expected)

		if result != tc.expected {
			fmt.Println("  ❌ MISMATCH!")
		} else {
			fmt.Println("  ✅ OK")
		}
	}
}
```

**Explanation:**
1. Use two pointers: slow (tortoise) and fast (hare)
2. Slow moves 1 step, fast moves 2 steps per iteration
3. If there's a cycle, fast will eventually catch up to slow
4. If fast reaches nil, no cycle exists
5. O(1) space, O(n) time complexity