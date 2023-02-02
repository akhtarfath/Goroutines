package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done() // When the Goroutine is finished, the Done function is executed

	group.Add(1)             // Add 1 to the group
	data.Store(value, value) // Store data to the Map, the key and value are the same
}

func TestMap(t *testing.T) {
	data := &sync.Map{} // Create a Map struct
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group) // Create 100 Goroutines
	}

	time.Sleep(3 * time.Second)                    // Wait for 3 seconds for the Goroutine to be executed
	data.Range(func(key, value interface{}) bool { // Range is used to traverse the Map
		fmt.Println(key, ":", value)
		return true
	})
}
