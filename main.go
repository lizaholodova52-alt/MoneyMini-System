package main

import (
	"fmt"
	"money-system/user"
	"money-system/paysystem"
)

func main() {
	
	ps := &paysystem.PaymentSystem{}
	
	fmt.Println("Создаю UserID: 1 с балансом 1000")
	fmt.Println("Создаю UserID: 2 с балансом 500")


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
	ps.AddUser(user1)
	ps.AddUser(user2)

	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму в размере 200")
	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму в размере 50")


	ps.AddTransaction(paysystem.Transaction{
		FromUID: "1",
		ToUID:   "2",
		Amount:     200,
	})
	ps.AddTransaction(paysystem.Transaction{
		FromUID: "2",
		ToUID:   "1",
		Amount:     50,
	})

	
	err := ps.ProcessingTransactions()
	if err != nil {
		fmt.Println("Ошибка при обработке транзакций:", err)
	}

	// Подсказка
	fmt.Println("Итого")
	fmt.Println("У первого пользователя должно получиться 850")
	fmt.Println("У второго пользователя должно получиться 650")
}