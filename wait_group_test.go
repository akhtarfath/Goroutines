package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// WaitGroup is used to wait for all Goroutines to finish executing
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // When the Goroutine is finished, the Done function is executed
	group.Add(1)       // Add 1 to the group

	fmt.Println("Asynchronous Goroutine") // Print the Goroutine
	time.Sleep(1 * time.Second)           // Wait for 1 second
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{} // Create a WaitGroup struct

	for i := 0; i < 10; i++ {
		go RunAsynchronous(group) // Create 10 Goroutines
	}

	group.Wait()                                     // Wait for all Goroutines to finish executing, and then continue to execute the following code
	fmt.Println("All Goroutines finished executing") // Print the final balance
}
