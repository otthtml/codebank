package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

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
