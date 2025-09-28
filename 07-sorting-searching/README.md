# 07 - Sorting and Searching

This example demonstrates Go's sorting and searching capabilities, including built-in functions, custom sorting, and algorithm implementations.

## Key Concepts

### Built-in Sort Functions

Go provides convenient sorting functions for basic types:

```go
sort.Ints([]int{3, 1, 4, 1, 5})        // Sort integers
sort.Strings([]string{"c", "a", "b"})   // Sort strings
sort.Float64s([]float64{3.1, 1.2, 4.3}) // Sort floats
```

### Custom Sorting with sort.Interface

Implement custom sorting by defining types that implement `sort.Interface`:

```go
type People []Person

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

people := People{{"Alice", 25}, {"Bob", 20}}
sort.Sort(people) // Sorts by age
```

### sort.Slice for Complex Sorting

Use `sort.Slice` for inline custom sorting logic:

```go
sort.Slice(products, func(i, j int) bool {
    // Custom comparison logic
    if products[i].Price != products[j].Price {
        return products[i].Price < products[j].Price
    }
    return products[i].Name < products[j].Name
})
```

### Binary Search

#### Manual Implementation
```go
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
```

#### Using sort.Search
```go
index := sort.Search(len(arr), func(i int) bool {
    return arr[i] >= target
})
```

### Sorting Algorithms

#### Bubble Sort (O(n²))
```go
func bubbleSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
```

#### Quick Sort (O(n log n) average)
```go
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
```

#### Merge Sort (O(n log n))
```go
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    return merge(left, right)
}
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
=== Built-in Sort Functions ===
Original: [3 1 4 1 5 9 2 6]
Sorted: [1 1 2 3 4 5 6 9]
Original strings: [zebra apple banana cherry]
Sorted strings: [apple banana cherry zebra]
Reverse sorted: [9 6 5 4 3 2 1 1]

=== Custom Sorting with sort.Interface ===
Before sorting:
  Alice: 25
  Bob: 30
  Charlie: 20
  David: 35
After sorting by age:
  Charlie: 20
  Alice: 25
  Bob: 30
  David: 35

=== Sorting Products by Different Criteria ===
Original products:
  Laptop: $999.99 (5 in stock)
  Mouse: $25.50 (20 in stock)
  Keyboard: $75.00 (15 in stock)
  Monitor: $299.99 (8 in stock)
  Mouse: $30.00 (10 in stock)

Sorted by price:
  Mouse: $25.50 (20 in stock)
  Mouse: $30.00 (10 in stock)
  Keyboard: $75.00 (15 in stock)
  Monitor: $299.99 (8 in stock)
  Laptop: $999.99 (5 in stock)

Sorted by name:
  Keyboard: $75.00 (15 in stock)
  Laptop: $999.99 (5 in stock)
  Monitor: $299.99 (8 in stock)
  Mouse: $25.50 (20 in stock)
  Mouse: $30.00 (10 in stock)

=== Binary Search ===
Found 7 at index 6 (manual binary search)
Found 7 at index 6 (sort.Search)
Found 'cherry' at index 2

=== Search Performance Comparison ===
Linear search found 5000 at index 5000
Binary search found 5000 at index 5000

=== Educational Sort Algorithms ===
Original array: [64 34 25 12 22 11 90]
Bubble sort: [11 12 22 25 34 64 90]
Selection sort: [11 12 22 25 34 64 90]
Insertion sort: [11 12 22 25 34 64 90]
Quick sort: [11 12 22 25 34 64 90]
Merge sort: [11 12 22 25 34 64 90]

=== Advanced Sorting with sort.Slice ===
Before custom sort:
  Widget A: $10.99
  Widget B: $15.50
  Widget C: $10.99
  Widget D: $20.00
After custom sort (price then name):
  Widget A: $10.99
  Widget C: $10.99
  Widget B: $15.50
  Widget D: $20.00

=== Utility Functions ===
Array: [3 1 4 1 5 9 2 6]
Min: 1, Max: 9
Is [1 2 3 4 5] sorted? true
Is [3 1 4 1 5] sorted? false
```

## Algorithm Complexities

| Algorithm | Time Complexity | Space Complexity | Stable |
|-----------|----------------|------------------|--------|
| Bubble Sort | O(n²) | O(1) | Yes |
| Selection Sort | O(n²) | O(1) | No |
| Insertion Sort | O(n²) | O(1) | Yes |
| Quick Sort | O(n log n) avg | O(log n) | No |
| Merge Sort | O(n log n) | O(n) | Yes |
| Built-in sort | O(n log n) | O(n) | Varies |

## Best Practices

### ✅ Do's
- **Use built-in sort functions** for basic types
- **Implement sort.Interface** for custom types
- **Use sort.Slice** for complex sorting logic
- **Use binary search** on sorted arrays
- **Consider stability** when choosing algorithms

### ❌ Don'ts
- **Don't implement sorting algorithms** unless for educational purposes
- **Don't use bubble sort** in production code
- **Don't forget to sort** before binary search
- **Don't modify arrays** during sorting unless you understand the implications

### When to Use Each Approach

1. **Built-in functions**: `sort.Ints()`, `sort.Strings()` - Simple cases
2. **sort.Interface**: Custom types with single sort criteria
3. **sort.Slice**: Complex sorting logic, multiple criteria
4. **Manual algorithms**: Educational purposes, specific constraints

## Common Interview Questions

- Implement binary search
- Sort an array of custom objects
- Find the kth largest element
- Sort strings by length
- Implement merge sort

## Performance Tips

- **Pre-sort data** if you'll search it multiple times
- **Use sort.Search** for cleaner binary search code
- **Consider sort.Slice** for complex comparisons
- **Profile your sorting** if performance is critical
- **Use stable sorts** when relative order matters

## Next Steps

- Implement more sorting algorithms (heap sort, radix sort)
- Practice solving LeetCode sorting problems in Go
- Compare performance of different approaches
- Implement custom comparers for complex data types