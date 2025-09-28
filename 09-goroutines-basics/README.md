# 09 - Goroutines Basics

This example demonstrates Go's concurrency primitives: goroutines, channels, and the select statement. Understanding these fundamentals is crucial for writing concurrent Go programs.

## Key Concepts

### Goroutines

#### Creating Goroutines
```go
go sayHello()  // Starts function in new goroutine
go func() {
    // Anonymous goroutine
}()
```
- Lightweight threads managed by Go runtime
- Start with `go` keyword
- Run concurrently with main goroutine

#### Goroutine Lifecycle
- Created with `go` keyword
- Run until function returns
- Automatically cleaned up by runtime
- Main goroutine exit terminates program

### Channels

#### Basic Channel Operations
```go
ch := make(chan int)  // Unbuffered channel

ch <- 42    // Send (blocks until received)
value := <-ch  // Receive (blocks until sent)
```
- Communication mechanism between goroutines
- Send: `channel <- value`
- Receive: `<-channel`

#### Buffered Channels
```go
ch := make(chan int, 3)  // Buffer capacity of 3
ch <- 1  // Non-blocking (buffer has space)
ch <- 2
ch <- 3
// ch <- 4  // Would block (buffer full)
```
- Can hold multiple values
- Send blocks only when buffer is full
- Receive blocks only when buffer is empty

#### Channel Directions
```go
func sendOnly(ch chan<- int) {  // Can only send
    ch <- 42
}

func receiveOnly(ch <-chan int) int {  // Can only receive
    return <-ch
}
```
- Type safety for channel usage
- `chan T`: bidirectional
- `chan<- T`: send-only
- `<-chan T`: receive-only

### Select Statement

#### Basic Select
```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
}
```
- Waits for first available channel operation
- Like switch statement for channels
- Can include timeout cases

#### Fan-In Pattern
```go
func fanIn(ch1, ch2 <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        for {
            select {
            case msg := <-ch1:
                out <- msg
            case msg := <-ch2:
                out <- msg
            }
        }
    }()
    return out
}
```
- Multiplex multiple input channels to one output channel
- Common concurrency pattern

### Producer-Consumer Pattern

```go
func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch) // Signal no more data
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Println("Consumed:", value)
    }
}
```
- Producer generates data, sends to channel
- Consumer receives and processes data
- Channel closing signals completion

### Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        result := process(job)
        results <- result
    }
}

// Usage
jobs := make(chan int, numJobs)
results := make(chan int, numJobs)

for w := 1; w <= numWorkers; w++ {
    go worker(w, jobs, results)
}
```
- Fixed number of worker goroutines
- Distribute work via jobs channel
- Collect results via results channel

### Non-blocking Operations

```go
select {
case ch <- value:
    fmt.Println("Sent successfully")
default:
    fmt.Println("Channel full, couldn't send")
}
```
- Use `default` case for non-blocking operations
- Avoid blocking when channel operations would block

### Goroutine Leaks

#### Bad: Leaking Goroutines
```go
go func() {
    for {
        // Infinite loop, never exits
        select {
        case value := <-ch:
            // Process value
        }
    }
}()
// Goroutine runs forever even after main exits
```

#### Good: Proper Cleanup
```go
go func() {
    for {
        select {
        case value, ok := <-ch:
            if !ok {
                return // Channel closed, exit
            }
            // Process value
        }
    }
}()

// Later...
close(ch) // Signal goroutine to exit
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
=== Basic Goroutines ===
Hello from goroutine!
Goroutine A: 1
Goroutine B: 1
Goroutine A: 2
Goroutine B: 2
Goroutine A: 3
Goroutine B: 3

=== Channels ===
Producing: 1
Consuming: 1
Producing: 2
Consuming: 2
Producing: 3
Consuming: 3
Producer-consumer finished

=== Buffered Channels ===
Sent 3 values to buffered channel
Received: 1
Received: 2
Received: 3

=== Channel Directions ===
Received via directional channels: 100

=== Select Statement ===
Received: Message from channel 1

=== Fan-In Pattern ===
Combined: Hello from ch1
Combined: Another from ch1
Combined: Hello from ch2

=== Timeout Example ===
Operation timed out

=== Non-blocking Operations ===
Sent value successfully
Received: 42
Channel is empty

=== Worker Pool Pattern ===
Worker 3 processing job 1
Worker 1 processing job 2
Worker 2 processing job 3
Worker 1 processing job 4
Worker 3 processing job 5
Result: 2
Result: 4
Result: 6
Result: 8
Result: 10

=== Goroutine Lifecycle ===
--- Bad: Goroutine leak ---
Received: 42
Goroutine is still running (leak!)

--- Good: Proper cleanup ---
Received: 42
Channel closed, goroutine exiting
Goroutine properly cleaned up

=== Concurrency Concepts Demonstrated ===
✓ Goroutines: Lightweight threads
✓ Channels: Communication between goroutines
✓ Select: Multiplexing channel operations
✓ Buffered channels: Asynchronous communication
✓ Channel directions: Type safety
✓ Worker pools: Divide work among goroutines
✓ Proper cleanup: Avoid goroutine leaks
```

## Best Practices

### ✅ Do's
- **Use channels** for goroutine communication
- **Close channels** when no more data will be sent
- **Use select** for multiplexing channel operations
- **Handle channel closure** with `value, ok := <-ch`
- **Use buffered channels** for performance when appropriate
- **Clean up goroutines** properly to avoid leaks

### ❌ Don'ts
- **Don't use shared memory** for communication (use channels)
- **Don't forget to close channels** (can cause deadlocks)
- **Don't send on closed channels** (runtime panic)
- **Don't use sleep** for synchronization (use channels)
- **Don't ignore goroutine leaks** (use proper cleanup)

### Channel Types and Use Cases

| Channel Type | Use Case | Blocking Behavior |
|--------------|----------|-------------------|
| `chan T` | General purpose | Send/Receive block |
| `chan T` (buffered) | Async communication | Block when buffer full/empty |
| `chan<- T` | Send-only parameter | Type safety |
| `<-chan T` | Receive-only parameter | Type safety |

### Common Patterns

1. **Generator Pattern**: Goroutine produces values, sends to channel
2. **Worker Pool**: Fixed workers process jobs from channel
3. **Fan-In**: Multiple inputs multiplexed to single output
4. **Fan-Out**: Single input distributed to multiple workers
5. **Pipeline**: Chain of goroutines processing data

## Interview Questions

- What are goroutines and how do they differ from threads?
- How do channels work in Go?
- What is the select statement used for?
- How do you avoid goroutine leaks?
- When should you use buffered vs unbuffered channels?
- How do you implement a worker pool in Go?

## Performance Considerations

- **Goroutines are cheap**: Thousands can run concurrently
- **Channels have overhead**: Use buffering when appropriate
- **Select is efficient**: Good for multiplexing
- **Avoid contention**: Design to minimize blocking

## Next Steps

- Study advanced concurrency patterns (fan-out, pipelines)
- Learn sync package primitives (Mutex, WaitGroup)
- Practice implementing concurrent data structures
- Understand race conditions and how to avoid them
- Explore context package for cancellation