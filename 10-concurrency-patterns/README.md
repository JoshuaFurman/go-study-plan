# 10 - Concurrency Patterns

This example demonstrates advanced concurrency patterns in Go, including pipelines, fan-out/fan-in, bounded parallelism, and other common concurrent programming techniques.

## Key Concepts

### Generator Pattern

#### Converting Values to Channel
```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}
```
- Convert data sources to channels
- Separate data generation from processing
- Return read-only channel for safety

### Pipeline Pattern

#### Chain of Processing Stages
```go
numbers := generator(1, 2, 3, 4, 5)
squared := square(numbers)  // Each stage is a goroutine

for result := range squared {
    fmt.Println(result)
}
```
- Chain of goroutines processing data
- Each stage receives from previous, sends to next
- Natural fit for streaming data processing

### Fan-Out Pattern

#### Distribute Work to Multiple Workers
```go
func fanOut(in <-chan int, numWorkers int) []<-chan int {
    outs := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        outs[i] = worker(i, in)
    }
    return outs
}
```
- Distribute work from one channel to multiple workers
- Parallel processing of jobs
- Each worker gets its own output channel

### Fan-In Pattern

#### Merge Multiple Channels
```go
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for value := range c {
                out <- value
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```
- Combine multiple input channels into one output channel
- Multiplexing multiple data sources
- Wait for all inputs to complete

### Bounded Parallelism

#### Limit Concurrent Operations
```go
semaphore := make(chan struct{}, maxWorkers)

go func() {
    semaphore <- struct{}{}  // Acquire
    defer func() { <-semaphore }()  // Release

    // Do work
}()
```
- Control number of concurrent operations
- Prevent resource exhaustion
- Use buffered channel as semaphore

### Cancellation Pattern

#### Graceful Shutdown
```go
func worker(jobs <-chan int, done <-chan struct{}) {
    for {
        select {
        case job := <-jobs:
            // Process job
        case <-done:
            return // Exit gracefully
        }
    }
}
```
- Signal goroutines to stop
- Use `done` channel for cancellation
- Clean up resources properly

### Error Handling in Concurrent Code

#### Propagate Errors Through Channels
```go
func worker(jobs <-chan int, results chan<- int, errors chan<- error) {
    for job := range jobs {
        if err := process(job); err != nil {
            errors <- err
            continue
        }
        results <- result
    }
}
```
- Separate error and result channels
- Handle errors without stopping processing
- Aggregate errors for reporting

### Rate Limiting

#### Control Operation Frequency
```go
func rateLimiter(requests <-chan int, rate time.Duration) <-chan int {
    ticker := time.NewTicker(rate)
    defer ticker.Stop()

    return func() <-chan int {
        out := make(chan int)
        go func() {
            defer close(out)
            for req := range requests {
                <-ticker.C  // Wait for next tick
                out <- req
            }
        }()
        return out
    }()
}
```
- Limit frequency of operations
- Use ticker for regular intervals
- Smooth out bursty traffic

### Future/Promise Pattern

#### Handle Async Results
```go
type Future struct {
    result chan int
    error  chan error
}

func (f *Future) Get() (int, error) {
    select {
    case result := <-f.result:
        return result, nil
    case err := <-f.error:
        return 0, err
    }
}
```
- Represent asynchronous computation
- Block until result is available
- Handle both success and error cases

## Running the Example

```bash
go run main.go
```

## Expected Output

```
=== Pipeline Pattern ===
Squared: 1
Squared: 4
Squared: 9
Squared: 16
Squared: 25

=== Fan-out/Fan-in Pattern ===
Worker 0 processing job 1
Worker 1 processing job 2
Worker 2 processing job 3
Worker 0 processing job 4
Worker 1 processing job 5
Worker 2 processing job 6
Worker 0 processing job 7
Worker 1 processing job 8
Worker 2 processing job 9
Worker 0 processing job 10
Result: 2
Result: 4
Result: 6
Result: 8
Result: 10
Result: 12
Result: 14
Result: 16
Result: 18
Result: 20

=== Bounded Parallelism ===
Bounded parallelism results: [1 4 9 16 25]

=== Cancellation Pattern ===
Worker 1 processing job 1
Worker 2 processing job 2
Worker 1 processing job 3
Worker 2 cancelled
Worker 1 cancelled
Cancellation demo results:
Got result: 3
Got result: 6
Got result: 9
Timeout waiting for results

=== Error Handling in Concurrent Code ===
Result: 1
Result: 4
Error: worker 1: negative job -1
Result: 16

=== Rate Limiting ===
Rate limiting demo (1 request per 200ms):
Processed request 1 at 200ms
Processed request 2 at 400ms
Processed request 3 at 600ms
Processed request 4 at 800ms
Processed request 5 at 1s

=== Future/Promise Pattern ===
Future1 result: 25
Future2 error: negative value: -1

=== Concurrency Patterns Summary ===
✓ Generator: Convert values to channel
✓ Pipeline: Chain processing stages
✓ Fan-out: Distribute work to workers
✓ Fan-in: Merge multiple channels
✓ Bounded parallelism: Limit concurrent operations
✓ Cancellation: Graceful shutdown
✓ Error handling: Propagate errors in concurrent code
✓ Rate limiting: Control operation frequency
✓ Future/Promise: Handle async results
```

## Pattern Comparison

| Pattern | Use Case | Benefits | Considerations |
|---------|----------|----------|----------------|
| **Pipeline** | Streaming data processing | Composable, clean separation | Backpressure handling |
| **Fan-out/Fan-in** | Parallel processing | Load distribution | Result ordering |
| **Bounded Parallelism** | Resource control | Prevents overload | Throughput vs latency |
| **Cancellation** | Graceful shutdown | Clean resource cleanup | Coordination complexity |
| **Rate Limiting** | Traffic control | Prevents abuse | Queue management |
| **Future/Promise** | Async operations | Non-blocking calls | Error propagation |

## Best Practices

### ✅ Do's
- **Use channels for communication** between goroutines
- **Close channels** when no more data will be sent
- **Handle channel closure** with `value, ok := <-ch`
- **Use sync.WaitGroup** for coordinating multiple goroutines
- **Implement cancellation** for long-running operations
- **Limit concurrency** with semaphores when needed
- **Handle errors** appropriately in concurrent code

### ❌ Don'ts
- **Don't use shared memory** for communication (use channels)
- **Don't forget to close channels** (can cause deadlocks)
- **Don't send on closed channels** (runtime panic)
- **Don't ignore goroutine leaks** (use proper cleanup)
- **Don't use time.Sleep** for synchronization
- **Don't block indefinitely** without cancellation

### Performance Considerations

- **Goroutines are cheap** but not free
- **Channels have overhead** - use buffering appropriately
- **Select is efficient** for multiplexing
- **Bounded parallelism** prevents resource exhaustion
- **Profile concurrent code** to identify bottlenecks

## Interview Questions

- How do you implement a worker pool in Go?
- What's the difference between fan-out and fan-in?
- How do you gracefully shut down goroutines?
- How do you limit concurrent operations?
- How do you handle errors in concurrent code?
- When should you use pipelines vs other patterns?

## Next Steps

- Study the sync package primitives (Mutex, WaitGroup)
- Learn about context for cancellation
- Practice implementing these patterns
- Study real-world examples in popular Go libraries
- Understand deadlock prevention and detection