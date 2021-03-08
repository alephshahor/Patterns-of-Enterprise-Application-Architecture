package services

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/domain_models"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/gateway"
)

func FindProductByID(productID uint) (*domain_models.Product, error) {
	var err error

	var product *domain_models.Product
	if product, err = gateway.Gateway().FindProductByID(productID); err != nil {
		return nil, err
	}

	product.RevenueRecognitionStrategy = domain_models.NewRevenueRecognitionStrategy(product.ProductType)

	return product, nil
}
