package table_modules

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
)

type contractTableModule struct {
	dataset gateway.IGateway
}

func NewContractTableModule(dataset gateway.IGateway) *contractTableModule {
	return &contractTableModule{
		dataset: dataset,
	}
}

func (m *contractTableModule) Create(productID uint, revenue float64, dateSigned time.Time) (uint, error) {
	var err error

	var contractID uint
	if contractID, err = m.dataset.CreateContract(productID, revenue, dateSigned); err != nil {
		return 0, err
	}

	var productTableModule *productTableModule
	productTableModule = NewProductTableModule(m.dataset)
	var productType enums.ProductType
	if productType, err = productTableModule.FindProductType(productID); err != nil {
		return 0, err
	}

	var revenueRecognitionTableModule *revenueRecognitionTableModule
	revenueRecognitionTableModule = NewRevenueRecognitionTableModule(m.dataset)

	switch productType {
	case enums.WordProcessor:
		if err = revenueRecognitionTableModule.Create(contractID, revenue, dateSigned); err != nil {
			return 0, err
		}
	case enums.Database:
		var dateRanges = []int{0, 30, 60}
		for _, date := range dateRanges {
			if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, date)); err != nil {
				return 0, err
			}
		}
	case enums.Spreadsheet:
		var dateRanges = []int{0, 60, 90}
		for _, date := range dateRanges {
			if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, date)); err != nil {
				return 0, err
			}
		}
	}

	return contractID, nil
}
