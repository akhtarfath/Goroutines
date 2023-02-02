package Goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}       // Once is used to execute a function only once
	group := sync.WaitGroup{} // Create a WaitGroup struct
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)       // Add 1 to the group
			defer group.Done() // When the Goroutine is finished, the Done function is executed
			once.Do(OnlyOnce)  // OnlyOnce will only be executed once, even if it is called multiple times
		}()
	}
	group.Wait() // Wait for all Goroutines to finish executing, and then continue to execute the following code
	fmt.Println("Counter:", counter)
}
