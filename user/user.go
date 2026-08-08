package user

import "fmt"

type User struct {
	ID     string
	Name   string
	Balance float64
}


func (u *User) Deposit(amount float64) {

	if amount <= 0 {
    	fmt.Println("invalid amount")
		return
	}

	u.Balance += amount
	fmt.Println("Депозит выполнен.")
}

func (u *User) Withdraw(amount float64) error {

	if amount <= 0 {
    	return fmt.Errorf("invalid amount")
	}
	
	if u.Balance < amount {
    	return fmt.Errorf("insufficient funds")
	}

	u.Balance -= amount
	return nil
}