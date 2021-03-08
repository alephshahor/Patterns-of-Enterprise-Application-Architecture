package domain_models

import (
	"time"
)

type Contract struct {
	ContractID uint      `pg:"contract_id"`
	ProductID  uint      `pg:"product_id"`
	Revenue    float64   `pg:"revenue"`
	DateSigned time.Time `pg:"date_signed"`
}

func NewContract(productID uint, revenue float64, dateSigned time.Time) *Contract {
	var contract = &Contract{
		ProductID:  productID,
		Revenue:    revenue,
		DateSigned: dateSigned,
	}
	return contract
}
