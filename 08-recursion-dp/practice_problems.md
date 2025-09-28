# Practice Problems - Recursion and Dynamic Programming

These practice problems reinforce recursive thinking and dynamic programming concepts. Focus on identifying base cases and optimal substructure.

## Problem 1: Fibonacci with Memoization

**Description:** Implement Fibonacci sequence calculation with memoization to avoid redundant calculations. Compare performance with naive recursive approach.

**Requirements:**
- Implement both naive recursive and memoized versions
- Use a map for memoization cache
- Handle large inputs efficiently
- Return the nth Fibonacci number

**Examples:**
- fibonacci(0) → 0
- fibonacci(1) → 1
- fibonacci(10) → 55

**Constraints:**
- 0 ≤ n ≤ 40 (to avoid integer overflow)
- Memoized version should be O(n) time and space

### Solution

```go
package main

import (
	"fmt"
	"time"
)

// Naive recursive approach - exponential time
func fibonacciNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciNaive(n-1) + fibonacciNaive(n-2)
}

// Memoized recursive approach - linear time
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	memo[n] = result
	return result
}

// Wrapper function for memoized version
func fibonacciMemoized(n int) int {
	memo := make(map[int]int)
	return fibonacciMemo(n, memo)
}

func timeFunction(fn func(int) int, n int, label string) {
	start := time.Now()
	result := fn(n)
	duration := time.Since(start)

	fmt.Printf("%s(%d) = %d (took %v)\n", label, n, result, duration)
}

func main() {
	n := 35 // Large enough to show performance difference

	fmt.Println("Comparing Fibonacci implementations:")
	fmt.Println("===================================")

	// Test small values first
	fmt.Printf("fibonacciMemoized(10) = %d\n", fibonacciMemoized(10))
	fmt.Printf("fibonacciNaive(10) = %d\n", fibonacciNaive(10))
	fmt.Println()

	// Performance comparison
	fmt.Println("Performance comparison for n=35:")
	timeFunction(fibonacciMemoized, n, "Memoized")
	timeFunction(fibonacciNaive, n, "Naive   ")
}
```

**Explanation:**
1. Naive version recalculates same values exponentially
2. Memoized version stores results in map to avoid recalculation
3. Performance difference becomes dramatic for larger n
4. Demonstrates trade-off between time and space complexity

## Problem 2: Coin Change (Minimum Coins)

**Description:** Given coins of different denominations and a total amount, find the minimum number of coins needed to make that amount. Return -1 if impossible.

**Requirements:**
- Use dynamic programming approach
- Handle cases where amount cannot be made
- Optimize for space where possible
- Return minimum coin count or -1

**Examples:**
- coins: [1, 2, 5], amount: 11 → 3 (5+5+1)
- coins: [2], amount: 3 → -1 (impossible)

**Constraints:**
- 1 ≤ coins.length ≤ 12
- 1 ≤ coins[i] ≤ 2^31 - 1
- 0 ≤ amount ≤ 10^4

### Solution

```go
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i] represents minimum coins needed for amount i
	dp := make([]int, amount+1)

	// Initialize dp array with maximum value
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 // 0 coins needed for amount 0

	// Fill dp array
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin] != math.MaxInt32 {
					dp[i] = min(dp[i], dp[i-coin]+1)
				}
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2}, 3},
		{[]int{1}, 0},
		{[]int{1, 2, 5}, 100},
		{[]int{3, 7, 405, 436}, 8839},
	}

	for i, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)
		fmt.Printf("Test %d: coins=%v, amount=%d → %d coins\n",
			i+1, tc.coins, tc.amount, result)
	}
}
```

**Explanation:**
1. dp[i] stores minimum coins needed for amount i
2. Initialize with MaxInt32, except dp[0] = 0
3. For each amount, try each coin and update dp[i] if better
4. Return -1 if dp[amount] still MaxInt32 (impossible)
5. Demonstrates bottom-up dynamic programming approach