# Practice Problems - Concurrency Patterns

These practice problems reinforce advanced concurrency patterns. Focus on worker pools, pipelines, and proper resource management.

## Problem 1: Worker Pool with Result Collection

**Description:** Implement a worker pool that processes tasks concurrently and collects results. Each worker should process multiple tasks and return results through channels.

**Requirements:**
- Create a pool of worker goroutines
- Distribute tasks among workers
- Collect and aggregate results
- Handle graceful shutdown
- Demonstrate load balancing

**Examples:**
- 5 workers processing 20 tasks
- Each task: square a number
- Results: sum of all squared values

**Constraints:**
- Use buffered channels for efficiency
- Handle worker shutdown properly
- No race conditions in result collection

### Solution

```go
package main

import (
	"fmt"
	"sync"
)

// Task represents work to be done
type Task struct {
	ID   int
	Data int
}

// Result represents completed work
type Result struct {
	TaskID int
	Value  int
	Error  error
}

func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		// Simulate processing: square the number
		result := task.Data * task.Data

		// Simulate occasional errors
		var err error
		if task.Data < 0 {
			err = fmt.Errorf("negative number not allowed")
		}

		results <- Result{
			TaskID: task.ID,
			Value:  result,
			Error:  err,
		}

		fmt.Printf("Worker %d processed task %d: %d → %d\n",
			id, task.ID, task.Data, result)
	}
}

func main() {
	numWorkers := 3
	numTasks := 10

	// Create channels
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Start workers
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks
	go func() {
		defer close(tasks)
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{ID: i, Data: i}
		}
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	total := 0
	validResults := 0

	fmt.Println("\nResults:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Task %d failed: %v\n", result.TaskID, result.Error)
		} else {
			fmt.Printf("Task %d result: %d\n", result.TaskID, result.Value)
			total += result.Value
			validResults++
		}
	}

	fmt.Printf("\nSummary: %d valid results, total sum: %d\n", validResults, total)
}
```

**Explanation:**
1. Worker pool with fixed number of goroutines
2. Tasks distributed through buffered channel
3. Results collected through another channel
4. Graceful shutdown using sync.WaitGroup
5. Error handling in task processing

## Problem 2: Pipeline Pattern

**Description:** Implement a three-stage pipeline: generator → processor → aggregator. Each stage runs concurrently and communicates through channels.

**Requirements:**
- Stage 1: Generate numbers
- Stage 2: Filter and transform numbers
- Stage 3: Aggregate results
- Use channel-based communication
- Demonstrate pipeline composition

**Examples:**
- Generator: produces 1,2,3,4,5,6,7,8,9,10
- Processor: filters even numbers, squares them
- Aggregator: sums the results

**Constraints:**
- Each stage is a separate goroutine
- Proper channel closing to signal completion
- No blocking between stages

### Solution

```go
package main

import "fmt"

// Pipeline stage 1: Generator
func generator(nums []int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Pipeline stage 2: Processor (filter even numbers and square them)
func processor(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 { // Filter even numbers
				out <- n * n // Square them
			}
		}
	}()
	return out
}

// Pipeline stage 3: Aggregator (sum all values)
func aggregator(in <-chan int) <-chan int {
	out := make(chan int, 1) // Buffered to avoid blocking
	go func() {
		defer close(out)
		sum := 0
		count := 0
		for n := range in {
			sum += n
			count++
			fmt.Printf("Aggregator received: %d (running sum: %d)\n", n, sum)
		}
		out <- sum
	}()
	return out
}

func main() {
	// Input data
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Pipeline processing:", numbers)

	// Create pipeline
	gen := generator(numbers)
	proc := processor(gen)
	agg := aggregator(proc)

	// Get final result
	finalResult := <-agg

	fmt.Printf("\nFinal result: %d\n", finalResult)

	// Expected: 2² + 4² + 6² + 8² + 10² = 4 + 16 + 36 + 64 + 100 = 220
}
```

**Explanation:**
1. Generator stage produces initial data
2. Processor stage filters and transforms data
3. Aggregator stage collects final results
4. Each stage communicates through channels
5. Pipeline composition allows easy extension
6. Proper channel closing signals completion