package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generator pattern: Convert values to channel
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

// Square function for pipeline
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			time.Sleep(100 * time.Millisecond) // Simulate work
			out <- n * n
		}
	}()
	return out
}

// Fan-out pattern: Distribute work to multiple workers
func fanOut(in <-chan int, numWorkers int) []<-chan int {
	outs := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		outs[i] = worker(i, in)
	}
	return outs
}

func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for job := range in {
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			out <- job * 2 // Double the value
		}
	}()
	return out
}

// Fan-in pattern: Merge multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Start goroutine for each input channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for value := range c {
				out <- value
			}
		}(ch)
	}

	// Close output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Pipeline pattern: Chain of processing stages
func pipelineDemo() {
	// Create pipeline: generator -> square -> print
	numbers := generator(1, 2, 3, 4, 5)
	squared := square(numbers)

	fmt.Println("Pipeline results:")
	for result := range squared {
		fmt.Printf("Squared: %d\n", result)
	}
}

// Fan-out/Fan-in combined pattern
func fanOutFanInDemo() {
	// Generate work
	jobs := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Fan-out: distribute to 3 workers
	workerOutputs := fanOut(jobs, 3)

	// Fan-in: merge results
	results := fanIn(workerOutputs...)

	// Collect results
	fmt.Println("Fan-out/Fan-in results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Bounded parallelism pattern
func boundedParallelism(jobs []int, maxWorkers int) []int {
	type jobResult struct {
		index int
		value int
	}

	jobCh := make(chan jobResult, len(jobs))
	semaphore := make(chan struct{}, maxWorkers) // Limit concurrent workers

	// Start workers
	var wg sync.WaitGroup
	for i, job := range jobs {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release semaphore

			// Process job
			time.Sleep(200 * time.Millisecond) // Simulate work
			result := value * value

			jobCh <- jobResult{index: index, value: result}
		}(i, job)
	}

	// Close channel when all workers done
	go func() {
		wg.Wait()
		close(jobCh)
	}()

	// Collect results in order
	results := make([]int, len(jobs))
	for result := range jobCh {
		results[result.index] = result.value
	}

	return results
}

// Context-based cancellation pattern
func cancellableWorker(id int, jobs <-chan int, results chan<- int, done <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return // Jobs channel closed
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(300 * time.Millisecond)
			results <- job * 3
		case <-done:
			fmt.Printf("Worker %d cancelled\n", id)
			return
		}
	}
}

func cancellationDemo() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	done := make(chan struct{})

	// Start workers
	for i := 1; i <= 2; i++ {
		go cancellableWorker(i, jobs, results, done)
	}

	// Send some jobs
	go func() {
		for i := 1; i <= 3; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Cancel after 500ms
	time.Sleep(500 * time.Millisecond)
	close(done)

	// Collect any results that completed
	fmt.Println("Cancellation demo results:")
	timeout := time.After(1 * time.Second)
	for {
		select {
		case result := <-results:
			fmt.Printf("Got result: %d\n", result)
		case <-timeout:
			fmt.Println("Timeout waiting for results")
			return
		}
	}
}

// Error handling in concurrent code
func workerWithError(id int, jobs <-chan int, results chan<- int, errors chan<- error) {
	for job := range jobs {
		if job < 0 {
			errors <- fmt.Errorf("worker %d: negative job %d", id, job)
			continue
		}
		time.Sleep(100 * time.Millisecond)
		results <- job * job
	}
}

func errorHandlingDemo() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	errors := make(chan error, 5)

	// Start workers
	for i := 1; i <= 2; i++ {
		go workerWithError(i, jobs, results, errors)
	}

	// Send jobs (including one that will error)
	jobs <- 1
	jobs <- 2
	jobs <- -1 // This will cause an error
	jobs <- 4
	close(jobs)

	// Collect results and errors
	fmt.Println("Error handling demo:")
	for i := 0; i < 4; i++ {
		select {
		case result := <-results:
			fmt.Printf("Result: %d\n", result)
		case err := <-errors:
			fmt.Printf("Error: %v\n", err)
		}
	}
}

// Rate limiting pattern
func rateLimiter(requests <-chan int, rate time.Duration) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		ticker := time.NewTicker(rate)
		defer ticker.Stop()

		for req := range requests {
			<-ticker.C // Wait for next tick
			out <- req
		}
	}()
	return out
}

func rateLimitingDemo() {
	requests := make(chan int, 5)

	// Send requests
	go func() {
		for i := 1; i <= 5; i++ {
			requests <- i
		}
		close(requests)
	}()

	// Rate limit to 1 per 200ms
	limited := rateLimiter(requests, 200*time.Millisecond)

	fmt.Println("Rate limiting demo (1 request per 200ms):")
	start := time.Now()
	for req := range limited {
		fmt.Printf("Processed request %d at %v\n", req, time.Since(start).Round(100*time.Millisecond))
	}
}

// Future/Promise pattern (simplified)
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

func asyncTask(value int) *Future {
	future := &Future{
		result: make(chan int, 1),
		error:  make(chan error, 1),
	}

	go func() {
		time.Sleep(500 * time.Millisecond) // Simulate async work
		if value < 0 {
			future.error <- fmt.Errorf("negative value: %d", value)
		} else {
			future.result <- value * value
		}
	}()

	return future
}

func futureDemo() {
	fmt.Println("Future/Promise demo:")

	future1 := asyncTask(5)
	future2 := asyncTask(-1)

	// Get results
	if result, err := future1.Get(); err != nil {
		fmt.Printf("Future1 error: %v\n", err)
	} else {
		fmt.Printf("Future1 result: %d\n", result)
	}

	if result, err := future2.Get(); err != nil {
		fmt.Printf("Future2 error: %v\n", err)
	} else {
		fmt.Printf("Future2 result: %d\n", result)
	}
}

func main() {
	fmt.Println("=== Pipeline Pattern ===")
	pipelineDemo()

	fmt.Println("\n=== Fan-out/Fan-in Pattern ===")
	fanOutFanInDemo()

	fmt.Println("\n=== Bounded Parallelism ===")
	jobs := []int{1, 2, 3, 4, 5}
	results := boundedParallelism(jobs, 2)
	fmt.Printf("Bounded parallelism results: %v\n", results)

	fmt.Println("\n=== Cancellation Pattern ===")
	cancellationDemo()

	fmt.Println("\n=== Error Handling in Concurrent Code ===")
	errorHandlingDemo()

	fmt.Println("\n=== Rate Limiting ===")
	rateLimitingDemo()

	fmt.Println("\n=== Future/Promise Pattern ===")
	futureDemo()

	fmt.Println("\n=== Concurrency Patterns Summary ===")
	fmt.Println("✓ Generator: Convert values to channel")
	fmt.Println("✓ Pipeline: Chain processing stages")
	fmt.Println("✓ Fan-out: Distribute work to workers")
	fmt.Println("✓ Fan-in: Merge multiple channels")
	fmt.Println("✓ Bounded parallelism: Limit concurrent operations")
	fmt.Println("✓ Cancellation: Graceful shutdown")
	fmt.Println("✓ Error handling: Propagate errors in concurrent code")
	fmt.Println("✓ Rate limiting: Control operation frequency")
	fmt.Println("✓ Future/Promise: Handle async results")
}
