package main

import (
	"fmt"
	"time"
)

// Basic recursion: Factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Basic recursion: Fibonacci (inefficient)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Fibonacci with memoization
var fibMemo = make(map[int]int)

func fibonacciMemo(n int) int {
	if n <= 1 {
		return n
	}

	if val, exists := fibMemo[n]; exists {
		return val
	}

	result := fibonacciMemo(n-1) + fibonacciMemo(n-2)
	fibMemo[n] = result
	return result
}

// Bottom-up dynamic programming: Fibonacci
func fibonacciDP(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Classic recursion: Tower of Hanoi
func towerOfHanoi(n int, from, to, aux string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", from, to)
		return
	}

	towerOfHanoi(n-1, from, aux, to)
	fmt.Printf("Move disk %d from %s to %s\n", n, from, to)
	towerOfHanoi(n-1, aux, to, from)
}

// Recursion with backtracking: Generate all subsets
func generateSubsets(nums []int, index int, current []int, result *[][]int) {
	// Add current subset to result
	subset := make([]int, len(current))
	copy(subset, current)
	*result = append(*result, subset)

	// Try adding each remaining element
	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])
		generateSubsets(nums, i+1, current, result)
		current = current[:len(current)-1] // backtrack
	}
}

// Wrapper function for subsets
func subsets(nums []int) [][]int {
	var result [][]int
	generateSubsets(nums, 0, []int{}, &result)
	return result
}

// Dynamic Programming: Climbing Stairs
// Problem: How many ways to climb n stairs taking 1 or 2 steps at a time?
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Space-optimized climbing stairs
func climbStairsOptimized(n int) int {
	if n <= 1 {
		return 1
	}

	prev2, prev1 := 1, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// Dynamic Programming: Coin Change (Minimum coins)
// Problem: Find minimum number of coins needed to make amount
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // impossible value
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1 // impossible
	}
	return dp[amount]
}

// Dynamic Programming: Longest Common Subsequence
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// Recursion with memoization: Knapsack Problem
func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, capacity+1)
		for j := range memo[i] {
			memo[i][j] = -1 // not computed
		}
	}

	var knapsackHelper func(int, int) int
	knapsackHelper = func(index, remainingCapacity int) int {
		if index >= n || remainingCapacity <= 0 {
			return 0
		}

		if memo[index][remainingCapacity] != -1 {
			return memo[index][remainingCapacity]
		}

		// Skip current item
		result := knapsackHelper(index+1, remainingCapacity)

		// Take current item if possible
		if weights[index] <= remainingCapacity {
			result = max(result, values[index]+knapsackHelper(index+1, remainingCapacity-weights[index]))
		}

		memo[index][remainingCapacity] = result
		return result
	}

	return knapsackHelper(0, capacity)
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Recursion: Binary Tree operations (simple representation)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion: Tree traversal - Inorder
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return result
}

// Recursion: Tree height
func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}

// Recursion: Check if binary tree is balanced
func isBalanced(root *TreeNode) bool {
	var checkHeight func(*TreeNode) int
	checkHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := checkHeight(node.Left)
		if leftHeight == -1 {
			return -1
		}

		rightHeight := checkHeight(node.Right)
		if rightHeight == -1 {
			return -1
		}

		if abs(leftHeight-rightHeight) > 1 {
			return -1
		}

		return 1 + max(leftHeight, rightHeight)
	}

	return checkHeight(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("=== Basic Recursion ===")

	// Factorial
	fmt.Printf("Factorial of 5: %d\n", factorial(5))

	// Fibonacci comparison
	n := 10
	fmt.Printf("Fibonacci of %d:\n", n)

	start := time.Now()
	result1 := fibonacci(n)
	duration1 := time.Since(start)
	fmt.Printf("  Naive recursion: %d (took %v)\n", result1, duration1)

	start = time.Now()
	result2 := fibonacciMemo(n)
	duration2 := time.Since(start)
	fmt.Printf("  With memoization: %d (took %v)\n", result2, duration2)

	start = time.Now()
	result3 := fibonacciDP(n)
	duration3 := time.Since(start)
	fmt.Printf("  Bottom-up DP: %d (took %v)\n", result3, duration3)

	fmt.Println("\n=== Tower of Hanoi ===")
	fmt.Println("Tower of Hanoi with 3 disks:")
	towerOfHanoi(3, "A", "C", "B")

	fmt.Println("\n=== Subsets Generation (Backtracking) ===")
	nums := []int{1, 2, 3}
	allSubsets := subsets(nums)
	fmt.Printf("Subsets of %v: %v\n", nums, allSubsets)
	fmt.Printf("Total subsets: %d\n", len(allSubsets))

	fmt.Println("\n=== Dynamic Programming: Climbing Stairs ===")
	stairs := 5
	ways := climbStairs(stairs)
	waysOptimized := climbStairsOptimized(stairs)
	fmt.Printf("Ways to climb %d stairs: %d (standard), %d (optimized)\n", stairs, ways, waysOptimized)

	fmt.Println("\n=== Dynamic Programming: Coin Change ===")
	coins := []int{1, 2, 5}
	amount := 11
	minCoins := coinChange(coins, amount)
	fmt.Printf("Minimum coins to make %d with %v: %d\n", amount, coins, minCoins)

	fmt.Println("\n=== Dynamic Programming: Longest Common Subsequence ===")
	str1, str2 := "abcde", "ace"
	lcs := longestCommonSubsequence(str1, str2)
	fmt.Printf("LCS of '%s' and '%s': %d\n", str1, str2, lcs)

	fmt.Println("\n=== Knapsack Problem (Memoization) ===")
	weights := []int{1, 2, 3, 4}
	values := []int{10, 20, 30, 40}
	capacity := 5
	maxValue := knapsack(weights, values, capacity)
	fmt.Printf("Knapsack - weights: %v, values: %v, capacity: %d\n", weights, values, capacity)
	fmt.Printf("Maximum value: %d\n", maxValue)

	fmt.Println("\n=== Binary Tree Operations ===")

	// Create a simple binary tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	inorder := inorderTraversal(root)
	fmt.Printf("Inorder traversal: %v\n", inorder)

	height := treeHeight(root)
	fmt.Printf("Tree height: %d\n", height)

	balanced := isBalanced(root)
	fmt.Printf("Tree is balanced: %t\n", balanced)

	// Create unbalanced tree
	unbalancedRoot := &TreeNode{Val: 1}
	unbalancedRoot.Left = &TreeNode{Val: 2}
	unbalancedRoot.Left.Left = &TreeNode{Val: 3}
	unbalancedRoot.Left.Left.Left = &TreeNode{Val: 4}

	balanced2 := isBalanced(unbalancedRoot)
	fmt.Printf("Unbalanced tree is balanced: %t\n", balanced2)
}
