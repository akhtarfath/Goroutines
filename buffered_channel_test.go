package Goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// buffered channel is a channel that has a buffer, the buffer is a data structure that stores data
// buffered channel is used to store data before it is received by the receiver
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // Create a channel with a buffer of 3, the data type must be specified
	// without 3 the channel will be unbuffered, so we can send data to the channel without receiving it first

	go func() { // Goroutine 1
		channel <- "Muhammad Fathan A." // Send data to the channel, the data type must match the channel data type
		channel <- "Yustika Ramadhani"  // Send data to the channel, the data type must match the channel data type
		channel <- "john Cena"          // Send data to the channel, the data type must match the channel data type
		// we can send data again because the buffer max capacity is 3
	}()

	go func() { // Goroutine 2
		fmt.Println("Length:", len(channel))   // length is the number of data in the buffer
		fmt.Println("Capacity:", cap(channel)) // capacity is the maximum number of data in the buffer

		// if the buffer is full, the Goroutine will be blocked or deadlocked
		for i := 0; i < len(channel); i++ {
			fmt.Println("Data:", <-channel) // Receive data from the channel, the data type must match the channel data type
		}
	}()

	time.Sleep(5 * time.Second) // Wait for 5 seconds (to make it easier to see the result
	defer close(channel)        // Close the channel to prevent the channel from being used again
	// if the buffer is empty, the Goroutine will be blocked or deadlocked because there is no data to be received
	fmt.Println("Finish") // Print Finish
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i+1)
		}
		close(channel) // Close the channel to prevent the channel from being used again
		// why to close channel? because if the channel is not closed, the range will not stop
	}()

	for data := range channel { // range is used to receive data from the channel
		fmt.Println("Data:", data)
	}
	fmt.Println("Finish")
}
