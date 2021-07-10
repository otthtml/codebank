package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{
		ID:              uuid.NewV4().String(),
		Name:            "",
		Number:          "",
		ExpirationMonth: 0,
		ExpirationYear:  0,
		CVV:             0,
		Balance:         0,
		Limit:           0,
		CreatedAt:       time.Now(),
	}
	return c
}
