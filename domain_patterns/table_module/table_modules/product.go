package table_modules

import "github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"

type productTableModule struct {
	dataset gateway.IGateway
}

func NewProductTableModule(dataset gateway.IGateway) *productTableModule {
	return &productTableModule{
		dataset: dataset,
	}
}
