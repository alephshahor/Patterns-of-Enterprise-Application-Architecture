package transaction_script

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/db"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"
	"github.com/stretchr/testify/suite"
)

type TransactionScriptTestSuite struct {
	suite.Suite
}

func (suite *TransactionScriptTestSuite) SetupTest() {
	cmd.Execute()
}

func (suite *TransactionScriptTestSuite) TestCreateWordProcessorContract() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(1, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(revenueRecognitions))
}

func (suite *TransactionScriptTestSuite) TestCreateSpreadsheetContract() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(2, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, len(revenueRecognitions))
}

func (suite *TransactionScriptTestSuite) TestDatabaseContract() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(3, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, len(revenueRecognitions))
}

func (suite *TransactionScriptTestSuite) TestCalculateWordProcessorRevenueRecognitions() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(1, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(revenueRecognitions))

	var totalRevenue float64
	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now())

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), totalRevenue, float64(70))
}

func (suite *TransactionScriptTestSuite) TestCalculateDatabaseRevenueRecognitions() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(2, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, len(revenueRecognitions))

	var totalRevenue float64
	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now())

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(23), math.Round(totalRevenue))

	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now().AddDate(0, 0, 30))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(47), math.Round(totalRevenue))

	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now().AddDate(0, 0, 60))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(70), math.Round(totalRevenue))

}

func (suite *TransactionScriptTestSuite) TestCalculateSpreadsheetRevenueRecognitions() {
	var err error
	var contract *models.Contract
	contract, err = CreateContract(3, 70, time.Now())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)

	var revenueRecognitions []*models.RevenueRecognition
	err = db.DB().Model(&revenueRecognitions).
		Where("contract_id = ?", contract.ContractID).
		Select()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, len(revenueRecognitions))

	var totalRevenue float64
	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now())

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(23), math.Round(totalRevenue))

	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now().AddDate(0, 0, 30))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(23), math.Round(totalRevenue))

	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now().AddDate(0, 0, 60))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(47), math.Round(totalRevenue))

	totalRevenue, err = CalculateRevenueRecognitions(contract.ContractID, time.Now().AddDate(0, 0, 90))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(70), math.Round(totalRevenue))

}

func TestTransactionScriptTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionScriptTestSuite))
}
