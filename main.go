package main

import (
	"fmt"
	"money-system/user"
)

func main() {
	user1 := &user.User{
		ID:      "1",
		Name:    "User 1",
		Balance: 1000,
	}

	user2 := &user.User{
		ID:      "2",
		Name:    "User 2",
		Balance: 500,
	}

	user1.Deposit(200)
	user2.Withdraw(50)

	fmt.Printf("UserID: 1, Баланс: %.2f\n", user1.Balance)
	fmt.Printf("UserID: 2, Баланс: %.2f\n", user2.Balance)
}