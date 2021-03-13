package table_modules

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
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
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned); err != nil {
			return 0, err
		}
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, 30)); err != nil {
			return 0, err
		}
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, 60)); err != nil {
			return 0, err
		}
	case enums.Spreadsheet:
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned); err != nil {
			return 0, err
		}
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, 60)); err != nil {
			return 0, err
		}
		if err = revenueRecognitionTableModule.Create(contractID, revenue/3, dateSigned.AddDate(0, 0, 90)); err != nil {
			return 0, err
		}
	}

	return contractID, nil
}

func (m *contractTableModule) CalculateRevenueRecognition(contractID uint) (float64, error) {
	var err error

	var revenueRecognitions []*models.RevenueRecognition
	// TODO: No puedes traer objetos, no existen instancias!
	if revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contractID, time.Now()); err != nil {
		return 0, err
	}

	var totalRevenueRecognition float64
	for _, revenueRecognition := range revenueRecognitions {
		// TODO: No puedes, no existen instancias!!?
		totalRevenueRecognition += revenueRecognition.Amount
	}

	return totalRevenueRecognition, err
}
