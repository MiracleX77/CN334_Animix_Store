package models

import "time"

type InsertPaymentModel struct {
	Type         string  `json:"type"`
	Total        float64 `json:"total"`
	ProofPayment string  `json:"proof_payment"`
}

type PaymentModel struct {
	ID           uint64    `json:"id"`
	Type         string    `json:"type"`
	Total        float64   `json:"total"`
	ProofPayment string    `json:"proof_payment"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
