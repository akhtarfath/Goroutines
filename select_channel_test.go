package Goroutines

import (
	"fmt"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0 // counter is used to count the number of channels that have data
	for {        // for is used to create an infinite loop
		select { // select is used to select a channel that has data
		case data := <-channel1: // <- is used to receive data from the channel
			fmt.Println("Receive Data From Channel 1:", data)
			counter++
		case data := <-channel2: // <- is used to receive data from the channel
			fmt.Println("Receive Data From Channel 2:", data)
			counter++
		default: // default is used to execute the code if there is no channel that has data
			fmt.Println("Menunggu Data")
		}
		if counter == 2 { // if the number of channels that have data is 2, then the loop will
			break // break is used to exit the loop
		}
	}
}
