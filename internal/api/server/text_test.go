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

type TextTestSuite struct {
	suite.Suite

	Server StorageServer
	ctx    context.Context
}

func (suite *TextTestSuite) SetupTest() {
	suite.Server = StorageServer{
		Storage: &db.InMemoryDB{
			Users: sync.Map{},
			Texts: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	sessions := []models.Session{
		{
			UUID:  "test session 1",
			Token: "initial login/initial password/salt",
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
	err = suite.Server.Storage.SaveText(suite.ctx, models.TextData{
		UUID: "initial UUID",
		Data: "initial text",
		Meta: "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving text: %e", err)
	}
}

func (suite *TextTestSuite) TestSaveTextInvalidToken() {
	req := &pb.SaveTextDataRequest{
		Token: "initial login/initial password/invalid salt",
		Data: &pb.TextData{
			Uuid: "test uuid",
			Data: "new text",
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveText(suite.ctx, req)
	require.Error(suite.T(), ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *TextTestSuite) TestSaveTextSuccess() {
	req := &pb.SaveTextDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.TextData{
			Uuid: "test uuid",
			Data: "new text",
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveText(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *TextTestSuite) TestLoadTextSuccess() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), "initial text", resp.Data.Data)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *TextTestSuite) TestLoadTextAuthError() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/wrong salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.Error(suite.T(), ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *TextTestSuite) TestLoadTextNoSuchID() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.Error(suite.T(), ErrDatabaseError, err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestTextTestSuite(t *testing.T) {
	suite.Run(t, new(TextTestSuite))
}
