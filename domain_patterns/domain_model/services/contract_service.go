package services

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/domain_models"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/gateway"
)

func CreateContract(productID uint, revenue float64, dateSigned time.Time) (*domain_models.Contract, error) {
	var err error

	var product *domain_models.Product
	if product, err = FindProductByID(productID); err != nil {
		return nil, err
	}

	var contract *domain_models.Contract
	contract = domain_models.NewContract(productID, revenue, dateSigned)

	if err = gateway.Gateway().CreateContract(contract); err != nil {
		return contract, err
	}

	var revenueRecognitionStrategy = product.RevenueRecognitionStrategy

	if _, err = CreateRevenueRecognitions(revenueRecognitionStrategy, contract.ContractID, revenue, dateSigned); err != nil {
		return nil, err
	}

	return contract, nil
}

func CalculateRevenueRecognitions(contractID uint, date time.Time) (float64, error) {
	var err error
	var revenueRecognitions []*domain_models.RevenueRecognition
	if revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contractID, date); err != nil {
		return 0, err
	}

	var totalRevenue float64
	for _, revenueRecognition := range revenueRecognitions {
		totalRevenue += revenueRecognition.Amount
	}

	return totalRevenue, nil
}
