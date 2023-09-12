package server

import (
	"context"
	"sync"
	"testing"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite

	Server StorageServer
	ctx    context.Context
}

func (suite *AuthTestSuite) SetupTest() {
	suite.Server = StorageServer{
		gctx: context.Background(),
		Storage: &db.InMemoryDB{
			Users:       sync.Map{},
			Credentials: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	sessions := []models.Session{
		{
			UUID:  "test session 1",
			Token: "initial login/initial password/test salt 1",
		},
		{
			UUID:  "test session 2",
			Token: "initial login/initial password/test salt 2",
		},
	}
	err := suite.Server.Storage.SaveUser(suite.ctx, models.User{
		Login:    "initial login",
		Password: "initial password",
		Sessions: sessions,
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving user: %e", err)
	}
	err = suite.Server.Storage.SaveCredentials(suite.ctx, models.CredentialsData{
		UUID:     "initial UUID",
		Login:    "initial login",
		Password: "initial password",
		Meta:     "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving credentials: %e", err)
	}
}

func (suite *AuthTestSuite) TestLoginInvalid() {
	req := &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "test login",
			Password: "test password",
		},
	}
	resp, err := suite.Server.Login(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "user test login not found", resp.Error)
	require.Equal(suite.T(), "", resp.Token)
}

func (suite *AuthTestSuite) TestLoginValid() {
	req := &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "initial login",
			Password: "initial password",
		},
	}
	resp, err := suite.Server.Login(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial login/initial password/salt", resp.Token)
}

func (suite *AuthTestSuite) TestRegisterValid() {
	req := &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    "new login",
			Password: "new password",
		},
	}
	resp, err := suite.Server.Register(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *AuthTestSuite) TestRegisterInvalid() {
	req := &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    "initial login",
			Password: "initial password",
		},
	}
	resp, err := suite.Server.Register(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "user already exists initial login", resp.Error)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
