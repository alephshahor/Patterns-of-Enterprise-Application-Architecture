package scripts

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/transaction_script/gateway"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
)

func CreateContract(productID uint, revenue float64, dateSigned time.Time) (*models.Contract, error) {
	var err error

	var product *models.Product
	if product, err = gateway.Gateway().FindProductByID(productID); err != nil {
		return nil, err
	}

	var newContract = &models.Contract{
		ProductID:  productID,
		Revenue:    revenue,
		DateSigned: dateSigned,
	}

	if err = gateway.Gateway().CreateContract(newContract); err != nil {
		return nil, err
	}

	var revenueRecognitions []*models.RevenueRecognition

	switch product.ProductType {

	case enums.WordProcessor:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue,
			RecognizedOn: dateSigned,
		})
	case enums.Spreadsheet:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned,
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned.AddDate(0, 0, 60),
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned.AddDate(0, 0, 90),
		})
	case enums.Database:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned,
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned.AddDate(0, 0, 30),
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID:   newContract.ContractID,
			Amount:       revenue / 3,
			RecognizedOn: dateSigned.AddDate(0, 0, 60),
		})
	}

	if err = gateway.Gateway().CreateRevenueRecognitions(revenueRecognitions); err != nil {
		return nil, err
	}

	return newContract, nil
}

func CalculateRevenueRecognitions(contractID uint, date time.Time) (float64, error) {
	var err error
	var revenueRecognitions []*models.RevenueRecognition
	if revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contractID, date); err != nil {
		return 0, err
	}

	var totalRevenue float64
	for _, revenueRecognition := range revenueRecognitions {
		totalRevenue += revenueRecognition.Amount
	}

	return totalRevenue, nil
}
