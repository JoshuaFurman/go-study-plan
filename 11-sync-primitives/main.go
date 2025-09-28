package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter demonstrates Mutex usage for thread-safe operations
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// SafeMap demonstrates RWMutex for read-heavy operations
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

// Worker function for WaitGroup demonstration
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

// Expensive operation that should only run once
func expensiveOperation() int {
	fmt.Println("Performing expensive operation...")
	time.Sleep(2 * time.Second)
	return 42
}

// Singleton demonstrates sync.Once usage
type Singleton struct {
	value int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance")
		instance = &Singleton{value: expensiveOperation()}
	})
	return instance
}

func main() {
	fmt.Println("=== Sync Primitives Demo ===\n")

	// 1. Mutex demonstration
	fmt.Println("1. Mutex - Thread-safe counter:")
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final counter value: %d (expected: 500)\n\n", counter.Value())

	// 2. RWMutex demonstration
	fmt.Println("2. RWMutex - Safe map with concurrent reads:")
	safeMap := NewSafeMap()

	// Start multiple readers
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				if value, exists := safeMap.Get(fmt.Sprintf("key%d", j)); exists {
					fmt.Printf("Reader %d got key%d: %d\n", id, j, value)
				}
			}
		}(i)
	}

	// Start writers
	for i := 0; i < 10; i++ {
		go func(i int) {
			safeMap.Set(fmt.Sprintf("key%d", i), i*10)
			fmt.Printf("Writer set key%d: %d\n", i, i*10)
		}(i)
	}

	time.Sleep(1 * time.Second) // Wait for operations to complete
	fmt.Println()

	// 3. WaitGroup demonstration
	fmt.Println("3. WaitGroup - Coordinating goroutines:")
	var wg2 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go worker(i, &wg2)
	}
	wg2.Wait()
	fmt.Println("All workers completed\n")

	// 4. sync.Once demonstration
	fmt.Println("4. sync.Once - Singleton pattern:")
	for i := 0; i < 3; i++ {
		go func(id int) {
			singleton := GetInstance()
			fmt.Printf("Goroutine %d got singleton value: %d\n", id, singleton.value)
		}(i)
	}

	time.Sleep(3 * time.Second) // Wait for goroutines
	fmt.Println("\n=== Demo Complete ===")
}
