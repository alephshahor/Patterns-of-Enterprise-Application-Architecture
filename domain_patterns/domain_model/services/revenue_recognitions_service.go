package services

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/domain_models"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/gateway"
)

func CreateRevenueRecognitions(revenueRecognitionStrategy domain_models.IRevenueRecognitionStrategy, contractID uint,
	revenue float64, dateSigned time.Time) ([]*domain_models.RevenueRecognition, error) {

	var err error

	var revenueRecognitions []*domain_models.RevenueRecognition
	revenueRecognitions = revenueRecognitionStrategy.CreateRevenueRecognitions(contractID, revenue, dateSigned)

	if err = gateway.Gateway().CreateRevenueRecognitions(revenueRecognitions); err != nil {
		return nil, err
	}

	return revenueRecognitions, err
}
