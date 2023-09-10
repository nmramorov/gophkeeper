package server

// import (
// 	"context"
// 	"sync"
// 	"testing"

// 	"github.com/nmramorov/gophkeeper/internal/api/models"
// 	"github.com/nmramorov/gophkeeper/internal/db"
// 	pb "github.com/nmramorov/gophkeeper/internal/proto"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// )

// type TextTestSuite struct {
// 	suite.Suite

// 	Server StorageServer
// 	ctx    context.Context
// }

// func (suite *TextTestSuite) SetupTest() {
// 	suite.Server = StorageServer{
// 		Storage: &db.InMemoryDB{
// 			Users:       sync.Map{},
// 			Credentials: sync.Map{},
// 			Text:        sync.Map{},
// 		},
// 	}
// 	suite.ctx = context.Background()
// 	err := suite.Server.Storage.SaveUser(suite.ctx, models.User{
// 		Login:    "initial login",
// 		Password: "initial password",
// 		Token:    "initial logininitial passwordsalt",
// 	})
// 	if err != nil {
// 		suite.T().Errorf("Error setup - saving user: %e", err)
// 	}
// 	err = suite.Server.Storage.SaveCredentials(suite.ctx, models.CredentialsData{
// 		UUID:     "initial UUID",
// 		Login:    "initial login",
// 		Password: "initial password",
// 		Meta:     "initial Meta",
// 	})
// 	if err != nil {
// 		suite.T().Errorf("Error setup - saving credentials: %e", err)
// 	}
// }

// func (suite *CredentialsTestSuite) TestSaveCredentialsInvalidToken() {
// 	req := &pb.SaveCredentialsDataRequest{
// 		Token: "invalid token",
// 		Data: &pb.CredentialsData{
// 			Uuid:     "test uuid",
// 			Login:    "invalid login",
// 			Password: "new password",
// 			Meta:     nil,
// 		},
// 	}
// 	resp, err := suite.Server.SaveCredentials(suite.ctx, req)
// 	require.NoError(suite.T(), err)
// 	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
// }

// func (suite *CredentialsTestSuite) TestSaveCredentialsSuccess() {
// 	req := &pb.SaveCredentialsDataRequest{
// 		Token: "initial logininitial passwordsalt",
// 		Data: &pb.CredentialsData{
// 			Uuid:     "test uuid",
// 			Login:    "new login",
// 			Password: "new password",
// 			Meta:     nil,
// 		},
// 	}
// 	resp, err := suite.Server.SaveCredentials(suite.ctx, req)
// 	require.NoError(suite.T(), err)
// 	require.Equal(suite.T(), "", resp.Error)
// }

// func (suite *CredentialsTestSuite) TestLoadCredentialsSuccess() {
// 	req := &pb.LoadCredentialsDataRequest{
// 		Token: "initial logininitial passwordsalt",
// 		Uuid:  "initial UUID",
// 	}
// 	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
// 	require.NoError(suite.T(), err)
// 	require.Equal(suite.T(), "", resp.Error)
// 	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
// 	require.Equal(suite.T(), "initial login", resp.Data.Login)
// 	require.Equal(suite.T(), "initial password", resp.Data.Password)
// 	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
// }

// func (suite *CredentialsTestSuite) TestLoadCredentialsAuthError() {
// 	req := &pb.LoadCredentialsDataRequest{
// 		Token: "wrong token",
// 		Uuid:  "initial UUID",
// 	}
// 	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
// 	require.NoError(suite.T(), err)
// 	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
// }

// func (suite *CredentialsTestSuite) TestLoadCredentialsNoSuchID() {
// 	req := &pb.LoadCredentialsDataRequest{
// 		Token: "wrong token",
// 		Uuid:  "wrong UUID",
// 	}
// 	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
// 	require.NoError(suite.T(), err)
// 	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
// }

// func TestCredentialsTestSuite(t *testing.T) {
// 	suite.Run(t, new(CredentialsTestSuite))
// }
