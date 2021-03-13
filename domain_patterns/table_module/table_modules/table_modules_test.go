package table_modules

import (
	"testing"
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"

	"github.com/stretchr/testify/assert"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
	"github.com/stretchr/testify/suite"
)

type TableModulesTestSuite struct {
	suite.Suite
}

func (suite *TableModulesTestSuite) SetupTest() {
	cmd.Execute()
}

func (suite *TableModulesTestSuite) TestCreateContract() {
	var err error
	var contract *models.Contract
	var now = time.Now()

	// Word processor
	contract, err = ContractTableModule().Create(1, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)
	assert.Equal(suite.T(), uint(1), contract.ProductID)
	assert.Equal(suite.T(), float64(60), contract.Revenue)
	assert.Equal(suite.T(), now, contract.DateSigned)

	var revenueRecognitions []*models.RevenueRecognition
	revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contract.ContractID, now)

	assert.Equal(suite.T(), 1, len(revenueRecognitions))
	assert.Equal(suite.T(), contract.ContractID, revenueRecognitions[0].ContractID)
	assert.Equal(suite.T(), float64(60), revenueRecognitions[0].Amount)

	// Database
	contract, err = ContractTableModule().Create(2, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)
	assert.Equal(suite.T(), uint(2), contract.ProductID)
	assert.Equal(suite.T(), float64(60), contract.Revenue)
	assert.Equal(suite.T(), now, contract.DateSigned)

	var timeOffsets = []int{0, 30, 60}

	for i, timeOffset := range timeOffsets {
		revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contract.ContractID, now.AddDate(0, 0, timeOffset))
		assert.Equal(suite.T(), i+1, len(revenueRecognitions))
		for j := range revenueRecognitions {
			assert.Equal(suite.T(), contract.ContractID, revenueRecognitions[j].ContractID)
			assert.Equal(suite.T(), contract.Revenue/3, revenueRecognitions[j].Amount)
		}
	}

	// Spreadsheet
	contract, err = ContractTableModule().Create(3, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)
	assert.Equal(suite.T(), uint(3), contract.ProductID)
	assert.Equal(suite.T(), float64(60), contract.Revenue)
	assert.Equal(suite.T(), now, contract.DateSigned)

	timeOffsets = []int{0, 60, 90}

	for i, timeOffset := range timeOffsets {
		revenueRecognitions, err = gateway.Gateway().FindRevenueRecognitionForContractBeforeDate(contract.ContractID, now.AddDate(0, 0, timeOffset))
		assert.Equal(suite.T(), i+1, len(revenueRecognitions))
		for j := range revenueRecognitions {
			assert.Equal(suite.T(), contract.ContractID, revenueRecognitions[j].ContractID)
			assert.Equal(suite.T(), contract.Revenue/3, revenueRecognitions[j].Amount)
		}
	}
}

func (suite *TableModulesTestSuite) TestCalculateRevenueRecognitions() {
	var err error
	var now = time.Now()

	var contractTableModule = ContractTableModule()
	assert.NotNil(suite.T(), contractTableModule)

	var contract *models.Contract
	contract, err = contractTableModule.Create(1, 60, now)
	assert.Nil(suite.T(), err)

	var revenueRecognition float64
	revenueRecognition, err = contractTableModule.CalculateRevenueRecognition(contract.ContractID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), revenueRecognition, float64(60))
}

func TestTableModulesTestSuite(t *testing.T) {
	suite.Run(t, new(TableModulesTestSuite))
}
