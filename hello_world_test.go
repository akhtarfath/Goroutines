package Goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) { // Test function
	t.Run("Create Goroutine", TestCreateGoroutine) // Create Goroutine
	t.Run("Many Goroutine", TestManyGoroutine)     // Many Goroutine
}

func RunHelloWorld() { // Goroutine function
	println("Hello World") // Print Hello World
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()          // Goroutine is created, the time is not known when it will be executed
	println("Ups")              // This line is executed first because the Goroutine is not executed immediately
	time.Sleep(1 * time.Second) // Wait for 1 second for the Goroutine to be executed
}

func DisplayNumber(number int) { // Goroutine function
	fmt.Println("DISPLAY:", number) // Print the number
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i <= 1000000; i++ { // Create 1000000 Goroutines
		go DisplayNumber(i) // Goroutine is created, the time is not known when it will be executed
		// the result is not printed in order, because the result is printed when the Goroutine is executed
	}
	time.Sleep(10 * time.Second) // Wait for 10 seconds for the Goroutine to be executed
}
