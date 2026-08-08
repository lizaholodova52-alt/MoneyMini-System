package user

import (
	"fmt"
	"sync"
)

type User struct {
	ID     string
	Name   string
	Balance float64
	mu	 sync.Mutex
}


func (u *User) Deposit(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if amount <= 0 {
    	fmt.Println("invalid amount")
		return
	}
	
	u.Balance += amount
	
	fmt.Println("Депозит выполнен.")
}

func (u *User) Withdraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	
	if amount <= 0 {
    	return fmt.Errorf("invalid amount")
	}
	
	if u.Balance < amount {
    	return fmt.Errorf("insufficient funds")
	}
	
	u.Balance -= amount
	
	return nil
}