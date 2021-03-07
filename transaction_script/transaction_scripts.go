package transaction_script

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
)

func CreateContract(productID uint, revenue float64, dateSigned time.Time) error {
	var err error

	var product *models.Product
	if product, err = Gateway().FindProductByID(productID); err != nil {
		return err
	}

	var newContract = &models.Contract{
		ProductID:  productID,
		Revenue:    revenue,
		DateSigned: dateSigned,
	}

	if err = Gateway().CreateContract(newContract); err != nil {
		return err
	}

	var revenueRecognitions []*models.RevenueRecognition

	switch product.ProductType {

	case enums.WordProcessor:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue,
			DateSigned: dateSigned,
		})
	case enums.Spreadsheet:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned,
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned.AddDate(0, 0, 60),
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned.AddDate(0, 0, 90),
		})
	case enums.Database:
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned,
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned.AddDate(0, 0, 30),
		})
		revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
			ContractID: newContract.ContractID,
			ProductID:  productID,
			Revenue:    revenue / 3,
			DateSigned: dateSigned.AddDate(0, 0, 60),
		})
	}

	if err = Gateway().CreateRevenueRecognitions(revenueRecognitions); err != nil {
		return err
	}

	return nil
}

func CalculateRevenueRecognitions(contractID uint, date time.Time) (float64, error) {
	var err error
	var revenueRecognitions []*models.RevenueRecognition
	if revenueRecognitions, err = Gateway().FindRevenueRecognitionForContractBeforeDate(contractID, date); err != nil {
		return 0, err
	}

	var totalRevenue float64
	for _, revenueRecognition := range revenueRecognitions {
		totalRevenue = revenueRecognition.Revenue
	}

	return totalRevenue, nil
}
