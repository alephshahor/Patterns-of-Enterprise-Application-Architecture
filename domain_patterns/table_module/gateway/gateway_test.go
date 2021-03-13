package gateway

import (
	"testing"
	"time"

	"github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/cmd"
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
	var contractID uint

	contractID, err = Gateway().CreateContract(1, 60, time.Now())

	assert.NotNil(suite.T(), err)
	assert.NotZero(suite.T(), contractID)
}

func (suite *GatewayTestSuite) TestCreateRevenueRecognition() {
	var err error
	var contractID uint

	contractID, err = Gateway().CreateContract(1, 60, time.Now())

	assert.NotNil(suite.T(), err)
	assert.NotZero(suite.T(), contractID)

	err = Gateway().CreateRevenueRecognition(contractID, 60, time.Now())
	assert.NotNil(suite.T(), err)
}

func TestGatewaySuite(t *testing.T) {
	suite.Run(t, new(GatewayTestSuite))
}
