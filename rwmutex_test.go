package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) { // AddBalance is used to add the balance to the
	account.RWMutex.Lock() // Lock the Goroutine, Lock is used to lock the Goroutine when writing data
	account.Balance += amount
	account.RWMutex.Unlock() // Unlock the Goroutine, Unlock is used to unlock the Goroutine when writing data
}

func (account *BankAccount) GetBalance() int { // GetBalance is used to get the balance from the
	account.RWMutex.RLock() // Lock the Goroutine, RLock is used to lock the Goroutine when reading data
	balance := account.Balance
	account.RWMutex.RUnlock() // Unlock the Goroutine, RUnlock is used to unlock the Goroutine when reading data
	return balance
}

// RWMutex is used to lock the Goroutine
func TestReadWriteMutex(t *testing.T) { // Test function
	account := BankAccount{}   // Create a BankAccount struct
	for i := 0; i < 100; i++ { // Create 100 Goroutines
		go func() { // Goroutine function
			for j := 0; j < 100; j++ { // create 100 data
				account.AddBalance(1)                         // Add 1 to the balance
				fmt.Println("Balance:", account.GetBalance()) // Print the balance
			}
		}() // Goroutine is created, the time is not known when it will be executed
	}
	time.Sleep(5 * time.Second)                         // Wait for 5 seconds for the Goroutine to be executed
	fmt.Println("Final Balance:", account.GetBalance()) // Print the final balance
}
