package domain_models

import "github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"

type Product struct {
	ProductID                  uint                        `pg:"product_id"`
	ProductName                string                      `pg:"product_name"`
	ProductType                enums.ProductType           `pg:"product_type"`
	RevenueRecognitionStrategy IRevenueRecognitionStrategy `pg:"-"`
}
