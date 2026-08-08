package paysystem

import (
	"fmt"
	"money-system/user"
	"sync"
)

type Transaction struct {
	FromID string
	ToID   string
	Amount     float64
}

type PaymentSystem struct {
	Users map[string]*user.User
	TransactionQueue []Transaction
}

func (ps *PaymentSystem) AddUser(u *user.User) {
	if ps.Users == nil {
		ps.Users = make(map[string]*user.User)
	}
	ps.Users[u.ID] = u
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	if ps.TransactionQueue == nil {
		ps.TransactionQueue = make([]Transaction, 0, 10)
	}
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem) ProcessTransaction(t Transaction) error {
	fromUser, ok := ps.Users[t.FromID]
	if !ok {
		return fmt.Errorf(
			"User %s not found",
			 t.FromID)
	}

	toUser, ok := ps.Users[t.ToID]
	if !ok {
		return fmt.Errorf(
			"User %s not found",
			 t.ToID)
	}


	err := fromUser.Withdraw(t.Amount)
	if err != nil {
		return fmt.Errorf("error withdrawing from user %s: %w",
			t.FromID,
			err)
	}


	toUser.Deposit(t.Amount)

	return nil
}


func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for t := range ch {
		err := ps.ProcessTransaction(t)
		if err != nil {
			fmt.Println("transaction error:", err)
			continue
		}

		fmt.Printf("Транзакция обработана: %s -> %s, сумма: %.2f\n", t.FromID, t.ToID, t.Amount)
	}
}