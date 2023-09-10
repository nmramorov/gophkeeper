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

type BinaryTestSuite struct {
	suite.Suite

	Server StorageServer
	ctx    context.Context
}

func (suite *BinaryTestSuite) SetupTest() {
	suite.Server = StorageServer{
		Storage: &db.InMemoryDB{
			Users:    sync.Map{},
			Binaries: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	err := suite.Server.Storage.SaveUser(suite.ctx, models.User{
		Login:    "initial login",
		Password: "initial password",
		Token:    "initial login/initial password/salt",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving user: %e", err)
	}
	err = suite.Server.Storage.SaveBinary(suite.ctx, models.BinaryData{
		UUID: "initial UUID",
		Data: []byte{255, 255, 1},
		Meta: "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving text: %e", err)
	}
}

func (suite *BinaryTestSuite) TestSaveBinaryInvalidToken() {
	req := &pb.SaveBinaryDataRequest{
		Token: "initial login/initial password/invalid salt",
		Data: &pb.BinaryData{
			Uuid: "test uuid",
			Data: []byte{255, 255, 1},
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *BinaryTestSuite) TestSaveBinarySuccess() {
	req := &pb.SaveBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.BinaryData{
			Uuid: "test uuid",
			Data: []byte{255, 255, 2},
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *BinaryTestSuite) TestLoadBinarySuccess() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), []byte{255, 255, 1}, resp.Data.Data)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *BinaryTestSuite) TestLoadBinaryAuthError() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/wrong salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *BinaryTestSuite) TestLoadBinaryNoSuchID() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestBinaryTestSuite(t *testing.T) {
	suite.Run(t, new(BinaryTestSuite))
}
