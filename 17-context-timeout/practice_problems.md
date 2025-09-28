# Practice Problems - Context and Timeout

These practice problems reinforce context usage for cancellation and timeouts. Focus on proper resource cleanup and graceful shutdown.

## Problem 1: HTTP Request with Timeout

**Description:** Make an HTTP GET request with a timeout using context. Handle cancellation and timeout scenarios properly.

**Requirements:**
- Use context.WithTimeout for request timeout
- Handle context cancellation
- Return appropriate errors for different failure modes
- Clean up resources properly

**Examples:**
- Successful request within timeout
- Request that times out
- Request cancelled by context

**Constraints:**
- Use net/http for requests
- Timeout should be configurable
- Handle all error cases

### Solution

```go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchWithTimeout(url string, timeout time.Duration) (string, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // Always call cancel to release resources

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Check if the error is due to context cancellation/timeout
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("request timed out after %v", timeout)
		}
		if ctx.Err() == context.Canceled {
			return "", fmt.Errorf("request was cancelled")
		}
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}

func simulateSlowServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Slow response"))
	})

	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Fast response"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		fmt.Println("Starting test server on :8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	return server
}

func main() {
	// Start a test server
	server := simulateSlowServer()
	defer server.Close()

	time.Sleep(100 * time.Millisecond) // Let server start

	// Test cases
	testCases := []struct {
		url     string
		timeout time.Duration
		desc    string
	}{
		{"http://localhost:8080/fast", 5 * time.Second, "Fast request with long timeout"},
		{"http://localhost:8080/slow", 5 * time.Second, "Slow request with long timeout"},
		{"http://localhost:8080/slow", 1 * time.Second, "Slow request with short timeout"},
	}

	for i, tc := range testCases {
		fmt.Printf("\nTest %d: %s\n", i+1, tc.desc)
		fmt.Printf("URL: %s, Timeout: %v\n", tc.url, tc.timeout)

		start := time.Now()
		result, err := fetchWithTimeout(tc.url, tc.timeout)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf("Error: %v (took %v)\n", err, duration)
		} else {
			fmt.Printf("Success: Got %d bytes (took %v)\n", len(result), duration)
		}
	}
}
```

**Explanation:**
1. Use context.WithTimeout to create cancellable context
2. Always defer cancel() to release resources
3. Check ctx.Err() to distinguish timeout from other errors
4. Handle HTTP errors and response reading
5. Demonstrate proper resource cleanup

## Problem 2: Worker Pool with Cancellation

**Description:** Create a worker pool that can be cancelled using context. Workers should stop processing when context is cancelled.

**Requirements:**
- Use context for cancellation
- Multiple workers processing tasks
- Graceful shutdown when cancelled
- Track completed vs cancelled work

**Examples:**
- Workers processing tasks normally
- Cancellation signal stops all workers
- Partial completion when cancelled

**Constraints:**
- Use context.WithCancel
- Workers should check context regularly
- Report work completed before cancellation

### Solution

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data int
}

type Result struct {
	TaskID    int
	Value     int
	Completed bool
}

func worker(ctx context.Context, id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				// Channel closed, no more tasks
				return
			}

			// Simulate work
			workTime := time.Duration(rand.Intn(500)+100) * time.Millisecond
			timer := time.NewTimer(workTime)
			defer timer.Stop()

			select {
			case <-timer.C:
				// Work completed
				result := task.Data * task.Data
				select {
				case results <- Result{TaskID: task.ID, Value: result, Completed: true}:
				case <-ctx.Done():
					return
				}
				fmt.Printf("Worker %d completed task %d: %d → %d\n", id, task.ID, task.Data, result)

			case <-ctx.Done():
				// Context cancelled during work
				fmt.Printf("Worker %d cancelled while working on task %d\n", id, task.ID)
				return
			}

		case <-ctx.Done():
			// Context cancelled while waiting for task
			fmt.Printf("Worker %d cancelled while waiting for task\n", id)
			return
		}
	}
}

func main() {
	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Channels for tasks and results
	tasks := make(chan Task, 20)
	results := make(chan Result, 20)

	// Start workers
	numWorkers := 3
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, tasks, results, &wg)
	}

	// Send tasks
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- Task{ID: i, Data: i * 10}
		}
		close(tasks)
	}()

	// Cancel after some time to demonstrate cancellation
	go func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("\nCancelling work...")
		cancel()
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	completed := 0
	cancelled := 0

	for result := range results {
		if result.Completed {
			completed++
			fmt.Printf("Result: Task %d = %d\n", result.TaskID, result.Value)
		}
	}

	fmt.Printf("\nSummary: %d tasks completed\n", completed)

	// Wait a bit for cancellation messages
	time.Sleep(100 * time.Millisecond)
}
```

**Explanation:**
1. Use context.WithCancel for cancellation control
2. Workers check context in multiple select statements
3. Graceful shutdown when context is cancelled
4. Track completed work vs cancelled work
5. Demonstrate proper goroutine cleanup