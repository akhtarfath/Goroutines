package Goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var group = sync.WaitGroup{}
	var counter int64 = 0 // Use the atomic package to ensure that the counter is incremented atomically
	for i := 0; i < 100; i++ {
		group.Add(1) // Add 1 to the group
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1) // Add 1 to the counter
			}
			group.Done() // When the Goroutine is finished, the Done function is executed
		}()
	}
	group.Wait() // Wait for all Goroutines to finish executing, and then continue to execute the following code
	fmt.Println("Counter:", counter)
}

func TestAtomicRaceCondition(t *testing.T) { // Test function
	var x int64 = 0 // x is used to count the number of Goroutines that have been executed
	group := sync.WaitGroup{}
	for i := 0; i < 1000; i++ { // Create 1000 Goroutines
		go func() { // Goroutine function
			group.Add(1)                // Add 1 to the group
			for j := 1; j <= 100; j++ { // create 100 data
				// x += 1 // x is used to count the number of Goroutines that have been executed
				atomic.AddInt64(&x, 1) // for primitive types, use the atomic package to ensure that the counter is incremented atomically
			}
			group.Done() // When the Goroutine is finished, the Done function is executed
		}() // Goroutine is created, the time is not known when it will be executed
	}
	group.Wait()               // Wait for 5 seconds for the Goroutine to be executed
	fmt.Println("Counter:", x) // the result is 100000 because the atomic package is used to ensure that the counter is incremented atomically
}
