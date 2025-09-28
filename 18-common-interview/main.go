package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Problem 1: Two Sum
// Given an array of integers and a target sum, return indices of two numbers that add up to target
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return nil
}

// Problem 2: Valid Parentheses
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid
func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != bracketMap[char] {
				return false
			}
		}
	}

	return len(stack) == 0
}

// Problem 3: Merge Two Sorted Lists
// Merge two sorted linked lists and return it as a sorted list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

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

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

// Problem 4: Maximum Subarray (Kadane's Algorithm)
// Given an integer array nums, find the contiguous subarray with the largest sum
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxCurrent := nums[0]
	maxGlobal := nums[0]

	for i := 1; i < len(nums); i++ {
		maxCurrent = max(nums[i], maxCurrent+nums[i])
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}

	return maxGlobal
}

// Problem 5: Climbing Stairs
// You are climbing a staircase. It takes n steps to reach the top. Each time you can climb 1 or 2 steps.
// In how many distinct ways can you climb to the top?
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Problem 6: Binary Tree Maximum Path Sum
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge.
// Find the maximum path sum in the tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))

		currentSum := node.Val + left + right
		maxSum = max(maxSum, currentSum)

		return node.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}

// Problem 7: Longest Palindromic Substring
// Given a string s, return the longest palindromic substring in s
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Check for odd length palindromes
		len1 := expandAroundCenter(s, i, i)
		// Check for even length palindromes
		len2 := expandAroundCenter(s, i, i+1)

		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// Problem 8: Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxX := math.MinInt32
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}

		minX := math.MaxInt32
		if partitionX < x {
			minX = nums1[partitionX]
		}

		maxY := math.MinInt32
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}

		minY := math.MaxInt32
		if partitionY < y {
			minY = nums2[partitionY]
		}

		if maxX <= minY && maxY <= minX {
			if (x+y)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2.0
			} else {
				return float64(max(maxX, maxY))
			}
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}

// Problem 9: Regular Expression Matching
// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*'
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	// Handle patterns like a*, a*b*, etc.
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] // Zero occurrence
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j] // One or more occurrences
				}
			}
		}
	}

	return dp[m][n]
}

// Problem 10: Word Break
// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// determine if s can be segmented into a space-separated sequence of one or more dictionary words
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// Problem 11: Number of Islands
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	islands := 0

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0' // Mark as visited

		// Visit all adjacent cells
		dfs(r-1, c) // Up
		dfs(r+1, c) // Down
		dfs(r, c-1) // Left
		dfs(r, c+1) // Right
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(i, j)
			}
		}
	}

	return islands
}

// Problem 12: Longest Increasing Subsequence
// Given an integer array nums, return the length of the longest strictly increasing subsequence
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper to create linked list from slice
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

// Helper to convert linked list to string
func linkedListToString(head *ListNode) string {
	var result []string
	current := head
	for current != nil {
		result = append(result, strconv.Itoa(current.Val))
		current = current.Next
	}
	return strings.Join(result, " -> ")
}

func main() {
	fmt.Println("=== Common Interview Problems Demo ===\n")

	// Problem 1: Two Sum
	fmt.Println("1. Two Sum:")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Array: %v, Target: %d\n", nums, target)
	fmt.Printf("Indices: %v (values: %d, %d)\n\n", result, nums[result[0]], nums[result[1]])

	// Problem 2: Valid Parentheses
	fmt.Println("2. Valid Parentheses:")
	testStrings := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, s := range testStrings {
		fmt.Printf("'%s' is valid: %t\n", s, isValid(s))
	}
	fmt.Println()

	// Problem 3: Merge Two Sorted Lists
	fmt.Println("3. Merge Two Sorted Lists:")
	list1 := createLinkedList([]int{1, 2, 4})
	list2 := createLinkedList([]int{1, 3, 4})
	fmt.Printf("List 1: %s\n", linkedListToString(list1))
	fmt.Printf("List 2: %s\n", linkedListToString(list2))
	merged := mergeTwoLists(list1, list2)
	fmt.Printf("Merged: %s\n\n", linkedListToString(merged))

	// Problem 4: Maximum Subarray
	fmt.Println("4. Maximum Subarray:")
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Max subarray sum: %d\n\n", maxSubArray(arr))

	// Problem 5: Climbing Stairs
	fmt.Println("5. Climbing Stairs:")
	steps := 5
	fmt.Printf("Ways to climb %d stairs: %d\n\n", steps, climbStairs(steps))

	// Problem 6: Binary Tree Maximum Path Sum
	fmt.Println("6. Binary Tree Maximum Path Sum:")
	// Create a simple tree: 1 -> (2, 3)
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Printf("Tree path sum: %d\n\n", maxPathSum(root))

	// Problem 7: Longest Palindromic Substring
	fmt.Println("7. Longest Palindromic Substring:")
	palindromeStr := "babad"
	fmt.Printf("String: %s\n", palindromeStr)
	fmt.Printf("Longest palindrome: %s\n\n", longestPalindrome(palindromeStr))

	// Problem 8: Median of Two Sorted Arrays
	fmt.Println("8. Median of Two Sorted Arrays:")
	arr1 := []int{1, 3}
	arr2 := []int{2}
	fmt.Printf("Arrays: %v and %v\n", arr1, arr2)
	fmt.Printf("Median: %.1f\n\n", findMedianSortedArrays(arr1, arr2))

	// Problem 9: Regular Expression Matching
	fmt.Println("9. Regular Expression Matching:")
	patterns := []string{"aa", "a", "ab", "a*b"}
	for _, pattern := range patterns {
		match := isMatch("aa", pattern)
		fmt.Printf("'%s' matches 'aa': %t\n", pattern, match)
	}
	fmt.Println()

	// Problem 10: Word Break
	fmt.Println("10. Word Break:")
	sentence := "leetcode"
	dict := []string{"leet", "code"}
	fmt.Printf("String: %s, Dict: %v\n", sentence, dict)
	fmt.Printf("Can break: %t\n\n", wordBreak(sentence, dict))

	// Problem 11: Number of Islands
	fmt.Println("11. Number of Islands:")
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Printf("Grid has %d islands\n\n", numIslands(grid))

	// Problem 12: Longest Increasing Subsequence
	fmt.Println("12. Longest Increasing Subsequence:")
	lisArr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("Array: %v\n", lisArr)
	fmt.Printf("LIS length: %d\n\n", lengthOfLIS(lisArr))

	fmt.Println("=== Demo Complete ===")
}
