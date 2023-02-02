package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{} // Create a WaitGroup struct
	group.Add(1)              // Add 1 to the group

	time.AfterFunc(1*time.Second, func() { // Create a timer, the timer will be executed after 5 seconds
		fmt.Println("Timer after 1 seconds")
		group.Done() // When the Goroutine is finished, the Done function is executed
	})

	group.Wait() // Wait for all Goroutines to finish executing, and then continue to execute the following code
	fmt.Println("All Goroutines finished executing")
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)       // Create a timer, the timer will be executed after 5 seconds
	timerWithAfter := time.After(5 * time.Second) // Create a timer, the timer will be executed after 5 seconds
	fmt.Println("Current Time", time.Now())       // Print the current times

	times := <-timer.C         // Wait for the timer to be executed
	times2 := <-timerWithAfter // Wait for the timer to be executed
	fmt.Println("Timer executed at", times)
	fmt.Println("Timer executed at", times2)
}
