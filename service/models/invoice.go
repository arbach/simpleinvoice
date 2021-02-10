package models

import (
	"encoding/json"
	"time"
)

type Invoice struct {
	ID             int       `json:"id" db:"id"`
	Status         string    `json:"status" db:"status"`
	Description    string    `json:"description" db:"description"`
	Amount         float64   `json:"amount" db:"amount"`
	PaymentAddress string    `json:"paymentAddress" db:"payment_address"`
	PaidAmount     float64   `json:"paidAmount" db:"paid_amount"`
	UpdatedAt      time.Time `json:"-" db:"updated_at"`
	CreatedAt      time.Time `json:"-" db:"created_at"`
}

type InvoiceRequest struct {
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentAddress string  `json:"-"`
}

func (in Invoice) MarshalJSON() ([]byte, error) {
	in.SetStatus()

	type Alias Invoice
	alias := (Alias)(in)

	return json.Marshal(struct {
		Alias
	}{
		Alias: alias,
	})
}

func (in *Invoice) isExpired() bool {
	return in.CreatedAt.Before(time.Now().Add(time.Hour * -1))
}

func (in *Invoice) SetStatus() {
	switch {
	case in.PaidAmount > in.Amount:
		in.Status = "Overpaid"
		break
	case in.PaidAmount == in.Amount:
		in.Status = "Paid"
		break
	case in.PaidAmount < in.Amount && in.isExpired():
		in.Status = "Expired"
		break
	case in.PaidAmount > 0 && in.PaidAmount < in.Amount && !in.isExpired():
		in.Status = "Partially paid"
		break
	case in.PaidAmount == 0:
		in.Status = "Unpaid"
		break
	default:
		in.Status = "Unpaid"
	}
}
