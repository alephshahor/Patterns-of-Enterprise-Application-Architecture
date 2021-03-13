package table_modules

import (
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"
	"time"
)

type revenueRecognitionTableModule struct {
	dataset gateway.IGateway
}

func NewRevenueRecognitionTableModule(dataset gateway.IGateway) *revenueRecognitionTableModule {
	return &revenueRecognitionTableModule{
		dataset: dataset,
	}
}

func (m *revenueRecognitionTableModule) Create(contractID uint, revenue float64, dateSigned time.Time) error {
	return m.dataset.CreateRevenueRecognition(contractID, revenue, dateSigned)
}
