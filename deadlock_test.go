package Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}
func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	var user1 = UserBalance{
		Name:    "Fathan",
		Balance: 1000000,
	}
	var user2 = UserBalance{
		Name:    "Yustika",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000) // Transfer 100000 from user1 to user2
	go Transfer(&user2, &user1, 200000) // Transfer 200000 from user2 to user1
	// Deadlock occurs when the Goroutine is locked and cannot be unlocked

	time.Sleep(3 * time.Second)

	fmt.Println("User 1:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User 2:", user2.Name, "Balance:", user2.Balance)
}
