package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditionMutex(t *testing.T) { // Test function
	var x = 0                   // x is used to count the number of Goroutines that have been executed
	var mutex = &sync.Mutex{}   // mutex is used to lock the Goroutine
	for i := 0; i < 1000; i++ { // Create 1000 Goroutines
		go func() { // Goroutine function
			for j := 1; j <= 100; j++ { // create 100 data
				mutex.Lock()   // Lock the Goroutine
				x += 1         // x is used to count the number of Goroutines that have been executed
				mutex.Unlock() // Unlock the Goroutine
			}
		}() // Goroutine is created, the time is not known when it will be executed
	}
	time.Sleep(5 * time.Second) // Wait for 5 seconds for the Goroutine to be executed
	fmt.Println("Counter:", x)
}
