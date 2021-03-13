package gateway

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
)

// Table Data Gateway (144)
type IGateway interface {
	CreateContract(productID uint, revenue float64, dateSigned time.Time) (uint, error)
	CreateRevenueRecognitions(revenueRecognitions []*models.RevenueRecognition) error
	FindContractByID(contractID uint) (*models.Contract, error)
	FindProductByID(productID uint) (*models.Product, error)
	FindRevenueRecognitionForContractBeforeDate(contractID uint, date time.Time) ([]*models.RevenueRecognition, error)
}

type gateway struct{}

var gatewayInstance *gateway

func Gateway() *gateway {
	if gatewayInstance == nil {
		gatewayInstance = newGateway()
	}
	return gatewayInstance
}

func newGateway() *gateway {
	return &gateway{}
}

func (g gateway) CreateContract(productID uint, revenue float64, dateSigned time.Time) (uint, error) {
	var err error
	var contract = &models.Contract{
		ProductID:  productID,
		Revenue:    revenue,
		DateSigned: dateSigned,
	}

	if _, err = db.DB().Model(contract).Insert(); err != nil {
		return 0, err
	}
	return contract.ContractID, nil
}

func (g gateway) CreateRevenueRecognitions(revenueRecognitions []*models.RevenueRecognition) error {
	var err error
	if _, err = db.DB().Model(&revenueRecognitions).
		Insert(); err != nil {
		return err
	}
	return nil
}

func (g gateway) FindContractByID(contractID uint) (*models.Contract, error) {
	var err error
	var newContract = new(models.Contract)
	if err = db.DB().Model(newContract).
		Where("contract_id = ?", contractID).
		Select(); err != nil {
		return nil, err
	}
	return newContract, nil
}

func (g gateway) FindProductByID(productID uint) (*models.Product, error) {
	var err error
	var product = new(models.Product)
	if err = db.DB().Model(product).
		Where("product_id = ?", productID).
		Select(); err != nil {
		return nil, err
	}
	return product, err
}

func (g gateway) FindRevenueRecognitionForContractBeforeDate(contractID uint, date time.Time) ([]*models.RevenueRecognition, error) {
	var err error
	var revenueRecognitions []*models.RevenueRecognition
	if err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contractID).
		Where("recognized_on <= ?", date).
		Select(); err != nil {
		return nil, err
	}
	return revenueRecognitions, nil
}
