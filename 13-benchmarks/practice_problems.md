# Practice Problems - Benchmarks

These practice problems reinforce performance testing and benchmarking concepts. Focus on comparing algorithm efficiency and memory usage.

## Problem 1: Benchmarking Sorting Algorithms

**Description:** Compare the performance of different sorting algorithms (bubble sort, quicksort, built-in sort) on various input sizes and distributions.

**Requirements:**
- Implement multiple sorting algorithms
- Benchmark each algorithm with different input sizes
- Test with sorted, reverse-sorted, and random data
- Compare time and memory performance
- Use sub-benchmarks for different scenarios

**Test Cases:**
- Small arrays (10 elements)
- Medium arrays (1000 elements)
- Large arrays (10000 elements)
- Different data distributions

**Constraints:**
- Use Go's testing.B for benchmarks
- Reset timer for accurate measurements
- Use b.ReportAllocs() for memory tracking

### Solution

```go
package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

// BubbleSort implements bubble sort algorithm
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

// QuickSort implements quicksort algorithm
func QuickSort(arr []int) {
	if len(arr) <= 1 {
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

// Generate test data
func generateRandomSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func generateSortedSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

func generateReverseSortedSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = size - i - 1
	}
	return arr
}

// Benchmarks for BubbleSort
func BenchmarkBubbleSort_Random10(b *testing.B) {
	data := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort_Random100(b *testing.B) {
	data := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort_Sorted100(b *testing.B) {
	data := generateSortedSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		BubbleSort(arr)
	}
}

// Benchmarks for QuickSort
func BenchmarkQuickSort_Random10(b *testing.B) {
	data := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_Random100(b *testing.B) {
	data := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_Random1000(b *testing.B) {
	data := generateRandomSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		QuickSort(arr)
	}
}

// Benchmarks for Built-in Sort
func BenchmarkBuiltinSort_Random10(b *testing.B) {
	data := generateRandomSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		sort.Ints(arr)
	}
}

func BenchmarkBuiltinSort_Random100(b *testing.B) {
	data := generateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		sort.Ints(arr)
	}
}

func BenchmarkBuiltinSort_Random1000(b *testing.B) {
	data := generateRandomSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		sort.Ints(arr)
	}
}

func BenchmarkBuiltinSort_Random10000(b *testing.B) {
	data := generateRandomSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(data))
		copy(arr, data)
		sort.Ints(arr)
	}
}
```

**Explanation:**
1. Implement multiple sorting algorithms
2. Create benchmarks for different input sizes and types
3. Use b.ResetTimer() to exclude setup time
4. Copy arrays to avoid modifying original data
5. Compare performance characteristics of different algorithms

## Problem 2: String Concatenation Performance

**Description:** Compare the performance of different string concatenation methods in Go and analyze memory allocation patterns.

**Requirements:**
- Test multiple concatenation approaches
- Measure both time and memory allocations
- Use different string sizes and counts
- Demonstrate performance trade-offs

**Methods to Compare:**
- Using `+` operator
- Using `strings.Builder`
- Using `bytes.Buffer`
- Using `fmt.Sprintf`
- Using string slices with `strings.Join`

**Constraints:**
- Use b.ReportAllocs() to track memory
- Test with various string sizes
- Include warmup iterations

### Solution

```go
package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// String concatenation methods
func concatWithPlus(strings []string) string {
	result := ""
	for _, s := range strings {
		result += s
	}
	return result
}

func concatWithBuilder(strings []string) string {
	var builder strings.Builder
	for _, s := range strings {
		builder.WriteString(s)
	}
	return builder.String()
}

func concatWithBuffer(strings []string) string {
	var buffer bytes.Buffer
	for _, s := range strings {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func concatWithSprintf(strings []string) string {
	result := ""
	for _, s := range strings {
		result = fmt.Sprintf("%s%s", result, s)
	}
	return result
}

func concatWithJoin(strings []string) string {
	return strings.Join(strings, "")
}

// Benchmark functions
func BenchmarkConcat_Plus_Small(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithPlus(data)
	}
}

func BenchmarkConcat_Plus_Medium(b *testing.B) {
	data := make([]string, 100)
	for i := range data {
		data[i] = "test_string_"
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithPlus(data)
	}
}

func BenchmarkConcat_Builder_Small(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithBuilder(data)
	}
}

func BenchmarkConcat_Builder_Medium(b *testing.B) {
	data := make([]string, 100)
	for i := range data {
		data[i] = "test_string_"
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithBuilder(data)
	}
}

func BenchmarkConcat_Builder_Large(b *testing.B) {
	data := make([]string, 1000)
	for i := range data {
		data[i] = "test_string_"
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithBuilder(data)
	}
}

func BenchmarkConcat_Buffer_Small(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithBuffer(data)
	}
}

func BenchmarkConcat_Buffer_Medium(b *testing.B) {
	data := make([]string, 100)
	for i := range data {
		data[i] = "test_string_"
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithBuffer(data)
	}
}

func BenchmarkConcat_Sprintf_Small(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithSprintf(data)
	}
}

func BenchmarkConcat_Join_Small(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithJoin(data)
	}
}

func BenchmarkConcat_Join_Medium(b *testing.B) {
	data := make([]string, 100)
	for i := range data {
		data[i] = "test_string_"
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = concatWithJoin(data)
	}
}

// Memory allocation focused benchmarks
func BenchmarkConcat_MemoryComparison(b *testing.B) {
	data := make([]string, 50)
	for i := range data {
		data[i] = fmt.Sprintf("string_%d_", i)
	}

	b.Run("Plus_Operator", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = concatWithPlus(data)
		}
	})

	b.Run("Strings_Builder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = concatWithBuilder(data)
		}
	})

	b.Run("Strings_Join", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = concatWithJoin(data)
		}
	})
}
```

**Explanation:**
1. Implement multiple string concatenation methods
2. Create benchmarks for different input sizes
3. Use b.ReportAllocs() to track memory allocations
4. Compare performance characteristics
5. Demonstrate why strings.Builder is preferred for concatenation