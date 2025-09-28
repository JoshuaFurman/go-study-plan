# 15. Linked Lists

This example demonstrates singly and doubly linked list implementations with common operations and LeetCode-style problems. Linked lists are fundamental data structures used extensively in coding interviews.

## Key Concepts

### Singly Linked List
- Each node has data and pointer to next node
- Head pointer tracks the start
- Operations: insert, delete, traverse, reverse
- Space: O(n), Time: O(1) for head operations, O(n) for search

### Doubly Linked List
- Each node has data, prev, and next pointers
- Head and tail pointers for efficient operations
- Bidirectional traversal
- More space but faster operations

### Common Operations
- **Insert**: Add node at beginning, end, or specific position
- **Delete**: Remove node by value or position
- **Search**: Find node by value
- **Traverse**: Visit all nodes
- **Reverse**: Reverse the list

## Data Structures

### Singly Linked List Node
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

### Doubly Linked List Node
```go
type DoublyListNode struct {
    Val  int
    Prev *DoublyListNode
    Next *DoublyListNode
}
```

## Basic Operations

### Singly Linked List Operations
```go
// Append to end
func (ll *SinglyLinkedList) Append(val int) {
    newNode := &ListNode{Val: val}
    if ll.Head == nil {
        ll.Head = newNode
        return
    }
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

// Prepend to beginning
func (ll *SinglyLinkedList) Prepend(val int) {
    newNode := &ListNode{Val: val, Next: ll.Head}
    ll.Head = newNode
}

// Delete by value
func (ll *SinglyLinkedList) Delete(val int) bool {
    if ll.Head == nil {
        return false
    }
    if ll.Head.Val == val {
        ll.Head = ll.Head.Next
        return true
    }
    current := ll.Head
    for current.Next != nil && current.Next.Val != val {
        current = current.Next
    }
    if current.Next != nil {
        current.Next = current.Next.Next
        return true
    }
    return false
}
```

### Doubly Linked List Operations
```go
// Append to end
func (dll *DoublyLinkedList) Append(val int) {
    newNode := &DoublyListNode{Val: val}
    if dll.Head == nil {
        dll.Head = newNode
        dll.Tail = newNode
        return
    }
    dll.Tail.Next = newNode
    newNode.Prev = dll.Tail
    dll.Tail = newNode
}

// Delete by value
func (dll *DoublyLinkedList) Delete(val int) bool {
    current := dll.Head
    for current != nil && current.Val != val {
        current = current.Next
    }
    if current == nil {
        return false
    }
    if current.Prev != nil {
        current.Prev.Next = current.Next
    } else {
        dll.Head = current.Next
    }
    if current.Next != nil {
        current.Next.Prev = current.Prev
    } else {
        dll.Tail = current.Prev
    }
    return true
}
```

## Advanced Operations

### Reverse Singly Linked List
```go
func (ll *SinglyLinkedList) Reverse() {
    var prev *ListNode
    current := ll.Head
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    ll.Head = prev
}
```

### Find Middle Node (Fast/Slow Pointers)
```go
func (ll *SinglyLinkedList) GetMiddle() *ListNode {
    if ll.Head == nil {
        return nil
    }
    slow, fast := ll.Head, ll.Head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```

### Detect Cycle (Floyd's Cycle Detection)
```go
func (ll *SinglyLinkedList) HasCycle() bool {
    if ll.Head == nil {
        return false
    }
    slow, fast := ll.Head, ll.Head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

## LeetCode-Style Problems

### Merge Two Sorted Lists
```go
func MergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    if l1 != nil {
        current.Next = l1
    }
    if l2 != nil {
        current.Next = l2
    }
    return dummy.Next
}
```

### Remove Nth Node From End
```go
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy
    // Move fast n+1 steps ahead
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }
    // Move both until fast reaches end
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    // Remove node
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

### Add Two Numbers
```go
func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    carry := 0
    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }
    return dummy.Next
}
```

### Palindrome Check
```go
func IsPalindrome(head *ListNode) bool {
    // Find middle
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // Reverse second half
    var prev *ListNode
    current := slow.Next
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    // Compare
    first, second := head, prev
    for second != nil {
        if first.Val != second.Val {
            return false
        }
        first = first.Next
        second = second.Next
    }
    return true
}
```

## Time Complexity

| Operation | Singly Linked List | Doubly Linked List |
|-----------|-------------------|-------------------|
| Insert Head | O(1) | O(1) |
| Insert Tail | O(n) | O(1) |
| Insert Middle | O(n) | O(n) |
| Delete Head | O(1) | O(1) |
| Delete Tail | O(n) | O(1) |
| Delete Middle | O(n) | O(n) |
| Search | O(n) | O(n) |
| Reverse | O(n) | O(n) |

## Space Complexity

- **Singly Linked List**: O(n) - stores data + next pointer
- **Doubly Linked List**: O(n) - stores data + prev + next pointers
- **Operations**: Usually O(1) auxiliary space

## When to Use Linked Lists

1. **Dynamic Size**: No need to pre-allocate space
2. **Frequent Insertions/Deletions**: Especially at beginning
3. **No Random Access**: Sequential access is fine
4. **Memory Efficiency**: Only allocate what you need

## Common Mistakes

1. **Null Pointer Dereference**: Always check for nil
2. **Lost Nodes**: Update all pointers when modifying
3. **Infinite Loops**: Ensure termination conditions
4. **Memory Leaks**: Properly handle node removal
5. **Edge Cases**: Empty lists, single nodes, head/tail operations

## Practice Problems

### Easy
- Reverse Linked List
- Merge Two Sorted Lists
- Linked List Cycle
- Remove Duplicates from Sorted List
- Palindrome Linked List

### Medium
- Add Two Numbers
- Remove Nth Node From End of List
- Rotate List
- Partition List
- Copy List with Random Pointer

### Hard
- Merge K Sorted Lists
- Reverse Nodes in k-Group
- Sort List
- LRU Cache (often implemented with doubly linked list)

## Running the Example

```bash
cd 15-linked-lists
go run main.go
```

This example provides comprehensive coverage of linked list operations and common interview problems, essential for backend software engineering interviews.