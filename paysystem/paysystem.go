package paysystem

import (
	"fmt"
	"money-system/user"
	"sync"
)

type Transaction struct {
	FromUserID string
	ToUserID   string
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
	fromUser, ok := ps.users[t.FromUserID]
	if !ok {
		return fmt.Errorf(
			"Пользователь %v не найден",
			 t.FromUserID)
	}

	toUser, ok := ps.users[t.ToUserID]
	if !ok {
		return fmt.Errorf(
			"Пользователь %v не найден",
			 t.ToUserID)
	}


	if !fromUser.Withdraw(t.Amount) {
		return fmt.Errorf(
			"Недостаточно средств у пользователя %v",
			 t.FromUserID)
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

		fmt.Printf("Транзакция успешно обработана: %s -> %s, сумма: %.2f\n", t.FromUserID, t.ToUserID, t.Amount)
	}
}