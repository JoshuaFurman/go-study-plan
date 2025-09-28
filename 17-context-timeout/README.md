# 17. Context & Timeout

This example demonstrates Go's context package for handling cancellation, timeouts, and request-scoped values. Context is essential for writing robust, cancellable concurrent programs and is a key part of Go's concurrency model.

## Key Concepts

### Context Types
- **context.Background()**: Root context, never cancelled
- **context.TODO()**: Placeholder for contexts not yet determined
- **context.WithCancel()**: Creates cancellable context
- **context.WithTimeout()**: Creates context that cancels after timeout
- **context.WithDeadline()**: Creates context that cancels at specific time
- **context.WithValue()**: Creates context with key-value pairs

### Context Interface
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### Cancellation Propagation
- Contexts form a tree hierarchy
- Cancelling parent context cancels all children
- Done channel signals cancellation
- Err() returns cancellation reason

## Context Creation Patterns

### Basic Cancellation
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Always call cancel to release resources

go func() {
    select {
    case <-time.After(time.Second):
        fmt.Println("Work done")
    case <-ctx.Done():
        fmt.Println("Work cancelled:", ctx.Err())
    }
}()

// Cancel after some time
time.Sleep(500 * time.Millisecond)
cancel()
```

### Timeout Context
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// This will timeout if work takes longer than 2 seconds
err := doWork(ctx)
if err != nil {
    if err == context.DeadlineExceeded {
        fmt.Println("Work timed out")
    }
}
```

### Context with Values
```go
// Add request-scoped values
ctx := context.WithValue(context.Background(), "userID", 12345)
ctx = context.WithValue(ctx, "requestID", "req-abc-123")

// Retrieve values in child functions
userID := ctx.Value("userID")
requestID := ctx.Value("requestID")
```

## Common Patterns

### HTTP Request with Timeout
```go
func httpRequest(ctx context.Context, url string) error {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }

    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return fmt.Errorf("request timed out")
        }
        return err
    }
    defer resp.Body.Close()
    return nil
}
```

### Worker Pool with Cancellation
```go
func workerPool(ctx context.Context, jobs <-chan Job) {
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    processJob(ctx, job)
                case <-ctx.Done():
                    fmt.Printf("Worker %d cancelled\n", workerID)
                    return
                }
            }
        }(i)
    }

    wg.Wait()
}
```

### Database Query with Timeout
```go
func queryDatabase(ctx context.Context, query string) (*sql.Rows, error) {
    // Set query timeout
    queryCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    return db.QueryContext(queryCtx, query)
}
```

## Context Hierarchy & Cancellation

### Parent-Child Relationship
```go
// Root context
rootCtx := context.Background()

// Service context with timeout
serviceCtx, serviceCancel := context.WithTimeout(rootCtx, 30*time.Second)

// Request context derived from service
requestCtx, requestCancel := context.WithCancel(serviceCtx)

// Database context with shorter timeout
dbCtx, dbCancel := context.WithTimeout(requestCtx, 5*time.Second)

// Cancelling serviceCtx will cancel all children
serviceCancel()
```

### Graceful Shutdown
```go
func serverShutdown(ctx context.Context, server *http.Server) error {
    // Create shutdown context with timeout
    shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
    defer shutdownCancel()

    // Graceful shutdown
    return server.Shutdown(shutdownCtx)
}
```

## Best Practices

### Always Cancel Contexts
```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel() // Prevents context leak
```

### Don't Store Contexts in Structs
```go
// Bad - context stored in struct
type Handler struct {
    ctx context.Context
}

// Good - pass context as parameter
func (h *Handler) Process(ctx context.Context) {
    // Use ctx for this operation
}
```

### Use Context for Cancellation Only
- Use context for cancellation signals
- Use context.WithValue sparingly
- Consider using structured logging instead of context values

### Context Timeout vs Client Timeout
```go
// Context timeout affects request
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

// Client timeout is separate
client := &http.Client{Timeout: 10 * time.Second}
```

## Common Mistakes

1. **Forgetting to cancel contexts** - causes memory leaks
2. **Storing contexts in structs** - leads to bugs
3. **Using context for configuration** - use parameters instead
4. **Not checking ctx.Done()** - goroutines won't be cancelled
5. **Blocking on context.Done()** - can cause deadlocks

## Testing with Context

### Testing Cancellation
```go
func TestWorkerCancellation(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())

    done := make(chan bool)
    go func() {
        worker(ctx)
        done <- true
    }()

    // Cancel and wait for worker to finish
    cancel()
    select {
    case <-done:
        // Worker cancelled successfully
    case <-time.After(time.Second):
        t.Error("Worker didn't cancel in time")
    }
}
```

### Testing Timeouts
```go
func TestTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()

    err := slowOperation(ctx)
    if err != context.DeadlineExceeded {
        t.Errorf("Expected timeout, got %v", err)
    }
}
```

## Performance Considerations

- **Context creation is cheap** - don't avoid creating contexts
- **Context cancellation is fast** - Done() channel is efficient
- **Context values have overhead** - use sparingly
- **Deep context trees** - can be expensive to traverse

## When to Use Context

1. **HTTP requests** - cancel long-running requests
2. **Database operations** - set query timeouts
3. **Worker pools** - coordinate goroutine lifecycles
4. **API calls** - prevent hanging requests
5. **File operations** - cancel I/O operations
6. **Long computations** - allow user cancellation

## Running the Example

```bash
cd 17-context-timeout
go run main.go
```

This example demonstrates context usage patterns essential for writing production-ready Go applications that handle cancellation, timeouts, and request-scoped values properly.