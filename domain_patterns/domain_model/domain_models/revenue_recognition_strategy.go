package domain_models

import (
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"
)

type IRevenueRecognitionStrategy interface {
	CreateRevenueRecognitions(contractID uint, amount float64, dateSigned time.Time) []*RevenueRecognition
}

type WordProcessorStrategy struct{}

func (s *WordProcessorStrategy) CreateRevenueRecognitions(contractID uint, amount float64, dateSigned time.Time) []*RevenueRecognition {
	var revenueRecognitions []*RevenueRecognition
	var revenueRecognition = NewRevenueRecognition(contractID, amount, dateSigned)
	revenueRecognitions = append(revenueRecognitions, &revenueRecognition)
	return revenueRecognitions
}

type SpreadsheetStrategy struct{}

func (s *SpreadsheetStrategy) CreateRevenueRecognitions(contractID uint, amount float64, dateSigned time.Time) []*RevenueRecognition {
	var revenueRecognitions []*RevenueRecognition

	var daysOffset = []int{0, 60, 90}
	for i := 0; i < 3; i++ {
		var revenueRecognition = NewRevenueRecognition(contractID, amount/3, dateSigned.AddDate(0, 0, daysOffset[i]))
		revenueRecognitions = append(revenueRecognitions, &revenueRecognition)
	}

	return revenueRecognitions
}

type DatabaseStrategy struct{}

func (s *DatabaseStrategy) CreateRevenueRecognitions(contractID uint, amount float64, dateSigned time.Time) []*RevenueRecognition {
	var revenueRecognitions []*RevenueRecognition

	var daysOffset = []int{0, 30, 60}
	for i := 0; i < 3; i++ {
		var revenueRecognition = NewRevenueRecognition(contractID, amount/3, dateSigned.AddDate(0, 0, daysOffset[i]))
		revenueRecognitions = append(revenueRecognitions, &revenueRecognition)
	}

	return revenueRecognitions
}

func NewRevenueRecognitionStrategy(productType enums.ProductType) IRevenueRecognitionStrategy {

	var revenueRecognitionStrategy IRevenueRecognitionStrategy

	switch productType {
	case enums.WordProcessor:
		revenueRecognitionStrategy = new(WordProcessorStrategy)
	case enums.Spreadsheet:
		revenueRecognitionStrategy = new(SpreadsheetStrategy)
	case enums.Database:
		revenueRecognitionStrategy = new(DatabaseStrategy)
	}

	return revenueRecognitionStrategy
}
