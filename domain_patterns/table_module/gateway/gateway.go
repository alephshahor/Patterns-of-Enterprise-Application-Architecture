package gateway

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
)

// Table Data Gateway (144)
type IGateway interface {
	CreateContract(productID uint, revenue float64, dateSigned time.Time) (uint, error)
	CreateRevenueRecognition(contractID uint, amount float64, recognizedOn time.Time) error
	FindProductType(productID uint) (enums.ProductType, error)
	FindRevenueRecognitionAmount(contractID uint, recognizedOn time.Time) ([]float64, error)
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

func (g *gateway) CreateContract(productID uint, revenue float64, dateSigned time.Time) (uint, error) {
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

func (g *gateway) CreateRevenueRecognition(contractID uint, amount float64, recognizedOn time.Time) error {
	var err error
	var revenueRecognition = &models.RevenueRecognition{
		ContractID:   contractID,
		Amount:       amount,
		RecognizedOn: recognizedOn,
	}
	_, err = db.DB().Model(revenueRecognition).Insert()
	return err
}

func (g *gateway) FindProductType(productID uint) (enums.ProductType, error) {
	var err error

	var productType enums.ProductType
	if err = db.DB().
		Model((*models.Product)(nil)).
		Column("product_type").
		Where("product_id = ?", productID).
		Select(&productType); err != nil {
		return enums.WordProcessor, err
	}

	return productType, nil
}

func (g *gateway) FindRevenueRecognitionAmount(contractID uint, recognizedOn time.Time) ([]float64, error) {
	var err error

	var revenueRecognitionAmount []float64
	if err = db.DB().
		Model((*models.RevenueRecognition)(nil)).
		Column("amount").
		Where("contract_id = ?", contractID).
		Where("recognized_on <= ?", recognizedOn).
		Select(&revenueRecognitionAmount); err != nil {
		return nil, err
	}

	return revenueRecognitionAmount, err
}
