# Practice Problems - Sync Primitives

These practice problems reinforce synchronization primitives. Focus on mutexes, RWMutex, WaitGroup, and Once for proper concurrent programming.

## Problem 1: Thread-Safe Counter with RWMutex

**Description:** Implement a thread-safe counter that supports increment, decrement, and value retrieval operations. Use RWMutex to optimize for multiple readers.

**Requirements:**
- Use RWMutex for efficient read operations
- Support concurrent reads and exclusive writes
- Demonstrate performance difference between Mutex and RWMutex
- Handle negative values appropriately

**Examples:**
- Multiple goroutines reading counter value
- Single writer updating counter
- RWMutex allows multiple concurrent readers

**Constraints:**
- Use RWMutex.RLock() for reads, RWMutex.Lock() for writes
- Ensure proper unlocking with defer
- Demonstrate lock contention scenarios

### Solution

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Thread-safe counter using RWMutex
type Counter struct {
	value int
	mutex sync.RWMutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Decrement() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value--
}

func (c *Counter) Value() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

func (c *Counter) Add(delta int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value += delta
}

func reader(id int, counter *Counter, reads *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		_ = counter.Value() // Read operation
		*reads++
	}
}

func writer(id int, counter *Counter, writes *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 50; i++ {
		counter.Increment() // Write operation
		*writes++
	}
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	numReaders := 10
	numWriters := 2

	// Track operations
	var totalReads, totalWrites int

	fmt.Printf("Starting with %d readers and %d writers\n", numReaders, numWriters)

	start := time.Now()

	// Start readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go reader(i, counter, &totalReads, &wg)
	}

	// Start writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go writer(i, counter, &totalWrites, &wg)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Completed in %v\n", duration)
	fmt.Printf("Final counter value: %d\n", counter.Value())
	fmt.Printf("Total reads: %d, Total writes: %d\n", totalReads, totalWrites)
	fmt.Printf("Expected final value: %d\n", totalWrites)
}
```

**Explanation:**
1. RWMutex allows multiple concurrent readers
2. Writers get exclusive access
3. RLock() for read operations, Lock() for write operations
4. Demonstrates performance benefits of RWMutex for read-heavy workloads

## Problem 2: Singleton Pattern with sync.Once

**Description:** Implement a singleton pattern using sync.Once to ensure a resource is initialized exactly once, even in concurrent environments.

**Requirements:**
- Use sync.Once for thread-safe initialization
- Create a database connection pool singleton
- Demonstrate that initialization happens only once
- Handle concurrent access safely

**Examples:**
- Multiple goroutines trying to get database instance
- Only one initialization occurs
- All goroutines get the same instance

**Constraints:**
- sync.Once ensures function runs exactly once
- Initialization must be thread-safe
- Return the same instance to all callers

### Solution

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Database represents a database connection pool
type Database struct {
	name      string
	connections int
	initialized time.Time
}

var (
	instance *Database
	once     sync.Once
)

// GetDatabaseInstance returns the singleton database instance
func GetDatabaseInstance() *Database {
	once.Do(func() {
		fmt.Println("Initializing database connection pool...")
		// Simulate expensive initialization
		time.Sleep(100 * time.Millisecond)

		instance = &Database{
			name:        "MainDB",
			connections: 10,
			initialized: time.Now(),
		}

		fmt.Println("Database initialized successfully!")
	})

	return instance
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	db := GetDatabaseInstance()
	fmt.Printf("Worker %d got database instance: %s (initialized at %v)\n",
		id, db.name, db.initialized.Format("15:04:05.000"))
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	fmt.Println("Starting concurrent database access test...")

	// Start multiple workers trying to access database
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	// Verify all workers got the same instance
	db1 := GetDatabaseInstance()
	db2 := GetDatabaseInstance()

	fmt.Printf("\nInstance verification:\n")
	fmt.Printf("db1 == db2: %t\n", db1 == db2)
	fmt.Printf("Same initialization time: %t\n", db1.initialized.Equal(db2.initialized))
	fmt.Printf("Database name: %s, Connections: %d\n", db1.name, db1.connections)
}
```

**Explanation:**
1. sync.Once ensures initialization function runs exactly once
2. Multiple goroutines can safely call GetDatabaseInstance()
3. Only one initialization occurs, even with concurrent access
4. All callers receive the same instance
5. Demonstrates lazy initialization pattern