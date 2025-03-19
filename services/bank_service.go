package services

import (
	"century/models/entities"
	"errors"
	"sync"
)

type BankService struct {
	accounts map[string]*entities.Account
	mu       sync.RWMutex
}

func NewBankService() *BankService {
	return &BankService{
		accounts: map[string]*entities.Account{
			"Mark": {AccountID: "Mark", Balance: 100},
			"Jane": {AccountID: "Jane", Balance: 50},
			"Adam": {AccountID: "Adam", Balance: 0},
		},
	}
}

func (b *BankService) Transfer(fromID, toID string, amount float64) error {
	if fromID == toID {
		return errors.New("cannot transfer money to the same account")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	fromAcc, exists := b.accounts[fromID]
	if !exists {
		return errors.New("sender account not found")
	}

	toAcc, exists := b.accounts[toID]
	if !exists {
		return errors.New("receiver account not found")
	}

	if fromAcc.Balance < amount {
		return errors.New("insufficient funds")
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount
	return nil
}

func (b *BankService) GetBalance(accountID string) (float64, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	acc, exists := b.accounts[accountID]
	if !exists {
		return 0, errors.New("account not found")
	}
	return acc.Balance, nil
}
