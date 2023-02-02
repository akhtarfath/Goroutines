package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

// async await is a feature that is not yet available in Go, but it is expected to be available in the future
// channel is a data structure that can be used to communicate between Goroutines
// chan is a keyword to create a channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // Create a channel, the data type must be specified
	// defer is used to execute the function after the function is finished
	defer close(channel) // Close the channel to prevent the channel from being used again
	go func() {
		time.Sleep(2 * time.Second)     // Wait for 2 seconds (to make it easier to see the result
		channel <- "Muhammad A. Fathan" // Send data to the channel, the data type must match the channel data type
		// if not send data to the channel, the Goroutine will be blocked or deadlocked
		//  is a situation where the Goroutine is blocked because it is waiting for data from the channel
		fmt.Println("Finish to send a data") // Print Finish to send a data
	}() // () is used to execute the function, without () the function will not be executed. The function is executed in a Goroutine

	message := <-channel                  // Receive data from the channel, the data type must match the channel data type
	fmt.Println("Receive Data:", message) // Print Hello World
	time.Sleep(5 * time.Second)           // Wait for 5 seconds (to make it easier to see the result
}

// channel can be used as a parameter
func TestCreateChannelAsParameter(t *testing.T) {
	channel := make(chan string) // Create a channel, the data type must be specified
	go GiveMeResponse(channel)   // Send channel as a parameter

	data := <-channel                  // Receive data from the channel, the data type must match the channel data type
	fmt.Println("Receive Data:", data) // Print Hello World (the result is not printed first because the Goroutine is not executed immediately)
	defer close(channel)               // Close the channel to prevent the channel from being used again
}

func GiveMeResponse(channel chan string) { // chan as parameter is not using pointer (*chan string)
	time.Sleep(2 * time.Second)     // Wait for 2 seconds (to make it easier to see the result
	channel <- "Muhammad A. Fathan" // Send data to the channel, the data type must match the channel data type
}
