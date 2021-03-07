package models

import "time"

type Contract struct {
	ContractID uint      `pg:"contract_id"`
	ProductID  uint      `pg:"product_id"`
	Revenue    float64   `pg:"revenue"`
	DateSigned time.Time `pg:"date_signed"`
}
