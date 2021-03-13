package table_modules

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
)

type productTableModule struct {
	dataset gateway.IGateway
}

func NewProductTableModule(dataset gateway.IGateway) *productTableModule {
	return &productTableModule{
		dataset: dataset,
	}
}

func (m *productTableModule) FindProductType(productID uint) (enums.ProductType, error) {
	var err error

	var productType enums.ProductType
	productType, err = m.dataset.FindProductType(productID)

	return productType, err
}
