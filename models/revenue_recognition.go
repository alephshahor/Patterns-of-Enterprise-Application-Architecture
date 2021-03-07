package models

import "time"

type RevenueRecognition struct {
	ContractID   uint      `pg:"contract_id"`
	Amount       float64   `pg:"amount"`
	RecognizedOn time.Time `pg:"recognized_on"`
}
