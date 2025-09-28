package main

import (
	"fmt"
	"time"
)

// Basic goroutine example
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// Goroutine with parameters
func printNumbers(prefix string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// Goroutine with channel communication
func producer(ch chan<- int, max int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i // Send to channel
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

func consumer(ch <-chan int, done chan<- bool) {
	for value := range ch {
		fmt.Printf("Consuming: %d\n", value)
		time.Sleep(300 * time.Millisecond) // Simulate processing
	}
	done <- true // Signal completion
}

// Buffered channels
func demonstrateBufferedChannels() {
	// Buffered channel with capacity 3
	ch := make(chan int, 3)

	// Can send without blocking (up to capacity)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Sent 3 values to buffered channel")

	// Receive values
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
}

// Channel directions
func sendOnly(ch chan<- int, value int) {
	ch <- value
}

func receiveOnly(ch <-chan int) int {
	return <-ch
}

// Select statement
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Select waits for first available channel
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout: no message received")
	}
}

// Multiple select cases
func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for {
			select {
			case msg, ok := <-ch1:
				if !ok {
					ch1 = nil // Disable this case
					continue
				}
				out <- msg
			case msg, ok := <-ch2:
				if !ok {
					ch2 = nil // Disable this case
					continue
				}
				out <- msg
			}

			// Exit when both channels are closed
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()

	return out
}

// Timeout with select
func timeoutExample() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Operation completed"
	}()

	select {
	case result := <-ch:
		fmt.Println("Result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
}

// Non-blocking operations
func nonBlockingExample() {
	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 42:
		fmt.Println("Sent value successfully")
	default:
		fmt.Println("Channel is full, couldn't send")
	}

	// Non-blocking receive
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("Channel is empty, nothing to receive")
	}

	// Try to receive again (should be empty)
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("Channel is empty")
	}
}

// Worker pool pattern (basic)
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // Simulate work
		results <- job * 2                 // Send result
	}
}

func workerPoolExample() {
	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to signal no more work

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

// Goroutine leak demonstration (bad practice)
func leakyGoroutine() {
	ch := make(chan int)

	go func() {
		// This goroutine will run forever if not properly managed
		for {
			select {
			case value := <-ch:
				fmt.Printf("Received: %d\n", value)
			}
		}
	}()

	// Send one value
	ch <- 42
	time.Sleep(100 * time.Millisecond)

	// The goroutine is still running, waiting for more values
	fmt.Println("Goroutine is still running (leak!)")
}

// Proper goroutine cleanup
func properGoroutine() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case value, ok := <-ch:
				if !ok {
					fmt.Println("Channel closed, goroutine exiting")
					return
				}
				fmt.Printf("Received: %d\n", value)
			}
		}
	}()

	// Send value
	ch <- 42
	time.Sleep(100 * time.Millisecond)

	// Properly close channel to signal goroutine to exit
	close(ch)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Goroutine properly cleaned up")
}

func main() {
	fmt.Println("=== Basic Goroutines ===")

	// Basic goroutine
	go sayHello()
	time.Sleep(100 * time.Millisecond) // Give goroutine time to run

	// Goroutines with parameters
	go printNumbers("Goroutine A", 3)
	go printNumbers("Goroutine B", 3)
	time.Sleep(500 * time.Millisecond) // Wait for completion

	fmt.Println("\n=== Channels ===")

	// Basic channel communication
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch, 3)
	go consumer(ch, done)

	<-done // Wait for consumer to finish
	fmt.Println("Producer-consumer finished")

	fmt.Println("\n=== Buffered Channels ===")
	demonstrateBufferedChannels()

	fmt.Println("\n=== Channel Directions ===")
	bidirectional := make(chan int, 1)

	go sendOnly(bidirectional, 100)
	value := receiveOnly(bidirectional)
	fmt.Printf("Received via directional channels: %d\n", value)

	fmt.Println("\n=== Select Statement ===")
	selectExample()

	fmt.Println("\n=== Fan-In Pattern ===")
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "Hello from ch1"
	ch1 <- "Another from ch1"
	close(ch1)

	ch2 <- "Hello from ch2"
	close(ch2)

	combined := fanIn(ch1, ch2)
	for msg := range combined {
		fmt.Println("Combined:", msg)
	}

	fmt.Println("\n=== Timeout Example ===")
	timeoutExample()

	fmt.Println("\n=== Non-blocking Operations ===")
	nonBlockingExample()

	fmt.Println("\n=== Worker Pool Pattern ===")
	workerPoolExample()

	fmt.Println("\n=== Goroutine Lifecycle ===")
	fmt.Println("--- Bad: Goroutine leak ---")
	leakyGoroutine()

	fmt.Println("\n--- Good: Proper cleanup ---")
	properGoroutine()

	fmt.Println("\n=== Concurrency Concepts Demonstrated ===")
	fmt.Println("✓ Goroutines: Lightweight threads")
	fmt.Println("✓ Channels: Communication between goroutines")
	fmt.Println("✓ Select: Multiplexing channel operations")
	fmt.Println("✓ Buffered channels: Asynchronous communication")
	fmt.Println("✓ Channel directions: Type safety")
	fmt.Println("✓ Worker pools: Divide work among goroutines")
	fmt.Println("✓ Proper cleanup: Avoid goroutine leaks")
}
