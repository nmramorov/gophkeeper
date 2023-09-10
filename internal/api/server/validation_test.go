package server

import (
	"context"
	"sync"
	"testing"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ValidationTestSuite struct {
	Server StorageServer

	suite.Suite
	ctx context.Context
}

func (suite *ValidationTestSuite) SetupTest() {
	suite.Server = StorageServer{
		Storage: &db.InMemoryDB{
			Users:       sync.Map{},
			Credentials: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	err := suite.Server.Storage.SaveUser(suite.ctx, models.User{
		Login:    "initial login",
		Password: "initial password",
		Token:    "initial login/initial password/test salt",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving user: %e", err)
	}
}

func (suite *ValidationTestSuite) TestValidateRequestSuccess() {
	result := suite.Server.ValidateRequest(suite.ctx, "initial login/initial password/test salt")
	require.Equal(suite.T(), "", result)
}

func (suite *ValidationTestSuite) TestValidateRequestInvalidUser() {
	result := suite.Server.ValidateRequest(suite.ctx, "wrong login/wrong password/test salt")
	require.Equal(suite.T(), "authorization error: wrong username", result)
}

func (suite *ValidationTestSuite) TestValidateRequestInvalidToken() {
	result := suite.Server.ValidateRequest(suite.ctx, "initial login/initial password/wrong salt")
	require.Equal(suite.T(), "authorization error: invalid token", result)
}

func (suite *ValidationTestSuite) TestValidateToken() {
	result := ValidateToken("test token", "stored token")
	require.Equal(suite.T(), false, result)
	result = ValidateToken("test token", "test token")
	require.Equal(suite.T(), true, result)
}

func TestValidationTestSuite(t *testing.T) {
	suite.Run(t, new(ValidationTestSuite))
}
