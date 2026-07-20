package paysystem

import (
	"fmt"
	"money-system/user"
	"sync"
)

type Transaction struct {
	FromUID string
	ToUID   string
	Amount     float64
}

type PaymentSystem struct {
	users map[string]*user.User
	TransactionQueue []Transaction
}

func (ps *PaymentSystem) AddUser(u *user.User) {
	if ps.users == nil {
		ps.users = make(map[string]*user.User)
	}
	ps.users[u.ID] = u
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	if ps.TransactionQueue == nil {
		ps.TransactionQueue = make([]Transaction, 0, 10)
	}
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem) ProcessingTransaction(t Transaction) error {
	fromUser, ok := ps.users[t.FromUID]
	if !ok {
		return fmt.Errorf(
			"Пользователь %v не найден",
			 t.FromUID)
	}

	toUser, ok := ps.users[t.ToUID]
	if !ok {
		return fmt.Errorf(
			"Пользователь %v не найден",
			 t.ToUID)
	}


	if !fromUser.Withdraw(t.Amount) {
		return fmt.Errorf(
			"Недостаточно средств у пользователя %v",
			 t.FromUID)
	}


	toUser.Deposit(t.Amount)

	return nil
}


func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for t := range ch {
		err := ps.ProcessingTransaction(t)
		if err != nil {
			fmt.Println("Ошибка при обработке транзакции:", err)
			continue
		}

		fmt.Printf("Транзакция обработана: %s -> %s, сумма: %.2f\n", t.FromUID, t.ToUID, t.Amount)
	}
}