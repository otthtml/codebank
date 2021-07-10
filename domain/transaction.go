package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func newTransaction() *Transaction {
	t := &Transaction{
		ID:           uuid.NewV4().String(),
		Amount:       0,
		Status:       "",
		Description:  "",
		Store:        "",
		CreditCardId: "",
		CreatedAt:    time.Now(),
	}
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}
