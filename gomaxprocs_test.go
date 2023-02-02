package Goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{} // Create a WaitGroup struct
	for i := 0; i < 100; i++ {
		group.Add(1) // Add 1 to the group
		go func() {
			time.Sleep(3 * time.Second)
			group.Done() // When the Goroutine is finished, the Done function is executed
		}()
	}

	cpuTotal := runtime.NumCPU() // Get the number of CPU cores
	fmt.Println("CPU Total:", cpuTotal)

	threadTotal := runtime.GOMAXPROCS(-1) // Get the number of threads
	fmt.Println("Thread Total:", threadTotal)

	goroutineTotal := runtime.NumGoroutine() // Get the number of Goroutines
	fmt.Println("Goroutine Total:", goroutineTotal)
	// first goroutine to running is main goroutine
	// second goroutine to garbage collection goroutine (GC)
	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	runtime.GOMAXPROCS(20)                // Set the number of threads to 20
	threadTotal := runtime.GOMAXPROCS(-1) // Get the number of threads
	fmt.Println("Thread Total:", threadTotal)
}
