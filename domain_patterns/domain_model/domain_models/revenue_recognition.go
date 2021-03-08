package domain_models

import (
	"time"
)

type RevenueRecognition struct {
	ContractID   uint      `pg:"contract_id"`
	Amount       float64   `pg:"amount"`
	RecognizedOn time.Time `pg:"recognized_on"`
}

func NewRevenueRecognition(contractID uint, amount float64, recognizedOn time.Time) RevenueRecognition {
	var revenueRecognition = RevenueRecognition{
		ContractID:   contractID,
		Amount:       amount,
		RecognizedOn: recognizedOn,
	}
	return revenueRecognition
}
