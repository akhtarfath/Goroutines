package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond is used to synchronize Goroutines, and can be used to implement a blocking queue,
// which is similar to the wait and notify mechanism in Java
var locker = sync.Mutex{}        // Create a Mutex struct
var cond = sync.NewCond(&locker) // Create a Cond struct
var group = sync.WaitGroup{}     // Create a WaitGroup struct

func WaitCondition(value int) {
	defer group.Done() // When the Goroutine is finished, the Done function is executed
	group.Add(1)       // Add 1 to the group, the number of Goroutines to be created

	cond.L.Lock() // Lock the Goroutine, so that only one Goroutine can be executed at a time
	cond.Wait()   // Wait for the signal, if the signal is not received, the Goroutine will be blocked
	fmt.Println("Done:", value)
	cond.L.Unlock() // Unlock the Goroutine, so that other Goroutines can be executed
}

func TestCond(t *testing.T) { // The output is not guaranteed to be in order
	for i := 0; i < 10; i++ { // Create 10 Goroutines
		go WaitCondition(i) // Create 10 Goroutines
	}

	go func() {                   // Create a Goroutine
		for i := 0; i < 10; i++ { // Create 10 Goroutines
			time.Sleep(1 * time.Second) // Wait for 1 second
			cond.Signal()               // Send a signal to one Goroutine
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second) // Wait for 1 second
	//	cond.Broadcast()            // Send a signal to all Goroutines
	//}()

	group.Wait() // Wait for all Goroutines to finish executing, and then continue to execute the following code
	fmt.Println("All Goroutines finished executing")
}
