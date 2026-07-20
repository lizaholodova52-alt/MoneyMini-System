package paysystem

import (
	"fmt"
	"money-system/user"
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

func (ps *PaymentSystem) ProcessingTransactions() error {
	for _, t := range ps.TransactionQueue {


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
	}

	return nil
}