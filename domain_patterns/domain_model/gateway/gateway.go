package gateway

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/domain_models"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
)

// Table Data Gateway (144)
type IGateway interface {
	CreateContract(contract *domain_models.Contract) error
	FindProductByID(productID uint) (*domain_models.Product, error)
	CreateRevenueRecognitions(revenueRecognitions []*domain_models.RevenueRecognition) error
	FindRevenueRecognitionForContractBeforeDate(contractID uint, date time.Time) ([]*domain_models.RevenueRecognition, error)
	FindContractByID(contractID uint) (*domain_models.Contract, error)
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

func (g gateway) CreateContract(contract *domain_models.Contract) error {
	var err error
	if _, err = db.DB().Model(contract).Insert(); err != nil {
		return err
	}
	return nil
}

func (g gateway) FindProductByID(productID uint) (*domain_models.Product, error) {
	var err error
	var product = new(domain_models.Product)
	if err = db.DB().Model(product).
		Where("product_id = ?", productID).
		Select(); err != nil {
		return nil, err
	}
	return product, err
}

func (g gateway) CreateRevenueRecognitions(revenueRecognitions []*domain_models.RevenueRecognition) error {
	var err error
	if _, err = db.DB().Model(&revenueRecognitions).
		Insert(); err != nil {
		return err
	}
	return nil
}

func (g gateway) FindContractByID(contractID uint) (*domain_models.Contract, error) {
	var err error
	var newContract = new(domain_models.Contract)
	if err = db.DB().Model(newContract).
		Where("contract_id = ?", contractID).
		Select(); err != nil {
		return nil, err
	}
	return newContract, nil
}

func (g gateway) FindRevenueRecognitionForContractBeforeDate(contractID uint, date time.Time) ([]*domain_models.RevenueRecognition, error) {
	var err error
	var revenueRecognitions []*domain_models.RevenueRecognition
	if err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contractID).
		Where("recognized_on <= ?", date).
		Select(); err != nil {
		return nil, err
	}
	return revenueRecognitions, nil
}
