package main

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkFibonacciRecursive benchmarks recursive fibonacci
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(20)
	}
}

// BenchmarkFibonacciIterative benchmarks iterative fibonacci
func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(20)
	}
}

// BenchmarkFibonacciMemoized benchmarks memoized fibonacci
func BenchmarkFibonacciMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMemoized(20)
	}
}

// BenchmarkStringConcatenationSlow benchmarks slow string concatenation
func BenchmarkStringConcatenationSlow(b *testing.B) {
	parts := []string{"Hello", ", ", "World", "!", " This", " is", " a", " test", "."}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringConcatenationSlow(parts)
	}
}

// BenchmarkStringConcatenationFast benchmarks fast string concatenation
func BenchmarkStringConcatenationFast(b *testing.B) {
	parts := []string{"Hello", ", ", "World", "!", " This", " is", " a", " test", "."}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringConcatenationFast(parts)
	}
}

// BenchmarkBubbleSort benchmarks bubble sort
func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 45, 72, 18}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BubbleSort(copyArr)
	}
}

// BenchmarkQuickSort benchmarks quicksort
func BenchmarkQuickSort(b *testing.B) {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 45, 72, 18}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		QuickSort(copyArr)
	}
}

// BenchmarkLinearSearch benchmarks linear search
func BenchmarkLinearSearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	target := 500
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(arr, target)
	}
}

// BenchmarkBinarySearch benchmarks binary search
func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	target := 500
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target)
	}
}

// BenchmarkSieveOfEratosthenes benchmarks sieve of Eratosthenes
func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SieveOfEratosthenes(1000)
	}
}

// BenchmarkTrialDivision benchmarks trial division prime generation
func BenchmarkTrialDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrialDivision(1000)
	}
}

// BenchmarkFactorialRecursive benchmarks recursive factorial
func BenchmarkFactorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialRecursive(15)
	}
}

// BenchmarkFactorialIterative benchmarks iterative factorial
func BenchmarkFactorialIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialIterative(15)
	}
}

// BenchmarkIsPalindromeSlow benchmarks slow palindrome check
func BenchmarkIsPalindromeSlow(b *testing.B) {
	testString := "A man, a plan, a canal: Panama"
	cleaned := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(testString, " ", ""), ",", ""), ":", ""))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeSlow(cleaned)
	}
}

// BenchmarkIsPalindromeFast benchmarks fast palindrome check
func BenchmarkIsPalindromeFast(b *testing.B) {
	testString := "A man, a plan, a canal: Panama"
	cleaned := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(testString, " ", ""), ",", ""), ":", ""))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeFast(cleaned)
	}
}

// BenchmarkMatrixMultiplicationNaive benchmarks naive matrix multiplication
func BenchmarkMatrixMultiplicationNaive(b *testing.B) {
	size := 10
	a := make([][]int, size)
	b_mat := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
		b_mat[i] = make([]int, size)
		for j := range a[i] {
			a[i][j] = i + j
			b_mat[i][j] = i - j
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MatrixMultiplicationNaive(a, b_mat)
	}
}

// Benchmark with different input sizes for Fibonacci
func BenchmarkFibonacciRecursive10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(10)
	}
}

func BenchmarkFibonacciRecursive20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(20)
	}
}

func BenchmarkFibonacciRecursive30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(30)
	}
}

// Sub-benchmarks for different sorting algorithms
func BenchmarkSorting(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("BubbleSort_%d", size), func(b *testing.B) {
			arr := make([]int, size)
			for i := range arr {
				arr[i] = size - i // Reverse sorted
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				copyArr := make([]int, len(arr))
				copy(copyArr, arr)
				BubbleSort(copyArr)
			}
		})

		b.Run(fmt.Sprintf("QuickSort_%d", size), func(b *testing.B) {
			arr := make([]int, size)
			for i := range arr {
				arr[i] = size - i // Reverse sorted
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				copyArr := make([]int, len(arr))
				copy(copyArr, arr)
				QuickSort(copyArr)
			}
		})
	}
}
