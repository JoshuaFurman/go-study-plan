package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// simulateWork simulates some work that takes time
func simulateWork(ctx context.Context, id int, duration time.Duration) error {
	fmt.Printf("Worker %d: Starting work for %v\n", id, duration)

	select {
	case <-time.After(duration):
		fmt.Printf("Worker %d: Work completed successfully\n", id)
		return nil
	case <-ctx.Done():
		fmt.Printf("Worker %d: Work cancelled due to: %v\n", id, ctx.Err())
		return ctx.Err()
	}
}

// simulateAPI simulates an API call with timeout
func simulateAPI(ctx context.Context, endpoint string) (string, error) {
	fmt.Printf("API: Calling %s\n", endpoint)

	// Simulate variable response time
	responseTime := time.Duration(rand.Intn(3000)+500) * time.Millisecond

	select {
	case <-time.After(responseTime):
		result := fmt.Sprintf("Response from %s", endpoint)
		fmt.Printf("API: Got response: %s\n", result)
		return result, nil
	case <-ctx.Done():
		fmt.Printf("API: Request to %s cancelled: %v\n", endpoint, ctx.Err())
		return "", ctx.Err()
	}
}

// workerWithContext demonstrates context usage in goroutines
func workerWithContext(ctx context.Context, id int) {
	fmt.Printf("Worker %d: Started\n", id)

	// Simulate some work
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Worker %d: Still working...\n", id)
		case <-ctx.Done():
			fmt.Printf("Worker %d: Received cancellation signal: %v\n", id, ctx.Err())
			return
		}
	}
}

// processWithTimeout demonstrates timeout handling
func processWithTimeout(ctx context.Context, name string, processTime time.Duration) error {
	fmt.Printf("Process %s: Starting (will take %v)\n", name, processTime)

	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-time.After(processTime):
		fmt.Printf("Process %s: Completed successfully\n", name)
		return nil
	case <-timeoutCtx.Done():
		fmt.Printf("Process %s: Timed out: %v\n", name, timeoutCtx.Err())
		return timeoutCtx.Err()
	}
}

// contextWithValue demonstrates passing values through context
func contextWithValue(ctx context.Context) {
	// Add user ID to context
	userCtx := context.WithValue(ctx, "userID", 12345)
	requestCtx := context.WithValue(userCtx, "requestID", "req-abc-123")

	// Simulate processing with context values
	processRequest(requestCtx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")

	fmt.Printf("Processing request %v for user %v\n", requestID, userID)

	// Simulate database query
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("Request %v: Database query completed\n", requestID)
	case <-ctx.Done():
		fmt.Printf("Request %v: Cancelled during database query\n", requestID)
		return
	}

	// Simulate API call
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("Request %v: External API call completed\n", requestID)
	case <-ctx.Done():
		fmt.Printf("Request %v: Cancelled during API call\n", requestID)
		return
	}

	fmt.Printf("Request %v: Processing completed\n", requestID)
}

// chainedContexts demonstrates context chaining and cancellation propagation
func chainedContexts() {
	fmt.Println("\n=== Chained Contexts Demo ===")

	// Root context
	rootCtx := context.Background()

	// Parent context with timeout
	parentCtx, parentCancel := context.WithTimeout(rootCtx, 3*time.Second)
	defer parentCancel()

	// Child context derived from parent
	childCtx, childCancel := context.WithCancel(parentCtx)
	defer childCancel()

	// Start workers
	go func() {
		if err := simulateWork(childCtx, 1, 1*time.Second); err != nil {
			fmt.Printf("Worker 1 error: %v\n", err)
		}
	}()

	go func() {
		if err := simulateWork(childCtx, 2, 4*time.Second); err != nil {
			fmt.Printf("Worker 2 error: %v\n", err)
		}
	}()

	// Let workers run for a bit
	time.Sleep(500 * time.Millisecond)

	// Cancel child context (should affect both workers)
	fmt.Println("Cancelling child context...")
	childCancel()

	// Wait for parent timeout
	<-parentCtx.Done()
	fmt.Printf("Parent context done: %v\n", parentCtx.Err())
}

// concurrentAPICalls demonstrates making multiple API calls with timeout
func concurrentAPICalls() {
	fmt.Println("\n=== Concurrent API Calls Demo ===")

	// Create context with timeout for all API calls
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	endpoints := []string{"api/user", "api/orders", "api/products", "api/analytics"}

	results := make(chan string, len(endpoints))
	errors := make(chan error, len(endpoints))

	// Start concurrent API calls
	for i, endpoint := range endpoints {
		go func(id int, ep string) {
			result, err := simulateAPI(ctx, ep)
			if err != nil {
				errors <- err
				return
			}
			results <- result
		}(i, endpoint)
	}

	// Collect results
	successCount := 0
	errorCount := 0

	for i := 0; i < len(endpoints); i++ {
		select {
		case result := <-results:
			fmt.Printf("SUCCESS: %s\n", result)
			successCount++
		case err := <-errors:
			fmt.Printf("ERROR: %v\n", err)
			errorCount++
		}
	}

	fmt.Printf("API calls completed: %d success, %d errors\n", successCount, errorCount)
}

// contextHierarchy demonstrates context hierarchy and cancellation
func contextHierarchy() {
	fmt.Println("\n=== Context Hierarchy Demo ===")

	rootCtx := context.Background()

	// Level 1: Service context
	serviceCtx, serviceCancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer serviceCancel()

	// Level 2: Request context
	requestCtx, requestCancel := context.WithCancel(serviceCtx)
	defer requestCancel()

	// Level 3: Database context
	dbCtx, dbCancel := context.WithTimeout(requestCtx, 2*time.Second)
	defer dbCancel()

	// Start database operation
	go func() {
		fmt.Println("Database: Starting transaction")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Database: Transaction completed")
		case <-dbCtx.Done():
			fmt.Printf("Database: Transaction cancelled: %v\n", dbCtx.Done())
		}
	}()

	// Start request processing
	go func() {
		fmt.Println("Request: Processing started")
		time.Sleep(1 * time.Second)

		// Simulate request-level cancellation
		fmt.Println("Request: Cancelling request context")
		requestCancel()
	}()

	// Wait for service timeout
	<-serviceCtx.Done()
	fmt.Printf("Service: All operations completed: %v\n", serviceCtx.Err())
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Context & Timeout Demo ===\n")

	// 1. Basic context cancellation
	fmt.Println("1. Basic Context Cancellation:")
	ctx, cancel := context.WithCancel(context.Background())

	go workerWithContext(ctx, 1)
	go workerWithContext(ctx, 2)

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling context...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println()

	// 2. Context with timeout
	fmt.Println("2. Context with Timeout:")
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer timeoutCancel()

	err := processWithTimeout(timeoutCtx, "Task1", 500*time.Millisecond)
	fmt.Printf("Task1 result: %v\n", err)

	err = processWithTimeout(timeoutCtx, "Task2", 1500*time.Millisecond)
	fmt.Printf("Task2 result: %v\n\n", err)

	// 3. Context with values
	fmt.Println("3. Context with Values:")
	contextWithValue(context.Background())
	fmt.Println()

	// 4. Multiple API calls with timeout
	concurrentAPICalls()

	// 5. Chained contexts
	chainedContexts()

	// 6. Context hierarchy
	contextHierarchy()

	fmt.Println("\n=== Demo Complete ===")
}
