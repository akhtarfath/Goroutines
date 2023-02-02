package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool Design Pattern is a design pattern that is used to manage a group of objects that are expensive to create.
// and can be reused. The objects that are managed by the Pool Design Pattern are called Pools
func TestPool(t *testing.T) {
	pool := sync.Pool{
		// default value
		New: func() interface{} { // New is used to create a new object, if the object is not available in the pool
			return "New" // Create a new object
		},
	} // Create a Pool struct

	pool.Put("Fathan")    // Put 1 into the pool
	pool.Put("Yustika")   // Put 2 into the pool
	pool.Put("Aulia")     // Put 3 into the pool
	pool.Put("Ramadhani") // Put 4 into the pool

	for i := 0; i < 10; i++ {
		go func() { // Goroutine function
			data := pool.Get() // Get the data from the pool
			fmt.Println(data)
			time.Sleep(1 * time.Second) // Wait for 1 second
			pool.Put(data)              // Put the data back into the pool
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("All Goroutines finished executing")
}
