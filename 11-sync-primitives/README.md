# 11. Synchronization Primitives

This example demonstrates Go's core synchronization primitives: `Mutex`, `RWMutex`, `WaitGroup`, and `Once`. These are essential for writing concurrent programs that are safe and correct.

## Key Concepts

### Mutex (Mutual Exclusion)
- Provides exclusive access to shared resources
- Only one goroutine can hold the lock at a time
- Use `Lock()` to acquire and `Unlock()` to release
- Critical for preventing race conditions

### RWMutex (Read-Write Mutex)
- Allows multiple readers or single writer
- Readers use `RLock()`/`RUnlock()`
- Writers use `Lock()`/`Unlock()`
- Ideal for read-heavy workloads

### WaitGroup
- Waits for a collection of goroutines to finish
- Use `Add()` to set expected count
- Call `Done()` when each goroutine completes
- `Wait()` blocks until all are done

### sync.Once
- Ensures a function executes only once
- Thread-safe initialization
- Perfect for singleton patterns and expensive setup

## Code Examples

### Thread-Safe Counter with Mutex
```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

### Safe Map with RWMutex
```go
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    value, exists := sm.data[key]
    return value, exists
}
```

### Coordinating Goroutines with WaitGroup
```go
var wg sync.WaitGroup
for i := 1; i <= 3; i++ {
    wg.Add(1)
    go worker(i, &wg)
}
wg.Wait() // Blocks until all workers done
```

### Singleton with sync.Once
```go
var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{value: expensiveOperation()}
    })
    return instance
}
```

## Common Patterns

1. **Resource Protection**: Use Mutex for any shared mutable state
2. **Read-Heavy Data**: Use RWMutex when reads >> writes
3. **Goroutine Coordination**: Use WaitGroup to wait for multiple goroutines
4. **Lazy Initialization**: Use sync.Once for expensive setup that should happen once

## Best Practices

- Always use `defer` to unlock mutexes
- Prefer RWMutex for read-heavy concurrent access
- Use WaitGroup for clean goroutine lifecycle management
- sync.Once is perfect for initialization that must happen exactly once
- Avoid deadlocks by acquiring locks in consistent order

## Running the Example

```bash
cd 11-sync-primitives
go run main.go
```

This will demonstrate all four synchronization primitives in action with concurrent operations.