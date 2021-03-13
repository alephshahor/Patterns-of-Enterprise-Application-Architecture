package table_modules

import (
	"testing"
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/table_module/gateway"

	"github.com/stretchr/testify/assert"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
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
	var dataset = gateway.Gateway()
	var now = time.Now()

	var contractTableModule = NewContractTableModule(dataset)
	assert.NotNil(suite.T(), contractTableModule)

	var revenueRecognitionTableModule = NewRevenueRecognitionTableModule(dataset)
	assert.NotNil(suite.T(), revenueRecognitionTableModule)

	var contractID uint

	// Word processor
	contractID, err = contractTableModule.Create(1, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), contractID)

	var recognizedAmount float64
	recognizedAmount, err = revenueRecognitionTableModule.CalculateRecognizedAmount(contractID, now)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(60), recognizedAmount)

	// Database
	contractID, err = contractTableModule.Create(2, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), contractID)

	var timeOffsets = []int{0, 30, 60}

	for i, timeOffset := range timeOffsets {
		recognizedAmount, err = revenueRecognitionTableModule.CalculateRecognizedAmount(contractID, now.AddDate(0, 0, timeOffset))
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), float64(20*(i+1)), recognizedAmount)
	}

	// Spreadsheet
	contractID, err = contractTableModule.Create(3, 60, now)

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), contractID)

	timeOffsets = []int{0, 60, 90}

	for i, timeOffset := range timeOffsets {
		recognizedAmount, err = revenueRecognitionTableModule.CalculateRecognizedAmount(contractID, now.AddDate(0, 0, timeOffset))
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), float64(20*(i+1)), recognizedAmount)
	}
}

func TestTableModulesTestSuite(t *testing.T) {
	suite.Run(t, new(TableModulesTestSuite))
}
