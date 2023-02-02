package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) { // chan receive only
	time.Sleep(2 * time.Second)     // Wait for 2 seconds (to make it easier to see the result
	channel <- "Muhammad A. Fathan" // Send data to the channel, the data type must match the channel data type
}

func OnlyOut(channel <-chan string) { // chan send only
	data := <-channel                  // Receive data from the channel, the data type must match the channel data type
	fmt.Println("Receive Data:", data) // Print Hello World (the result is not printed first because the Goroutine is not executed immediately)
}

func TestInOutChannel(t *testing.T) { // Test function
	channel := make(chan string) // Create a channel, the data type must be specified

	// Send channel as a parameter
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second) // Wait for 3 seconds (to make it easier to see the result
	defer close(channel)        // Close the channel to prevent the channel from being used again
}
