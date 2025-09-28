# 08 - Recursion and Dynamic Programming

This example demonstrates recursive algorithms, memoization techniques, and basic dynamic programming problems - essential concepts for technical interviews.

## Key Concepts

### Basic Recursion

#### Factorial
```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```
- Base case: when to stop recursion
- Recursive case: how to break down the problem
- Call stack builds up, then unwinds

#### Fibonacci (Naive)
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
- Exponential time complexity: O(2^n)
- Many redundant calculations
- Good example of inefficient recursion

### Memoization

#### Fibonacci with Memoization
```go
var fibMemo = make(map[int]int)

func fibonacciMemo(n int) int {
    if val, exists := fibMemo[n]; exists {
        return val
    }
    result := fibonacciMemo(n-1) + fibonacciMemo(n-2)
    fibMemo[n] = result
    return result
}
```
- Cache results to avoid redundant calculations
- Time complexity: O(n)
- Space complexity: O(n) for cache

### Bottom-Up Dynamic Programming

#### Fibonacci DP
```go
func fibonacciDP(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1

    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```
- Build solution from smaller subproblems
- Often more efficient than memoization
- No recursion overhead

### Classic Recursion Problems

#### Tower of Hanoi
```go
func towerOfHanoi(n int, from, to, aux string) {
    if n == 1 {
        fmt.Printf("Move disk 1 from %s to %s\n", from, to)
        return
    }
    towerOfHanoi(n-1, from, aux, to)
    fmt.Printf("Move disk %d from %s to %s\n", n, from, to)
    towerOfHanoi(n-1, aux, to, from)
}
```
- Divide and conquer approach
- Time complexity: O(2^n)

#### Subset Generation (Backtracking)
```go
func generateSubsets(nums []int, index int, current []int, result *[][]int) {
    // Add current subset
    *result = append(*result, append([]int{}, current...))

    for i := index; i < len(nums); i++ {
        current = append(current, nums[i])
        generateSubsets(nums, i+1, current, result)
        current = current[:len(current)-1] // backtrack
    }
}
```
- Generate all possible combinations
- Backtracking: try, recurse, undo

### Dynamic Programming Problems

#### Climbing Stairs
```go
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1

    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```
- dp[i] = ways to reach step i
- Each step can come from i-1 or i-2

#### Coin Change (Minimum Coins)
```go
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0

    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i], dp[i-coin]+1)
            }
        }
    }
    return dp[amount]
}
```
- dp[i] = minimum coins to make amount i
- Consider each coin for each amount

#### Longest Common Subsequence
```go
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
```
- 2D DP table
- Compare characters, build up solution

### Knapsack Problem (Memoization)

```go
func knapsack(weights, values []int, capacity int) int {
    memo := make([][]int, len(weights))
    for i := range memo {
        memo[i] = make([]int, capacity+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var helper func(int, int) int
    helper = func(index, remaining int) int {
        if index >= len(weights) || remaining <= 0 {
            return 0
        }
        if memo[index][remaining] != -1 {
            return memo[index][remaining]
        }

        // Skip item
        result := helper(index+1, remaining)

        // Take item if possible
        if weights[index] <= remaining {
            result = max(result, values[index] + helper(index+1, remaining-weights[index]))
        }

        memo[index][remaining] = result
        return result
    }

    return helper(0, capacity)
}
```

### Tree Recursion

#### Binary Tree Height
```go
func treeHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}
```

#### Check Balanced Tree
```go
func isBalanced(root *TreeNode) bool {
    var checkHeight func(*TreeNode) int
    checkHeight = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left := checkHeight(node.Left)
        right := checkHeight(node.Right)

        if left == -1 || right == -1 || abs(left-right) > 1 {
            return -1
        }

        return 1 + max(left, right)
    }

    return checkHeight(root) != -1
}
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
=== Basic Recursion ===
Factorial of 5: 120
Fibonacci of 10:
  Naive recursion: 55 (took ~instant)
  With memoization: 55 (took ~instant)
  Bottom-up DP: 55 (took ~instant)

=== Tower of Hanoi ===
Tower of Hanoi with 3 disks:
Move disk 1 from A to C
Move disk 2 from A to B
Move disk 1 from C to B
Move disk 3 from A to C
Move disk 1 from B to A
Move disk 2 from B to C
Move disk 1 from A to C

=== Subsets Generation (Backtracking) ===
Subsets of [1 2 3]: [[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3] []]
Total subsets: 8

=== Dynamic Programming: Climbing Stairs ===
Ways to climb 5 stairs: 8 (standard), 8 (optimized)

=== Dynamic Programming: Coin Change ===
Minimum coins to make 11 with [1 2 5]: 3

=== Dynamic Programming: Longest Common Subsequence ===
LCS of 'abcde' and 'ace': 3

=== Knapsack Problem (Memoization) ===
Knapsack - weights: [1 2 3 4], values: [10 20 30 40], capacity: 5
Maximum value: 50

=== Binary Tree Operations ===
Inorder traversal: [4 2 5 1 3]
Tree height: 3
Tree is balanced: true
Unbalanced tree is balanced: false
```

## Complexity Analysis

| Problem | Recursion | Memoization | DP | Time | Space |
|---------|-----------|-------------|----|------|-------|
| Fibonacci | O(2^n) | O(n) | O(n) | O(n) | O(n) |
| Climbing Stairs | O(2^n) | O(n) | O(n) | O(n) | O(n) |
| Coin Change | O(2^n) | O(amount) | O(coins×amount) | O(coins×amount) | O(amount) |
| LCS | O(2^(m+n)) | O(m×n) | O(m×n) | O(m×n) | O(m×n) |
| Knapsack | O(2^n) | O(n×capacity) | O(n×capacity) | O(n×capacity) | O(n×capacity) |

## Best Practices

### ✅ Do's
- **Identify base cases** clearly
- **Use memoization** for overlapping subproblems
- **Consider bottom-up DP** for better performance
- **Draw recursion trees** to understand call patterns
- **Be aware of stack limits** (default ~1MB in Go)

### ❌ Don'ts
- **Don't use recursion** for problems with deep call stacks
- **Don't forget base cases** (infinite recursion)
- **Don't modify shared state** in recursive functions
- **Don't use recursion** when iteration is simpler
- **Don't ignore time/space complexity**

### When to Use Each Approach

1. **Recursion**: Natural for tree/graph problems, divide-and-conquer
2. **Memoization**: When recursion has overlapping subproblems
3. **Bottom-up DP**: When you can build solution iteratively
4. **Backtracking**: When exploring all possibilities (subsets, permutations)

## Common Interview Questions

- Implement Fibonacci with different approaches
- Solve the climbing stairs problem
- Find minimum coins for amount
- Generate all subsets of a set
- Solve knapsack problem
- Check if binary tree is balanced

## Next Steps

- Practice more DP problems (edit distance, longest increasing subsequence)
- Implement recursive tree traversals
- Solve backtracking problems (N-Queens, Sudoku)
- Compare recursive vs iterative solutions
- Profile recursive functions for stack usage