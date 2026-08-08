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

	// Депозиты
	user1.Deposit(200)
	fmt.Printf("UserID: %s, Баланс: %.2f\n", user1.ID, user1.Balance)

	user1.Deposit(100)
	fmt.Printf("UserID: %s, Баланс: %.2f\n", user1.ID, user1.Balance)

	// Снятия
	if err := user2.Withdraw(50); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("UserID: %s, Баланс: %.2f\n", user2.ID, user2.Balance)

	if err := user2.Withdraw(100); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("UserID: %s, Баланс: %.2f\n", user2.ID, user2.Balance)
}