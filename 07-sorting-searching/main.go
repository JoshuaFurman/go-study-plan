package main

import (
	"fmt"
	"sort"
)

// Person struct for custom sorting
type Person struct {
	Name string
	Age  int
}

// People slice type for custom sorting
type People []Person

// Implement sort.Interface for People
func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Product struct for sorting by multiple criteria
type Product struct {
	Name  string
	Price float64
	Stock int
}

// ByPrice implements sort.Interface for sorting products by price
type ByPrice []Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface for sorting products by name
type ByName []Product

func (p ByName) Len() int           { return len(p) }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Binary search for integers
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // not found
}

// Binary search for strings
func binarySearchString(arr []string, target string) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // not found
}

// Generic binary search using sort.Search
func findInSorted(arr []int, target int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}

// Linear search for comparison
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// Bubble sort (educational - don't use in production)
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Selection sort (educational - don't use in production)
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Insertion sort (educational - don't use in production)
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Quick sort implementation
func quickSort(arr []int) {
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

	quickSort(arr[:right+1])
	quickSort(arr[left:])
}

// Merge sort implementation
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Custom sort with sort.Slice
func sortByMultipleCriteria(products []Product) {
	// Sort by price ascending, then by name ascending for same price
	sort.Slice(products, func(i, j int) bool {
		if products[i].Price != products[j].Price {
			return products[i].Price < products[j].Price
		}
		return products[i].Name < products[j].Name
	})
}

// Find minimum and maximum in unsorted array
func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	min, max := arr[0], arr[0]
	for _, v := range arr[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// Check if array is sorted
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Built-in Sort Functions ===")

	// Sort integers
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("Original: %v\n", ints)
	sort.Ints(ints)
	fmt.Printf("Sorted: %v\n", ints)

	// Sort strings
	strings := []string{"zebra", "apple", "banana", "cherry"}
	fmt.Printf("Original strings: %v\n", strings)
	sort.Strings(strings)
	fmt.Printf("Sorted strings: %v\n", strings)

	// Sort in reverse
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Printf("Reverse sorted: %v\n", ints)

	fmt.Println("\n=== Custom Sorting with sort.Interface ===")

	// Custom sorting with sort.Interface
	people := People{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
		{"David", 35},
	}

	fmt.Println("Before sorting:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	sort.Sort(people)

	fmt.Println("After sorting by age:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	fmt.Println("\n=== Sorting Products by Different Criteria ===")

	products := []Product{
		{"Laptop", 999.99, 5},
		{"Mouse", 25.50, 20},
		{"Keyboard", 75.00, 15},
		{"Monitor", 299.99, 8},
		{"Mouse", 30.00, 10}, // Same name as another mouse
	}

	fmt.Println("Original products:")
	for _, p := range products {
		fmt.Printf("  %s: $%.2f (%d in stock)\n", p.Name, p.Price, p.Stock)
	}

	// Sort by price
	sort.Sort(ByPrice(products))
	fmt.Println("\nSorted by price:")
	for _, p := range products {
		fmt.Printf("  %s: $%.2f (%d in stock)\n", p.Name, p.Price, p.Stock)
	}

	// Sort by name
	sort.Sort(ByName(products))
	fmt.Println("\nSorted by name:")
	for _, p := range products {
		fmt.Printf("  %s: $%.2f (%d in stock)\n", p.Name, p.Price, p.Stock)
	}

	fmt.Println("\n=== Binary Search ===")

	// Binary search requires sorted array
	sortedInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	// Manual binary search
	index := binarySearch(sortedInts, target)
	if index != -1 {
		fmt.Printf("Found %d at index %d (manual binary search)\n", target, index)
	} else {
		fmt.Printf("%d not found\n", target)
	}

	// Using sort.Search
	searchResult := findInSorted(sortedInts, target)
	if searchResult < len(sortedInts) && sortedInts[searchResult] == target {
		fmt.Printf("Found %d at index %d (sort.Search)\n", target, searchResult)
	}

	// Binary search for strings
	sortedStrings := []string{"apple", "banana", "cherry", "date", "elderberry"}
	targetStr := "cherry"

	strIndex := binarySearchString(sortedStrings, targetStr)
	if strIndex != -1 {
		fmt.Printf("Found '%s' at index %d\n", targetStr, strIndex)
	}

	fmt.Println("\n=== Search Performance Comparison ===")

	// Compare linear vs binary search
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = i
	}

	target = 5000

	// Linear search
	linearIndex := linearSearch(largeArray, target)
	fmt.Printf("Linear search found %d at index %d\n", target, linearIndex)

	// Binary search
	binaryIndex := binarySearch(largeArray, target)
	fmt.Printf("Binary search found %d at index %d\n", target, binaryIndex)

	fmt.Println("\n=== Educational Sort Algorithms ===")

	// Demonstrate different sorting algorithms
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Printf("Original array: %v\n", unsorted)

	// Bubble sort
	bubbleArr := make([]int, len(unsorted))
	copy(bubbleArr, unsorted)
	bubbleSort(bubbleArr)
	fmt.Printf("Bubble sort: %v\n", bubbleArr)

	// Selection sort
	selectionArr := make([]int, len(unsorted))
	copy(selectionArr, unsorted)
	selectionSort(selectionArr)
	fmt.Printf("Selection sort: %v\n", selectionArr)

	// Insertion sort
	insertionArr := make([]int, len(unsorted))
	copy(insertionArr, unsorted)
	insertionSort(insertionArr)
	fmt.Printf("Insertion sort: %v\n", insertionArr)

	// Quick sort
	quickArr := make([]int, len(unsorted))
	copy(quickArr, unsorted)
	quickSort(quickArr)
	fmt.Printf("Quick sort: %v\n", quickArr)

	// Merge sort
	mergeArr := mergeSort(unsorted)
	fmt.Printf("Merge sort: %v\n", mergeArr)

	fmt.Println("\n=== Advanced Sorting with sort.Slice ===")

	products2 := []Product{
		{"Widget A", 10.99, 5},
		{"Widget B", 15.50, 3},
		{"Widget C", 10.99, 8}, // Same price as Widget A
		{"Widget D", 20.00, 2},
	}

	fmt.Println("Before custom sort:")
	for _, p := range products2 {
		fmt.Printf("  %s: $%.2f\n", p.Name, p.Price)
	}

	sortByMultipleCriteria(products2)

	fmt.Println("After custom sort (price then name):")
	for _, p := range products2 {
		fmt.Printf("  %s: $%.2f\n", p.Name, p.Price)
	}

	fmt.Println("\n=== Utility Functions ===")

	// Find min/max
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := findMinMax(arr)
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// Check if sorted
	sortedArr := []int{1, 2, 3, 4, 5}
	unsortedArr := []int{3, 1, 4, 1, 5}

	fmt.Printf("Is %v sorted? %t\n", sortedArr, isSorted(sortedArr))
	fmt.Printf("Is %v sorted? %t\n", unsortedArr, isSorted(unsortedArr))
}
