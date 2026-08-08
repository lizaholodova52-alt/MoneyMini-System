package main

import (
	"fmt"
	"money-system/paysystem"
	"money-system/user"
	"sync"
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


	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму 200")
	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму 50")

	t1 := paysystem.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount:     200,
	}

	t2 := paysystem.Transaction{
		FromID: "2",
		ToID:   "1",
		Amount:     50,
	}

	
	ps.AddTransaction(t1)
	ps.AddTransaction(t2)


	ch := make(
		chan paysystem.Transaction,
		len(ps.TransactionQueue),
	)

	
	var wg sync.WaitGroup


	workersCount := 3

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go ps.Worker(ch, &wg)
	}

	
	for _, transaction := range ps.TransactionQueue {
		ch <- transaction
	}


	close(ch)


	wg.Wait()

	fmt.Println("Итого")
	fmt.Printf("UserID: 1, Баланс: %.2f\n", user1.Balance)
	fmt.Printf("UserID: 2, Баланс: %.2f\n", user2.Balance)
}