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

func TestContractServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ContractServiceTestSuite))
}
