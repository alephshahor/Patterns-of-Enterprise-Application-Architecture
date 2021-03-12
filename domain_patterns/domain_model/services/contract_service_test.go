package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/domain_patterns/domain_model/domain_models"
	"github.com/stretchr/testify/suite"
)

type ContractServiceTestSuite struct {
	suite.Suite
}

func (suite *ContractServiceTestSuite) SetupTest() {
	cmd.Execute()
}

func (suite *ContractServiceTestSuite) TestCreateContract() {
	var err error
	var contract *domain_models.Contract

	var now = time.Now()
	contract, err = CreateContract(1, 60, now)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), contract.ProductID, uint(1))
	assert.Equal(suite.T(), contract.Revenue, float64(60))
	assert.Equal(suite.T(), contract.DateSigned, now)
}

func (suite *ContractServiceTestSuite) TestCalculateRevenueRecognitionForWordProcessor() {
	var err error
	var contract *domain_models.Contract

	var now = time.Now()
	contract, err = CreateContract(1, 60, now)

	assert.Nil(suite.T(), err)

	var revenueRecognition float64
	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(60), revenueRecognition)
}

func (suite *ContractServiceTestSuite) TestCalculateRevenueRecognitionForDatabase() {
	var err error
	var contract *domain_models.Contract

	var now = time.Now()
	contract, err = CreateContract(2, 60, now)

	assert.Nil(suite.T(), err)

	var revenueRecognition float64
	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(20), revenueRecognition)

	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned.AddDate(0, 0, 30))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(40), revenueRecognition)

	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned.AddDate(0, 0, 60))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(60), revenueRecognition)
}

func (suite *ContractServiceTestSuite) TestCalculateRevenueRecognitionForSpreadsheet() {
	var err error
	var contract *domain_models.Contract

	var now = time.Now()
	contract, err = CreateContract(3, 60, now)

	assert.Nil(suite.T(), err)

	var revenueRecognition float64
	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(20), revenueRecognition)

	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned.AddDate(0, 0, 60))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(40), revenueRecognition)

	revenueRecognition, err = CalculateRevenueRecognitions(contract.ContractID, contract.DateSigned.AddDate(0, 0, 90))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(60), revenueRecognition)
}

func TestContractServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ContractServiceTestSuite))
}
