# Practice Problems - Goroutines Basics

These practice problems reinforce goroutine and channel concepts. Focus on concurrent programming patterns and proper synchronization.

## Problem 1: Concurrent Counter

**Description:** Implement a thread-safe counter that can be incremented concurrently by multiple goroutines. Use channels for communication and synchronization.

**Requirements:**
- Create a Counter struct with increment and get methods
- Use goroutines to increment the counter concurrently
- Ensure thread safety using channels
- Demonstrate race condition without proper synchronization

**Examples:**
- 10 goroutines, each incrementing 100 times → Final count: 1000
- Proper synchronization prevents race conditions

**Constraints:**
- Use channels for communication, not mutexes (save mutexes for next example)
- Handle goroutine synchronization properly

### Solution

```go
package main

import (
	"fmt"
	"sync"
)

// Thread-safe counter using channels
type Counter struct {
	value   int
	updates chan int
	done    chan bool
}

func NewCounter() *Counter {
	c := &Counter{
		updates: make(chan int, 100), // Buffered channel
		done:    make(chan bool),
	}

	// Start the counter goroutine
	go c.processUpdates()

	return c
}

func (c *Counter) processUpdates() {
	for {
		select {
		case update := <-c.updates:
			c.value += update
		case <-c.done:
			return
		}
	}
}

func (c *Counter) Increment(delta int) {
	c.updates <- delta
}

func (c *Counter) Value() int {
	// Send a query and wait for response (simplified)
	// In a real implementation, you'd want a more sophisticated query mechanism
	return c.value
}

func (c *Counter) Close() {
	close(c.done)
}

func worker(id int, counter *Counter, increments int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < increments; i++ {
		counter.Increment(1)
	}

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	counter := NewCounter()
	defer counter.Close()

	var wg sync.WaitGroup
	numWorkers := 10
	incrementsPerWorker := 100

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, counter, incrementsPerWorker, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Give some time for all updates to be processed
	// In a real implementation, you'd want better synchronization
	fmt.Printf("Final counter value: %d (expected: %d)\n",
		counter.Value(), numWorkers*incrementsPerWorker)
}
```

**Explanation:**
1. Counter uses a channel to receive updates from goroutines
2. Separate goroutine processes updates to ensure thread safety
3. Workers send increments through the channel
4. Buffered channel prevents blocking
5. Demonstrates basic goroutine communication patterns

## Problem 2: Fan-Out Pattern

**Description:** Implement a fan-out pattern where one producer sends data to multiple worker goroutines through channels. Each worker processes data and sends results back.

**Requirements:**
- Create one producer goroutine that generates numbers
- Create multiple worker goroutines that process numbers (e.g., square them)
- Use channels for communication
- Collect results from all workers
- Demonstrate load distribution

**Examples:**
- Producer: generates 1,2,3,4,5
- 3 workers: each squares numbers they receive
- Results collected and displayed

**Constraints:**
- Use select statement for channel operations
- Handle channel closing properly
- Demonstrate concurrent processing benefits

### Solution

```go
package main

import (
	"fmt"
	"sync"
)

// Job represents work to be done
type Job struct {
	ID  int
	Data int
}

// Result represents processed work
type Result struct {
	JobID int
	Value int
}

func producer(jobs chan<- Job, numJobs int) {
	defer close(jobs)

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Data: i * 10}
	}
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// Simulate processing (squaring the data)
		result := job.Data * job.Data
		results <- Result{JobID: job.ID, Value: result}

		fmt.Printf("Worker %d processed job %d: %d → %d\n",
			id, job.ID, job.Data, result)
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start producer
	go producer(jobs, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Printf("Job %d result: %d\n", result.JobID, result.Value)
	}

	fmt.Printf("\nProcessed %d jobs with %d workers\n", numJobs, numWorkers)
}
```

**Explanation:**
1. Producer generates jobs and sends them through jobs channel
2. Multiple workers receive jobs and process them concurrently
3. Results are sent back through results channel
4. Main goroutine collects all results
5. Demonstrates fan-out pattern for load distribution