package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) { // Test function
	var x = 0                   // x is used to count the number of Goroutines that have been executed
	for i := 0; i < 1000; i++ { // Create 1000 Goroutines
		go func() { // Goroutine function
			for j := 1; j <= 100; j++ { // create 100 data
				x += 1 // x is used to count the number of Goroutines that have been executed
			}
		}() // Goroutine is created, the time is not known when it will be executed
	}
	time.Sleep(5 * time.Second) // Wait for 5 seconds for the Goroutine to be executed
	fmt.Println("Counter:", x)
}
