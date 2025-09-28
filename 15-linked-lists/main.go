package main

import (
	"fmt"
)

// ListNode represents a node in singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// DoublyListNode represents a node in doubly linked list
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// SinglyLinkedList represents a singly linked list
type SinglyLinkedList struct {
	Head *ListNode
	Size int
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	Head *DoublyListNode
	Tail *DoublyListNode
	Size int
}

// NewSinglyLinkedList creates a new singly linked list
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// NewDoublyLinkedList creates a new doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Singly Linked List Methods

// Append adds a node to the end of singly linked list
func (ll *SinglyLinkedList) Append(val int) {
	newNode := &ListNode{Val: val}
	ll.Size++

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

// Prepend adds a node to the beginning of singly linked list
func (ll *SinglyLinkedList) Prepend(val int) {
	newNode := &ListNode{Val: val, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

// Delete removes the first occurrence of val from singly linked list
func (ll *SinglyLinkedList) Delete(val int) bool {
	if ll.Head == nil {
		return false
	}

	// Handle head deletion
	if ll.Head.Val == val {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	current := ll.Head
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
		ll.Size--
		return true
	}

	return false
}

// Find returns true if val exists in singly linked list
func (ll *SinglyLinkedList) Find(val int) bool {
	current := ll.Head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Print displays the singly linked list
func (ll *SinglyLinkedList) Print() {
	current := ll.Head
	fmt.Print("SinglyLinkedList: ")
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// Reverse reverses the singly linked list
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

// GetMiddle returns the middle node of singly linked list
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

// HasCycle detects if singly linked list has a cycle
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

// Doubly Linked List Methods

// Append adds a node to the end of doubly linked list
func (dll *DoublyLinkedList) Append(val int) {
	newNode := &DoublyListNode{Val: val}
	dll.Size++

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
}

// Prepend adds a node to the beginning of doubly linked list
func (dll *DoublyLinkedList) Prepend(val int) {
	newNode := &DoublyListNode{Val: val}
	dll.Size++

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	newNode.Next = dll.Head
	dll.Head.Prev = newNode
	dll.Head = newNode
}

// Delete removes the first occurrence of val from doubly linked list
func (dll *DoublyLinkedList) Delete(val int) bool {
	if dll.Head == nil {
		return false
	}

	current := dll.Head

	// Find the node to delete
	for current != nil && current.Val != val {
		current = current.Next
	}

	if current == nil {
		return false
	}

	// Update pointers
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

	dll.Size--
	return true
}

// Find returns true if val exists in doubly linked list
func (dll *DoublyLinkedList) Find(val int) bool {
	current := dll.Head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Print displays the doubly linked list
func (dll *DoublyLinkedList) Print() {
	current := dll.Head
	fmt.Print("DoublyLinkedList: nil <- ")
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// PrintReverse displays the doubly linked list in reverse
func (dll *DoublyLinkedList) PrintReverse() {
	current := dll.Tail
	fmt.Print("DoublyLinkedList (reverse): nil <- ")
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.Prev
	}
	fmt.Println(" -> nil")
}

// Common Linked List Problems

// MergeTwoSortedLists merges two sorted singly linked lists
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

// RemoveNthFromEnd removes the nth node from the end of singly linked list
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// Move fast pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head // n is larger than list length
		}
		fast = fast.Next
	}

	// Move both pointers until fast reaches end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the node
	slow.Next = slow.Next.Next

	return dummy.Next
}

// AddTwoNumbers adds two numbers represented as linked lists
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

// IsPalindrome checks if singly linked list is a palindrome
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

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

// Helper function to create linked list from slice
func createLinkedList(nums []int) *ListNode {
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

// Helper function to print linked list
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("=== Linked Lists Demo ===\n")

	// 1. Singly Linked List Operations
	fmt.Println("1. Singly Linked List Operations:")
	sll := NewSinglyLinkedList()
	sll.Append(1)
	sll.Append(2)
	sll.Append(3)
	sll.Print()

	sll.Prepend(0)
	sll.Print()

	sll.Delete(2)
	sll.Print()

	fmt.Printf("Find 3: %t\n", sll.Find(3))
	fmt.Printf("Find 5: %t\n", sll.Find(5))
	fmt.Printf("Size: %d\n\n", sll.Size)

	// 2. Singly Linked List Advanced Operations
	fmt.Println("2. Singly Linked List Advanced Operations:")
	sll2 := NewSinglyLinkedList()
	for i := 1; i <= 5; i++ {
		sll2.Append(i)
	}
	sll2.Print()

	middle := sll2.GetMiddle()
	if middle != nil {
		fmt.Printf("Middle node: %d\n", middle.Val)
	}

	sll2.Reverse()
	sll2.Print()

	fmt.Printf("Has cycle: %t\n\n", sll2.HasCycle())

	// 3. Doubly Linked List Operations
	fmt.Println("3. Doubly Linked List Operations:")
	dll := NewDoublyLinkedList()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Print()

	dll.Prepend(0)
	dll.Print()

	dll.Delete(2)
	dll.Print()

	dll.PrintReverse()
	fmt.Printf("Size: %d\n\n", dll.Size)

	// 4. Merge Two Sorted Lists
	fmt.Println("4. Merge Two Sorted Lists:")
	list1 := createLinkedList([]int{1, 3, 5})
	list2 := createLinkedList([]int{2, 4, 6})
	fmt.Print("List 1: ")
	printLinkedList(list1)
	fmt.Print("List 2: ")
	printLinkedList(list2)

	merged := MergeTwoSortedLists(list1, list2)
	fmt.Print("Merged: ")
	printLinkedList(merged)
	fmt.Println()

	// 5. Remove Nth Node From End
	fmt.Println("5. Remove Nth Node From End:")
	list3 := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original: ")
	printLinkedList(list3)

	result := RemoveNthFromEnd(list3, 2)
	fmt.Print("Remove 2nd from end: ")
	printLinkedList(result)
	fmt.Println()

	// 6. Add Two Numbers
	fmt.Println("6. Add Two Numbers (as linked lists):")
	num1 := createLinkedList([]int{2, 4, 3}) // represents 342
	num2 := createLinkedList([]int{5, 6, 4}) // represents 465
	fmt.Print("Number 1: ")
	printLinkedList(num1)
	fmt.Print("Number 2: ")
	printLinkedList(num2)

	sum := AddTwoNumbers(num1, num2)
	fmt.Print("Sum (342 + 465 = 807): ")
	printLinkedList(sum)
	fmt.Println()

	// 7. Palindrome Check
	fmt.Println("7. Palindrome Check:")
	palindromeList := createLinkedList([]int{1, 2, 3, 2, 1})
	nonPalindromeList := createLinkedList([]int{1, 2, 3, 4, 5})

	fmt.Print("List: ")
	printLinkedList(palindromeList)
	fmt.Printf("Is palindrome: %t\n", IsPalindrome(palindromeList))

	fmt.Print("List: ")
	printLinkedList(nonPalindromeList)
	fmt.Printf("Is palindrome: %t\n\n", IsPalindrome(nonPalindromeList))

	fmt.Println("=== Demo Complete ===")
}
