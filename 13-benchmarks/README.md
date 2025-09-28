# 13. Benchmarks & Performance Testing

This example demonstrates Go's benchmarking capabilities and performance comparison techniques. Understanding algorithmic complexity and measuring performance is crucial for writing efficient code.

## Key Concepts

### Benchmark Functions
- Functions named `BenchmarkXxx` with `*testing.B` parameter
- Use `b.N` to control iteration count
- `b.ResetTimer()` to exclude setup time
- Compare different implementations

### Performance Analysis
- Big O notation understanding
- Time complexity vs space complexity
- Algorithm selection based on input size
- Profiling and optimization techniques

### Benchmark Best Practices
- Reset timer after setup
- Use realistic data sizes
- Compare algorithms fairly
- Use sub-benchmarks for different inputs

## Algorithm Comparisons

### Fibonacci Calculation
- **Recursive**: O(2^n) - Exponential time, very slow
- **Iterative**: O(n) - Linear time, efficient
- **Memoized**: O(n) - Linear time with O(n) space

### String Concatenation
- **+ operator**: O(n²) - Quadratic time due to reallocations
- **strings.Builder**: O(n) - Linear time, efficient

### Sorting Algorithms
- **Bubble Sort**: O(n²) - Quadratic time, simple but slow
- **Quick Sort**: O(n log n) - Log-linear time, efficient

### Search Algorithms
- **Linear Search**: O(n) - Linear time
- **Binary Search**: O(log n) - Logarithmic time (requires sorted data)

### Prime Generation
- **Sieve of Eratosthenes**: O(n log log n) - Very efficient
- **Trial Division**: O(n²) - Quadratic time, slower

## Benchmark Examples

### Basic Benchmark
```go
func BenchmarkFibonacciIterative(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FibonacciIterative(20)
    }
}
```

### Benchmark with Setup
```go
func BenchmarkBubbleSort(b *testing.B) {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    b.ResetTimer() // Exclude setup time
    for i := 0; i < b.N; i++ {
        copyArr := make([]int, len(arr))
        copy(copyArr, arr)
        BubbleSort(copyArr)
    }
}
```

### Sub-benchmarks
```go
func BenchmarkSorting(b *testing.B) {
    sizes := []int{10, 100, 1000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("QuickSort_%d", size), func(b *testing.B) {
            // Benchmark for specific size
        })
    }
}
```

## Running Benchmarks

### Run all benchmarks
```bash
go test -bench=.
```

### Run specific benchmark
```bash
go test -bench=BenchmarkFibonacci
```

### Run with memory allocation info
```bash
go test -bench=. -benchmem
```

### Run with CPU profiling
```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Compare benchmarks
```bash
go test -bench=. -count=5 | tee benchmark.txt
benchstat benchmark.txt
```

## Performance Results Interpretation

### Benchmark Output Format
```
BenchmarkFibonacciIterative-8    1000000    1234 ns/op
```
- **BenchmarkFibonacciIterative-8**: Function name and CPU cores
- **1000000**: Number of iterations
- **1234 ns/op**: Nanoseconds per operation

### Memory Allocation Output
```
BenchmarkStringConcatenation-8    100000    12345 ns/op    512 B/op    10 allocs/op
```
- **512 B/op**: Bytes allocated per operation
- **10 allocs/op**: Number of allocations per operation

## Common Performance Patterns

1. **Algorithm Selection**: Choose O(n log n) over O(n²) for large inputs
2. **Space-Time Tradeoffs**: Memoization uses space to save time
3. **Cache Efficiency**: Sequential access is faster than random access
4. **Memory Allocation**: Minimize allocations in hot paths
5. **String Operations**: Use strings.Builder for concatenation

## Profiling Tools

### CPU Profiling
```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Memory Profiling
```bash
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

### Trace Profiling
```bash
go test -bench=. -trace=trace.out
go tool trace trace.out
```

## Optimization Techniques

1. **Algorithm Improvement**: Better algorithms yield biggest gains
2. **Data Structure Selection**: Choose appropriate data structures
3. **Memory Management**: Reduce allocations and copying
4. **Concurrency**: Use goroutines for CPU-bound tasks
5. **Caching**: Cache expensive computations
6. **Lazy Evaluation**: Compute values only when needed

## Running the Example

```bash
cd 13-benchmarks

# Run the demo
go run main.go

# Run all benchmarks
go test -bench=.

# Run benchmarks with memory info
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkFibonacci

# Compare sorting algorithms
go test -bench=BenchmarkSorting
```

This example provides practical experience with performance testing and optimization, essential skills for professional Go development.