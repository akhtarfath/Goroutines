package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // Create a ticker, which will be executed every 1 second
	tickStop := false                         // Set the ticker to stop

	go func() { // Create a Goroutine
		time.Sleep(5 * time.Second) // Wait for 5 seconds
		ticker.Stop()               // Stop the ticker
		tickStop = true             // Set the ticker to stop
	}()

	//for tick := range ticker.C { // Wait for the ticker to be executed
	//	fmt.Println("Ticker executed at", tick) // Print the execution time
	//}

	for { // Wait for the ticker to be executed
		select { // Use select to wait for the ticker to be executed
		case tick := <-ticker.C: // Wait for the ticker to be executed
			fmt.Println("Ticker executed at", tick) // Print the execution time
		}
		if tickStop { // If the ticker is stopped,
			break // Exit the loop
		}
	}
}
