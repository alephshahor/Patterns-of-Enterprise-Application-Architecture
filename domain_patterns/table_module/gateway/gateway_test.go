package gateway

import (
	"testing"
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/enums"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GatewayTestSuite struct {
	suite.Suite
}

func (suite *GatewayTestSuite) SetupTest() {
	cmd.Execute()
}

func (suite *GatewayTestSuite) TestGateway() {
	assert.NotNil(suite.T(), Gateway())
}

func (suite *GatewayTestSuite) TestCreateContract() {
	var err error
	var contract = &models.Contract{
		ProductID:  1,
		Revenue:    60,
		DateSigned: time.Now(),
	}
	err = Gateway().CreateContract(contract)
	assert.Nil(suite.T(), err)
}

func (suite *GatewayTestSuite) TestCreateRevenueRecognitions() {
	var err error
	var revenueRecognitions []*models.RevenueRecognition

	revenueRecognitions = append(revenueRecognitions, &models.RevenueRecognition{
		ContractID:   1,
		Amount:       60,
		RecognizedOn: time.Now(),
	})

	err = Gateway().CreateRevenueRecognitions(revenueRecognitions)
	assert.Nil(suite.T(), err)
}

func (suite *GatewayTestSuite) TestFindContractByID() {
	var err error
	var contract *models.Contract

	contract, err = Gateway().FindContractByID(1)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), contract)
	assert.Equal(suite.T(), contract.ContractID, uint(1))
	assert.Equal(suite.T(), contract.ProductID, uint(1))
	assert.Equal(suite.T(), contract.Revenue, float64(60))
}

func (suite *GatewayTestSuite) TestFindProductByID() {
	var err error
	var product *models.Product

	product, err = Gateway().FindProductByID(1)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), product)
	assert.Equal(suite.T(), product.ProductID, uint(1))
	assert.Equal(suite.T(), product.ProductType, enums.WordProcessor)
	assert.Equal(suite.T(), product.ProductName, "word_processor_product")
}

func (suite *GatewayTestSuite) TestFindRevenueRecognitionForContractBeforeDate() {
	var err error
	var revenueRecognitions []*models.RevenueRecognition
	revenueRecognitions, err = Gateway().FindRevenueRecognitionForContractBeforeDate(1, time.Now())

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), len(revenueRecognitions))
}

func TestGatewaySuite(t *testing.T) {
	suite.Run(t, new(GatewayTestSuite))
}
