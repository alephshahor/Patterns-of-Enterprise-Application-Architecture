package gateway

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
)

// Table Data Gateway (144)
type IGateway interface {
	CreateContract(productID uint, revenue float64, dateSigned time.Time) (uint, error)
	CreateRevenueRecognition(contractID uint, amount float64, recognizedOn time.Time) error
	FindProductType(productID uint) (enums.ProductType, error)
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

	var product = new(models.Product)
	if err = db.DB().
		Model(product).
		Where("product_id = ?", productID).
		First(); err != nil {
		return enums.WordProcessor, err
	}

	return product.ProductType, nil
}
