package table_modules

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"
)

type revenueRecognitionTableModule struct {
	dataset gateway.IGateway
}

func NewRevenueRecognitionTableModule(dataset gateway.IGateway) *revenueRecognitionTableModule {
	return &revenueRecognitionTableModule{
		dataset: dataset,
	}
}

func (m *revenueRecognitionTableModule) Create(contractID uint, amount float64, recognizedOn time.Time) error {
	return m.dataset.CreateRevenueRecognition(contractID, amount, recognizedOn)
}

func (m *revenueRecognitionTableModule) CalculateRecognizedAmount(contractID uint, recognizedOn time.Time) (float64, error) {
	var err error

	var recognizedAmount []float64
	if recognizedAmount, err = m.dataset.FindRevenueRecognitionAmount(contractID, recognizedOn); err != nil {
		return 0, err
	}

	var totalAmount float64
	for _, amount := range recognizedAmount {
		totalAmount += amount
	}

	return totalAmount, err
}
