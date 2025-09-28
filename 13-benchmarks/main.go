package main

import (
	"fmt"
	"strings"
)

// FibonacciRecursive calculates fibonacci using recursion (inefficient)
func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciIterative calculates fibonacci using iteration (efficient)
func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// FibonacciMemoized calculates fibonacci using memoization
func FibonacciMemoized(n int) int {
	memo := make(map[int]int)
	var fib func(int) int
	fib = func(x int) int {
		if x <= 1 {
			return x
		}
		if val, exists := memo[x]; exists {
			return val
		}
		memo[x] = fib(x-1) + fib(x-2)
		return memo[x]
	}
	return fib(n)
}

// StringConcatenationSlow concatenates strings using + operator
func StringConcatenationSlow(parts []string) string {
	result := ""
	for _, part := range parts {
		result += part
	}
	return result
}

// StringConcatenationFast concatenates strings using strings.Builder
func StringConcatenationFast(parts []string) string {
	var builder strings.Builder
	for _, part := range parts {
		builder.WriteString(part)
	}
	return builder.String()
}

// BubbleSort sorts slice using bubble sort (O(n²))
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// QuickSort sorts slice using quicksort (O(n log n))
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	QuickSort(arr[:right+1])
	QuickSort(arr[left:])
}

// LinearSearch searches for target in slice (O(n))
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// BinarySearch searches for target in sorted slice (O(log n))
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// SieveOfEratosthenes finds all primes up to n
func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}
	primes[0], primes[1] = false, false

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	var result []int
	for i, isPrime := range primes {
		if isPrime {
			result = append(result, i)
		}
	}
	return result
}

// TrialDivision finds all primes up to n using trial division
func TrialDivision(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// MatrixMultiplicationNaive performs naive matrix multiplication (O(n³))
func MatrixMultiplicationNaive(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// FactorialRecursive calculates factorial recursively
func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

// FactorialIterative calculates factorial iteratively
func FactorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// IsPalindromeSlow checks if string is palindrome using string reversal
func IsPalindromeSlow(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) == s
}

// IsPalindromeFast checks if string is palindrome using two pointers
func IsPalindromeFast(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Performance Comparison Demo ===\n")

	// Fibonacci comparison
	fmt.Println("1. Fibonacci Calculation:")
	n := 20
	fmt.Printf("Fibonacci(%d):\n", n)
	fmt.Printf("  Recursive: %d\n", FibonacciRecursive(n))
	fmt.Printf("  Iterative: %d\n", FibonacciIterative(n))
	fmt.Printf("  Memoized:  %d\n\n", FibonacciMemoized(n))

	// String concatenation comparison
	fmt.Println("2. String Concatenation:")
	parts := []string{"Hello", ", ", "World", "!", " This", " is", " a", " test", "."}
	fmt.Printf("Concatenating %d strings:\n", len(parts))
	slow := StringConcatenationSlow(parts)
	fast := StringConcatenationFast(parts)
	fmt.Printf("  Slow result length: %d\n", len(slow))
	fmt.Printf("  Fast result length: %d\n", len(fast))
	fmt.Printf("  Results match: %t\n\n", slow == fast)

	// Sorting comparison
	fmt.Println("3. Sorting Algorithms:")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original array: %v\n", arr)

	bubbleArr := make([]int, len(arr))
	copy(bubbleArr, arr)
	BubbleSort(bubbleArr)
	fmt.Printf("  Bubble sort: %v\n", bubbleArr)

	quickArr := make([]int, len(arr))
	copy(quickArr, arr)
	QuickSort(quickArr)
	fmt.Printf("  Quick sort:  %v\n\n", quickArr)

	// Search comparison
	fmt.Println("4. Search Algorithms:")
	sortedArr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	fmt.Printf("Searching for %d in %v:\n", target, sortedArr)
	fmt.Printf("  Linear search: index %d\n", LinearSearch(sortedArr, target))
	fmt.Printf("  Binary search: index %d\n\n", BinarySearch(sortedArr, target))

	// Prime generation comparison
	fmt.Println("5. Prime Generation:")
	limit := 100
	fmt.Printf("Primes up to %d:\n", limit)
	sievePrimes := SieveOfEratosthenes(limit)
	trialPrimes := TrialDivision(limit)
	fmt.Printf("  Sieve method: %d primes\n", len(sievePrimes))
	fmt.Printf("  Trial method: %d primes\n", len(trialPrimes))
	fmt.Printf("  Results match: %t\n\n", fmt.Sprintf("%v", sievePrimes) == fmt.Sprintf("%v", trialPrimes))

	// Factorial comparison
	fmt.Println("6. Factorial Calculation:")
	factN := 10
	fmt.Printf("Factorial(%d):\n", factN)
	fmt.Printf("  Recursive: %d\n", FactorialRecursive(factN))
	fmt.Printf("  Iterative: %d\n\n", FactorialIterative(factN))

	// Palindrome comparison
	fmt.Println("7. Palindrome Check:")
	testString := "A man, a plan, a canal: Panama"
	// Clean the string for palindrome check
	cleaned := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(testString, " ", ""), ",", ""), ":", ""))
	fmt.Printf("Testing: %s\n", testString)
	fmt.Printf("  Slow method: %t\n", IsPalindromeSlow(cleaned))
	fmt.Printf("  Fast method: %t\n\n", IsPalindromeFast(cleaned))

	fmt.Println("Run 'go test -bench=.' to see benchmark results!")
}
