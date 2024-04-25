package entities

import (
	"gorm.io/gorm"
)

type (
	Payment struct {
		gorm.Model
		Type         string  `json:"type"`
		Total        float64 `json:"total"`
		ProofPayment string  `json:"proof_payment"`
		Status       string  `json:"status"`
	}

	InsertPayment struct {
		Type         string  `json:"type"`
		Total        float64 `json:"total"`
		ProofPayment string  `json:"proof_payment"`
		Status       string  `json:"status"`
	}
)
